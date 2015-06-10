package gojsa

import (
	"encoding/json"
	"net/url"
)

type Properties struct {
	// SchemaMap
	Schemas map[string]*Schema
}

func (s *Properties) UnmarshalJSON(data []byte) error {
	var schemas map[string]*Schema
	if err := json.Unmarshal(data, &schemas); err != nil {
		return err
	}
	s.Schemas = schemas
	return nil
}

func (s *Properties) Resolve(schemas *map[string]*Schema) error {
	for _, schema := range s.Schemas {
		if err := schema.Resolve(schemas); err != nil {
			return err
		}
	}
	return nil
}

func (p *Properties) QueryString() string {
	v := url.Values{}
	for name, schema := range p.Schemas {
		v.Set(name, schema.Example.String())
	}
	return v.Encode()
}
