package gojsa

import (
	"encoding/json"
	"fmt"
)

type SchemaOrSchemas struct {
	IsSchema bool
	Schema   *Schema
	Schemas  []*Schema
}

func (b *SchemaOrSchemas) UnmarshalJSON(data []byte) (err error) {
	var i interface{}
	if err = json.Unmarshal(data, &i); err != nil {
		return err
	}
	switch t := i.(type) {
	default:
		return fmt.Errorf("unexpected type %T", t)
	case map[string]interface{}:
		s := new(Schema)
		if err = json.Unmarshal(data, &s); err != nil {
			return err
		}
		b.IsSchema = true
		b.Schema = s
	case []interface{}:
		s := new([]*Schema)
		if err = json.Unmarshal(data, &s); err != nil {
			return err
		}
		b.Schemas = *s
	}
	return nil
}

func (s *SchemaOrSchemas) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var schema *Schema
	var schemas []*Schema
	if err = unmarshal(&schema); err == nil {
		s.IsSchema = true
		s.Schema = schema
		return nil
	}
	if err = unmarshal(&schemas); err == nil {
		s.Schemas = schemas
		return nil
	}
	return fmt.Errorf("unexpected type")
}

func (s *SchemaOrSchemas) Collect(schemas *map[string]*Schema) (err error) {
	if s.IsSchema {
		return s.Schema.Collect(schemas)
	}
	return nil
}