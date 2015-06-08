package gojsa

import (
	"encoding/json"
	"io"
)

type Link struct {
	HRef         string
	Rel          string
	Title        string
	TargetSchema Schema
	MediaType    string
	Method       string
	EncType      string
	Schema       Schema
}

func NewLink(href, rel string) Link {
	return Link{
		HRef:    href,
		Rel:     rel,
		Method:  "GET",
		EncType: "application/json",
	}
}

type Media struct {
	Type           string
	BinaryEncoding string
}

func Parse(r io.Reader) (s Schema, err error) {
	dec := json.NewDecoder(r)
	if err := dec.Decode(&s); err != nil {
		return s, err
	}
	return s, nil
}
