package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"outline/pkg/config"
)

type Client struct {
	Config *config.Config
}

func (c *Client) Documents() *Documents {
	return &Documents{client: c}
}

func (c *Client) Collections() *Collections {
	return &Collections{client: c}
}

func (c *Client) get(path string) (*http.Response, error) {
	client := http.Client{}

	url, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("could not parse url: %v", err)
	}
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("bearer %s", c.Config.Token))

	return client.Do(req)
}

func (c *Client) post(path string, body interface{}) (*http.Response, error) {
	client := http.Client{}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	readerBody := bytes.NewReader(jsonBody)

	url, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("could not parse url: %v", err)
	}
	req, err := http.NewRequest("POST", url.String(), readerBody)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("Authorization", fmt.Sprintf("bearer %s", c.Config.Token))

	// log.Printf("Path: %s", req.URL.Path)
	// log.Printf("Body: %s", string(jsonBody))
	// log.Printf("Headers: %s", req.Header)

	return client.Do(req)
}

func (c *Client) Login() error {
	resp, err := c.post(fmt.Sprintf("%s/api/auth.info", c.Config.Host), nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("authentication failed")
	}
	err = c.Config.Write()
	if err != nil {
		return err
	}
	return nil
}
