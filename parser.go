package gojsa

import (
	"encoding/json"
	"io"
	"log"
)

type Reference struct {
	Ref string `json:"$ref"`
}

type Node struct {
	Description string
	Type        []string
}

type Root struct {
	Node

	ID          string
	Definitions map[string]Schema
	Properties  map[string]Schema
	Required    []string

	Title string
	Links []Link
}

type Schema struct {
	json.Unmarshaler
	SchemaUnmarshaler
	Node

	ID          string
	Definitions map[string]Property
	Properties  map[string]Property
	Required    []string

	Title string
	Links []Link
}

type SchemaUnmarshaler struct{}

func (s SchemaUnmarshaler) UnmarshalJSON(b []byte) error {
	var tmp map[string]interface{}
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	if id := tmp["id"]; id != nil {
		log.Println(id)
	}
	if ref := tmp["$ref"]; ref != nil {
		log.Println(ref)
	}
	return nil
}

type Property struct {
	Node

	Example  interface{}
	Format   string
	ReadOnly bool
}

type Link struct {
	Title        string
	Description  string
	Rel          string
	Method       string
	HRef         string
	EncType      string
	Schema       Schema
	MediaType    string
	TargetSchema Schema
}

type AnyOf struct {
	Ref string
}

func Parse(r io.Reader) (root Root) {
	dec := json.NewDecoder(r)
	if err := dec.Decode(&root); err != nil {
		panic(err)
	}

	return root
}
