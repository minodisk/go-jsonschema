package gojsa

import (
	"encoding/json"
	"fmt"
	"io"
)

type Schema struct {
	// JSON Schema: core definitions and terminology
	// json-schema-core
	// See http://json-schema.org/latest/json-schema-core.html
	ID string

	// JSON Schema: interactive and non interactive validation
	// json-schema-validation
	// See http://json-schema.org/latest/json-schema-validation.html
	// 5.   Validation keywords sorted by instance types
	// 5.1. Validation keywords for numeric instances (number and integer)
	MultipleOf       float64
	Maximum          float64
	ExclusiveMaximum bool
	Minimum          float64
	ExclusiveMinimum bool
	// 5.2. Validation keywords for strings
	MaxLength int
	MinLength int
	Pattern   string // regexp
	// 5.3. Validation keywords for arrays
	AdditionalItems SchemaOrBool    // Schema or bool
	Items           SchemaOrSchemas // Schema or []Schema
	MaxItems        int
	MinItems        int
	UniqueItems     bool
	// 5.4. Validation keywords for objects
	MaxProperties        int
	MinProperties        int
	Required             []string
	AdditionalProperties SchemaOrBool // Schema or bool
	Properties           map[string]*Schema
	PatternProperties    map[string]string // regexp
	Dependencies         SchemaOrStrings   // Schema or []string
	// 5.5. Validation keywords for any instance type
	Enum        []interface{}
	Type        Strings // string or []string
	AllOf       []*Schema
	AnyOf       []*Schema
	OneOf       []*Schema
	Not         *Schema
	Definitions []*Schema
	// 6. Metadata keywords
	Title       string
	Description string
	Default     interface{}
	// 7. Semantic validation with "format"
	Format string // "date-time" | "email" | "hostname" | "ipv4" | "ipv6" | "uri"

	// JSON Hyper-Schema: Hypertext definitions for JSON Schema
	// json-schema-hypermedia
	// See http://json-schema.org/latest/json-schema-hypermedia.html
	Links     []Link
	Media     Media
	ReadOnly  bool
	PathStart string
}

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

type Strings []string

func (s *Strings) UnmarshalJSON(data []byte) (err error) {
	type Tmp Strings
	var tmp Tmp
	if err = json.Unmarshal(data, &tmp); err == nil {
		*s = Strings(tmp)
		return
	}
	var str string
	if err = json.Unmarshal(data, &str); err == nil {
		*s = []string{str}
	}
	return
}

type Link struct {
	HRef         string
	Rel          string
	Title        string
	TargetSchema Schema
	MediaType    string
	Method       string
	EncType      string
	Schema       Schema
}

func NewLink(href, rel string) Link {
	return Link{
		HRef:    href,
		Rel:     rel,
		Method:  "GET",
		EncType: "application/json",
	}
}

type Media struct {
	Type           string
	BinaryEncoding string
}

func Parse(r io.Reader) (s Schema, err error) {
	dec := json.NewDecoder(r)
	if err := dec.Decode(&s); err != nil {
		return s, err
	}
	return s, nil
}
