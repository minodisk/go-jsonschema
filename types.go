package gojsa

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

type StringOrStrings struct {
	IsString bool
	String   string
	Strings  []string
}

func (s *StringOrStrings) UnmarshalJSON(data []byte) (err error) {
	var i interface{}
	if err = json.Unmarshal(data, &i); err != nil {
		return err
	}
	switch t := i.(type) {
	default:
		return fmt.Errorf("unexpected type %T", t)
	case string:
		s.IsString = true
		s.String = t
	case []interface{}:
		var strs []string
		for _, v := range t {
			str, ok := v.(string)
			if !ok {
				fmt.Errorf("unexpected type %T", t)
			}
			strs = append(strs, str)
		}
		s.Strings = strs
	}
	return nil
}
