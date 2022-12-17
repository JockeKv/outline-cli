package api

import (
	"fmt"
)

type Collections struct {
	client *Client
}

func (c *Collections) getPath(path string) string {
	// fmt.Println("getPath")
	return fmt.Sprintf("%s/api/collections.%s", c.client.Config.Host, path)
}

func (c *Collections) List() (*[]Collection, error) {
	// fmt.Printf("Posting to: %s\n", d.getPath("list"))
	resp, err := c.client.post(c.getPath("list"), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to post: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		body := Error{}
		err = unmarshal(resp.Body, &body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("API Error: %d, %s", resp.StatusCode, body.Message)
	}
	doc := []Collection{}
	err = unmarshalData(resp.Body, &doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}
