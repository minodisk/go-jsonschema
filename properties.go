package jsonschema

import (
	"log"
	"net/url"
)

type Properties struct {
	SchemaMap
}

func (p *Properties) QueryString() string {
	v := url.Values{}
	for name, schema := range p.Schemas {
		v.Set(name, schema.Example.RawString())
	}
	return v.Encode()
}

func (p *Properties) ExampleData(includesReadOnly bool) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	for name, schema := range p.Schemas {
		if schema == nil {
			log.Printf("schema doesn't exist at '%s' in properties", name)
			continue
		}
		if !includesReadOnly && schema.ReadOnly {
			continue
		}
		m[name] = schema.Example
		v, err := schema.ExampleData(includesReadOnly)
		if err != nil {
			return nil, err
		}
		m[name] = v
	}
	return m, nil
}
