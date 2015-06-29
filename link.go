package jsonschema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"text/template"
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

	Boundary = "example_boundary"
	part     = `--{{.boundary}}
Content-Disposition: form-data; name="{{.name}}"

{{.content}}
`
)

var (
	partTmpl = template.Must(template.New("part").Parse(part))
)

type Link struct {
	// http://json-schema.org/latest/json-schema-hypermedia.html
	HRef         *HRef
	Rel          string
	Title        string
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

func (l *Link) Resolve(schemas *map[string]*Schema, root *Schema) (err error) {
	if l.TargetSchema != nil {
		err = l.TargetSchema.Resolve(schemas, root)
		if err != nil {
			return err
		}
	}
	if l.Schema != nil {
		err = l.Schema.Resolve(schemas, root)
		if err != nil {
			return err
		}
	}
	if l.HRef != nil {
		err = l.HRef.Resolve(schemas)
		if err != nil {
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
		return fmt.Sprintf("%s; boundary=%s", l.EncType, Boundary)
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

	d, err := l.Schema.ExampleData(false)
	if err != nil {
		return ""
	}

	if l.IsContentTypeMultipart() {
		s, err := multipartString(d)
		if err != nil {
			log.Printf("can't create a response body as multipart: %s", err)
			return ""
		}
		return s
	}

	b, err := json.MarshalIndent(d, examplePrefix, exampleIndent)
	if err != nil {
		log.Printf("can't create a response body as JSON: %s", err)
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
		d, err = l.TargetSchema.ExampleData(true)
	} else {
		d, err = l.parent.ExampleData(true)
	}
	if err != nil {
		log.Printf("fail to create example data: %s", err)
		return ""
	}

	if l.Rel == "instances" {
		d = []interface{}{d}
	}

	b, err := json.MarshalIndent(d, examplePrefix, exampleIndent)
	if err != nil {
		log.Printf("fail to marshal as JSON: %s", err)
		return ""
	}
	return string(b)
}

func multipartString(data interface{}) (string, error) {
	var str string
	switch d := data.(type) {
	default:
		return "", fmt.Errorf("unsupported data type: %T", d)
	case map[string]interface{}:
		for name, content := range d {
			switch c := content.(type) {
			default:
				return "", fmt.Errorf("unsupported content type: %T", c)
			case *Example:
				b, err := marshalPart(map[string]interface{}{
					"boundary": Boundary,
					"name":     name,
					"content":  c.RawString(),
				})
				if err != nil {
					return "", err
				}
				str = string(b)
			case []interface{}:
				n := fmt.Sprintf("%s[]", name)
				results := []string{}
				for _, v := range c {
					b, err := marshalPart(map[string]interface{}{
						"boundary": Boundary,
						"name":     n,
						"content":  v,
					})
					if err != nil {
						return "", err
					}
					results = append(results, string(b))
				}
				str = strings.Join(results, "")
			}
		}
		return fmt.Sprintf("%s\n--%s--", str, Boundary), nil
	}
}

func marshalPart(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := partTmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
