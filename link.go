package jsonschema

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/minodisk/jsonschema/multipart"
)

const (
	examplePrefix = ""
	exampleIndent = "  "

	contentTypeJSON      = "application/json"
	contentTypeMultipart = "multipart/form-data"

	methodPost   = "POST"
	methodGet    = "GET"
	methodPut    = "PUT"
	methodPatch  = "PATCH"
	methodDelete = "DELETE"
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

	parent *Schema
}

func (l *Link) UnmarshalJSON(data []byte) error {
	type Tmp Link
	var tmp Tmp
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if tmp.Method == "" {
		tmp.Method = methodGet
	}
	tmp.Method = strings.ToUpper(tmp.Method)
	if tmp.EncType == "" {
		tmp.EncType = contentTypeJSON
	}
	if tmp.MediaType == "" {
		tmp.MediaType = contentTypeJSON
	}
	*l = Link(tmp)
	return nil
}

func (l *Link) SetParent(s *Schema) {
	l.parent = s
}

func (l *Link) Resolve(schemas *map[string]*Schema, root *Schema) error {
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
	return l.HRef.ColonString()
}

func (l Link) ExampleEndpoint() string {
	return l.HRef.ExampleString()
}

func (l Link) MethodEndpoint() string {
	return fmt.Sprintf("%s %s", l.Method, l.Endpoint())
}

func (l Link) QueryString() string {
	if l.Method != methodGet {
		return ""
	}
	if l.Schema == nil {
		return ""
	}
	return fmt.Sprintf("?%s", l.Schema.QueryString())
}

func (l Link) IsContentTypeMultipart() bool {
	return l.EncType == contentTypeMultipart
}

func (l Link) RequestContentType() string {
	if l.IsContentTypeMultipart() {
		return fmt.Sprintf("%s; boundary=%s", l.EncType, multipart.Boundary)
	}
	return l.EncType
}

func (l Link) HasRequestBody() bool {
	return l.Schema != nil
}

func (l Link) RequestBody() string {
	if !l.HasRequestBody() {
		return ""
	}

	d, err := l.Schema.ExampleData()
	if err != nil {
		return ""
	}

	if l.IsContentTypeMultipart() {
		s, err := multipart.Marshal(d)
		if err != nil {
			log.Println("fail to marshal as form data: %s", err)
			return ""
		}
		return s
	}

	b, err := json.MarshalIndent(d, examplePrefix, exampleIndent)
	if err != nil {
		log.Println("fail to marshal as JSON: %s", err)
		return ""
	}
	return string(b)
}

func (l Link) HasResponseBody() bool {
	return l.MediaType != "null"
}

func (l Link) ResponseBody() string {
	if !l.HasResponseBody() {
		return ""
	}

	var d interface{}
	var err error
	if l.TargetSchema != nil {
		d, err = l.TargetSchema.ExampleData()
	} else {
		d, err = l.parent.ExampleData()
	}
	if err != nil {
		log.Println("fail to create example data: %s", err)
		return ""
	}

	if l.Rel == "instances" {
		d = []interface{}{d}
	}

	b, err := json.MarshalIndent(d, examplePrefix, exampleIndent)
	if err != nil {
		log.Println("fail to marshal as JSON: %s", err)
		return ""
	}
	return string(b)
}

func (l Link) ResponseStatus() int {
	switch {
	case l.Method == methodPost:
		return 201
	case !l.HasResponseBody():
		return 204
	default:
		return 200
	}
}

func (l Link) ResponseReasonPhrase() string {
	switch {
	case l.Method == methodPost:
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
