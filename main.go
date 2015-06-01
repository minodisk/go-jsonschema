package main

import (
	"encoding/json"
	"log"
	"os"
)

type Ref struct {
	Ref string `json:"$ref"`
}

type Unmershaler interface{}

func (u *Unmershaler) UnmarshalJSON(buf []byte) error {
	log.Println("UnmarshalJSON:", string(buf))
	var r Ref
	if err := json.Unmarshal(buf, &r); err == nil {
		u.Ref = r.Ref
		return nil
	}
	return nil
}

func main() {
	file, err := os.Open("heroku.schema.json")
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(file)
	var s Schema
	if err := dec.Decode(&s); err != nil {
		panic(err)
	}
	log.Printf("%+v", s)
}

type Schema struct {
	// Type        string
	// ID          string
	// Title       string
	// Description string
	Definitions map[string]Unmershaler
	Properties  map[string]Unmershaler
	// Required    []string
	// Links       []map[string]interface{}
}

type Link struct {
	Title        string
	Description  string
	Rel          string
	EncType      string
	Method       string
	HRef         string
	Schema       Schema
	MediaType    string
	TargetSchema Schema
}

type AnyOf struct {
	Ref string
}

type Prop struct {
	Description string
	Example     interface{}
	Format      string
	ReadOnly    bool
	Type        []string
}
