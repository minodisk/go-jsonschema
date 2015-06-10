package gojsa

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"path"
	"regexp"
)

var (
	rBraceBracket = regexp.MustCompile(`{\((.*)\)}`)
)

type Link struct {
	HRef         *HRef
	Rel          string
	Title        string
	TargetSchema *Schema
	MediaType    string
	Method       string
	EncType      string
	Schema       *Schema

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

func (l Link) QueryString() string {
	if l.Method != "GET" {
		return ""
	}
	if l.Schema == nil {
		return ""
	}
	return fmt.Sprintf("?%s", l.Schema.QueryString())
}

func (l Link) ContentType() string {
	return l.EncType
}

func (l Link) HasRequestBody() bool {
	return l.Schema != nil
}

func (l Link) RequestBody() string {
	if l.HasRequestBody() {
		return l.Schema.Body()
	}
	return ""
}

func (l Link) HasResponseBody() bool {
	return l.TargetSchema != nil
}

func (l Link) ResponseBody() string {
	if l.HasResponseBody() {
		return l.TargetSchema.Body()
	}
	return ""
}

type Media struct {
	Type           string
	BinaryEncoding string
}

// func Parse(r io.Reader) (s Schema, err error) {
// 	dec := json.NewDecoder(r)
// 	if err := dec.Decode(&s); err != nil {
// 		return s, err
// 	}
// 	return s, nil
// }

type HRef struct {
	id      int
	value   string
	example string
}

func (h *HRef) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	h.value = rBraceBracket.ReplaceAllStringFunc(s, func(s string) string {
		d, err := url.QueryUnescape(s)
		if err != nil {
			return s
		}
		return d
	})
	return nil
}

func (h *HRef) Resolve(schemas *map[string]*Schema) error {
	h.example = h.replaceBraceBracket(func(s string) string {
		schema := (*schemas)[s]
		return schema.Example.String()
	})
	h.id = rand.Int()
	return nil
}

func (h HRef) String() string {
	return h.value
}

func (h HRef) ColonString() string {
	return h.replaceBraceBracket(func(s string) string {
		return fmt.Sprintf(":%s", path.Base(s))
	})
}

func (h *HRef) ExampleString() string {
	return h.example
}

func (h HRef) replaceBraceBracket(replacer func(string) string) string {
	return rBraceBracket.ReplaceAllStringFunc(h.value, func(s string) string {
		m := rBraceBracket.FindStringSubmatch(s)
		return replacer(m[1])
	})
}
