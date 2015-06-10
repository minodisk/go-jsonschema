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

// func (l *Link) UnmarshalJSON(data []byte) error {
// 	type Tmp Link
// 	var tmp Tmp
// 	if err := json.UnmarshalJSON(data, &tmp); err != nil {
// 		return err
// 	}
// 	*l = tmp
// 	return nil
// }

// func NewLink(href, rel string) Link {
// 	return Link{
// 		HRef:    href,
// 		Rel:     rel,
// 		Method:  "GET",
// 		EncType: "application/json",
// 	}
// }

func (l *Link) Resolve(schemas *map[string]*Schema) error {
	return l.HRef.Resolve(schemas)
}

func (l *Link) QueryString() string {
	if l.Method != "GET" {
		return ""
	}
	if l.Schema == nil {
		return ""
	}
	return fmt.Sprintf("?%s", l.Schema.QueryString())
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
