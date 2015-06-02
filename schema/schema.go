package schema

type Node struct {
	Ref string `json:"$ref"`

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
	Node

	ID          string
	Definitions map[string]Property
	Properties  map[string]Property
	Required    []string

	Title string
	Links []Link
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
