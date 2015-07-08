package jsonschema

import (
	"encoding/json"
	"fmt"
)

type SchemaOrStrings struct {
	IsSchema bool
	Schema   *Schema
	Strings  []string
}

func (b *SchemaOrStrings) UnmarshalJSON(data []byte) (err error) {
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
		var strs []string
		for _, v := range t {
			str, ok := v.(string)
			if !ok {
				fmt.Errorf("unexpected type %T", t)
			}
			strs = append(strs, str)
		}
		b.Strings = strs
	}
	return nil
}

// func (s *SchemaOrStrings) Collect(schemas map[string]*Schema, p string) (err error) {
// 	if s.IsSchema {
// 		return s.Collect(schemas, p)
// 	}
// 	return nil
// }

func (s *SchemaOrStrings) Resolve(schemas map[string]*Schema, root *Schema) (err error) {
	if s.IsSchema {
		return s.Resolve(schemas, root)
	}
	return nil
}
