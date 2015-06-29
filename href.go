package jsonschema

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"
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
