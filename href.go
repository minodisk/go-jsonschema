package jsonschema

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"path"
	"regexp"
)

var (
	rBraceBracket = regexp.MustCompile(`{\((.*?)\)}`)
)

type HRef struct {
	value   string
	example string
}

func (h *HRef) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	h.value = rBraceBracket.ReplaceAllStringFunc(s, func(id string) string {
		d, err := url.QueryUnescape(id)
		if err != nil {
			return id
		}
		return d
	})
	return nil
}

func (h *HRef) Resolve(schemas map[string]*Schema) error {
	h.example = h.replaceBraceBracket(func(id string) string {
		schema := schemas[id]
		if schema == nil {
			log.Printf("schema '%s' doesn't exist, it is referred from url", id)
			return ""
		}
		return schema.Example.RawString()
	})
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
