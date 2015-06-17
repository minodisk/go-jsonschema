package jsonschema

import (
	"encoding/json"
	"fmt"
)

type SchemaOrBool struct {
	IsSchema bool
	Schema   *Schema
	Bool     bool
}

func (b *SchemaOrBool) UnmarshalJSON(data []byte) (err error) {
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
	case bool:
		b.Bool = t
	}
	return nil
}

// func (s *SchemaOrBool) Collect(schemas *map[string]*Schema, p string) error {
// 	if s.IsSchema {
// 		return s.Schema.Collect(schemas, p)
// 	}
// 	return nil
// }

func (s *SchemaOrBool) Resolve(schemas *map[string]*Schema, root *Schema) error {
	if s.IsSchema {
		return s.Schema.Resolve(schemas, root)
	}
	return nil
}
