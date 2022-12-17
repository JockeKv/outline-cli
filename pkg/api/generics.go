package api

import (
	"encoding/json"
	"io"
)

type SortOpts struct {
	// Field to sort on
	Sort string
	// ASC or DESC
	Direction string
}

type FilterOpts struct {
	UserId             string
	CollectionId       string
	BacklinkDocumentId string
	ParentDocumentId   string
}

type SearchOpts struct {
	UserId       string
	CollectionId string

	IncludeArchived bool
	IncludeDrafts   bool
	// Enum: "day" "week" "month" "year"
	DateFilter string
}

func unmarshalData(data io.Reader, out interface{}) error {
	// fmt.Println("unmarshalData")
	result := map[string]json.RawMessage{}
	bytes, err := io.ReadAll(data)
	if err != nil {
		return err
	}
	// fmt.Printf("Response: %s\n", string(bytes))
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return err
	}

	err = json.Unmarshal(result["data"], out)
	if err != nil {
		return err
	}
	return nil
}

func unmarshal(data io.Reader, out interface{}) error {
	// fmt.Println("unmarshal")
	result := out
	bytes, err := io.ReadAll(data)
	if err != nil {
		return err
	}
	// fmt.Printf("Response: %s\n", string(bytes))
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return err
	}
	return nil
}
