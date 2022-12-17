package api

import (
	"fmt"
)

type Documents struct {
	client *Client
}

func (d *Documents) getPath(path string) string {
	// fmt.Println("getPath")
	return fmt.Sprintf("%s/api/documents.%s", d.client.Config.Host, path)
}

// Get Document info
//
// id <string>
// Unique identifier for the document. Either the UUID or the urlId is acceptable.
//
// shareId <string>
// Unique identifier for a document share, a shareId may be used in place of a document UUID
func (d *Documents) Info(id string) (*Document, error) {
	body := map[string]string{"id": id}
	resp, err := d.client.post(d.getPath("info"), body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		e := Error{}
		err = unmarshal(resp.Body, &e)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("API Error: %v", e)
	}

	doc := Document{}
	err = unmarshalData(resp.Body, &doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

func (d *Documents) List(collectionId string, sort *SortOpts, p *Pagination) (*[]Document, error) {
	// fmt.Printf("Posting to: %s\n", d.getPath("list"))
	body := map[string]string{"collectionId": collectionId}
	resp, err := d.client.post(d.getPath("list"), &body)
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
	doc := []Document{}
	err = unmarshalData(resp.Body, &doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

func (d *Documents) Drafts(sort *SortOpts, p *Pagination) (*[]Document, error) {
	resp, err := d.client.post(d.getPath("drafts"), nil)
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
	doc := []Document{}
	err = unmarshalData(resp.Body, &doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

func (d *Documents) Viewed(sort SortOpts, p Pagination) ([]Document, error) {

	return nil, nil

}

func (d *Documents) Search(query string, searchOpts SearchOpts, p Pagination) ([]Document, error) {
	return nil, nil
}

// *id <string> Document id
//
// title <string> Document title
//
// text <string> Document content
//
// append <bool> Append or replace
//
// publish <bool> Publish document
func (d *Documents) Update(doc *Document) (*Document, error) {
	body := map[string]interface{}{
		"id":           *doc.Id,
		"title":        *doc.Title,
		"text":         *doc.Text,
		"publish":      doc.PublishedAt != nil,
		"collectionId": *doc.CollectionId,
	}
	resp, err := d.client.post(d.getPath("update"), body)
	if err != nil {
		return nil, err
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
	res := Document{}
	err = unmarshalData(resp.Body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (d *Documents) Create(doc *Document) (*Document, error) {
	body := map[string]interface{}{
		"id":           *doc.Id,
		"title":        *doc.Title,
		"text":         *doc.Text,
		"publish":      doc.PublishedAt != nil,
		"collectionId": *doc.CollectionId,
	}
	resp, err := d.client.post(d.getPath("create"), body)
	if err != nil {
		return nil, err
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
	res := Document{}
	err = unmarshalData(resp.Body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
