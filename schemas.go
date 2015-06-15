package jsonschema

import "encoding/json"

type Schemas struct {
	Schemas []*Schema
}

func (s *Schemas) UnmarshalJSON(data []byte) error {
	var schemas []*Schema
	if err := json.Unmarshal(data, &schemas); err != nil {
		return err
	}
	s.Schemas = schemas
	return nil
}

func (s *Schemas) Collect(schemas *map[string]*Schema, p string) error {
	for _, schema := range s.Schemas {
		if err := schema.Collect(schemas, p); err != nil {
			return err
		}
	}
	return nil
}

func (s *Schemas) Resolve(schemas *map[string]*Schema, root *Schema) error {
	for _, schema := range s.Schemas {
		if err := schema.Resolve(schemas, root); err != nil {
			return err
		}
	}
	return nil
}
