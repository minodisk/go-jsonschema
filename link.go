package gojsa

import (
	"encoding/json"
	"fmt"
	"regexp"
)

var (
	rBraceBracket = regexp.MustCompile(`{\((.*)\)}`)
)

type Link struct {
	// http://json-schema.org/latest/json-schema-hypermedia.html
	HRef         *HRef
	Rel          string
	Title        Title
	TargetSchema *Schema
	MediaType    string
	Method       string
	EncType      string
	Schema       *Schema

	// Description is not defined but appears in lots of schema.json
	Description string
}

func (l *Link) UnmarshalJSON(data []byte) error {
	type Tmp Link
	var tmp Tmp
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if tmp.Method == "" {
		tmp.Method = "GET"
	}
	if tmp.EncType == "" {
		tmp.EncType = "application/json"
	}
	if tmp.MediaType == "" {
		tmp.MediaType = "application/json"
	}
	*l = Link(tmp)
	return nil
}

// func NewLink(href, rel string) Link {
// 	return Link{
// 		HRef:    href,
// 		Rel:     rel,
// 		Method:  "GET",
// 		EncType: "application/json",
// 	}
// }

func (l *Link) Resolve(schemas *map[string]*Schema, root *Schema) error {
	if err := l.HRef.Resolve(schemas); err != nil {
		return err
	}
	if l.TargetSchema != nil {
		if err := l.TargetSchema.Resolve(schemas, root); err != nil {
			return err
		}
	}
	if l.Schema != nil {
		if err := l.Schema.Resolve(schemas, root); err != nil {
			return err
		}
	}
	return nil
}

func (l Link) Endpoint() string {
	return l.HRef.ExampleString()
}

func (l Link) QueryString() string {
	if l.Method != "GET" {
		return ""
	}
	if l.Schema == nil {
		return ""
	}
	return fmt.Sprintf("?%s", l.Schema.QueryString())
}

func (l Link) RequestContentType() string {
	return l.EncType
}

func (l Link) HasRequestBody() bool {
	return l.Schema != nil
}

func (l Link) RequestBody() string {
	if l.HasRequestBody() {
		if body, err := l.Schema.ExampleJSON(); err == nil {
			return body
		}
	}
	return ""
}

func (l Link) HasResponseBody() bool {
	return l.TargetSchema != nil
}

func (l Link) ResponseBody() string {
	if l.HasResponseBody() {
		if body, err := l.TargetSchema.ExampleJSON(); err == nil {
			return body
		}
	}
	return ""
}

func (l Link) ResponseStatus() int {
	switch {
	case l.Method == "POST":
		return 201
	case !l.HasResponseBody():
		return 204
	default:
		return 200
	}
}

func (l Link) ResponseReasonPhrase() string {
	switch {
	case l.Method == "POST":
		return "Created"
	case !l.HasResponseBody():
		return "No Content"
	default:
		return "OK"
	}
}

func (l Link) ResponseContentType() string {
	return l.MediaType
}
