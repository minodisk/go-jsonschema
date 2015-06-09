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

func (s *StringOrStrings) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var str string
	var strs []string
	if err = unmarshal(&str); err == nil {
		s.IsString = true
		s.String = str
		return nil
	}
	if err = unmarshal(&strs); err == nil {
		s.Strings = strs
		return nil
	}
	return fmt.Errorf("unexpected type")
}
