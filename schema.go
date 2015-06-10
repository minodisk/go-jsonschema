package gojsa

import (
	"encoding/json"
	"fmt"
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
	AdditionalProperties SchemaOrBool      // Schema or bool
	Properties           SchemaMap         //map[string]*Schema
	PatternProperties    map[string]string // regexp
	Dependencies         SchemaOrStrings   // Schema or []string
	// 5.5. Validation keywords for any instance type
	Enum        []interface{}
	Type        Type    // string or []string
	AllOf       Schemas //[]*Schema
	AnyOf       Schemas //[]*Schema
	OneOf       Schemas //[]*Schema
	Not         *Schema
	Definitions Definitions
	// 6. Metadata keywords
	Title       string
	Description string
	Default     interface{}
	// 7. Semantic validation with "format"
	Format string // "date-time" | "email" | "hostname" | "ipv4" | "ipv6" | "uri"

	// JSON Hyper-Schema: Hypertext definitions for JSON Schema
	// json-schema-hypermedia
	// See http://json-schema.org/latest/json-schema-hypermedia.html
	// 4.1. links
	Links []Link
	// 4.3. media
	Media   Media
	Example Example
	// 4.4. readOnly
	ReadOnly bool
	// 4.5. pathStart
	PathStart string

	Schema string `json:"$schema"`
	Ref    string `json:"$ref"`
}

func NewJSON(b []byte) (*Schema, error) {
	s := new(Schema)
	if err := json.Unmarshal(b, s); err != nil {
		return nil, err
	}
	if err := s.Reference(); err != nil {
		return nil, err
	}
	return s, nil
}

// func NewYAML(b []byte) (*Schema, error) {
// 	s := new(Schema)
// 	if err := yaml.Unmarshal(b, s); err != nil {
// 		return nil, err
// 	}
// 	s.Resolve()
//
// 	return s, nil
// }

func (s *Schema) Reference() (err error) {
	schemas := make(map[string]*Schema)
	if err := s.Collect(&schemas, "#"); err != nil {
		return err
	}
	if err := s.Resolve(&schemas); err != nil {
		return err
	}
	return nil
}

func (s *Schema) Collect(schemas *map[string]*Schema, p string) (err error) {
	// if err := s.AdditionalItems.Collect(schemas, p); err != nil {
	// 	return err
	// }
	// if err := s.Items.Collect(schemas, p); err != nil {
	// 	return err
	// }
	// if err := s.AdditionalProperties.Collect(schemas, p); err != nil {
	// 	return err
	// }
	// if err := s.Properties.Collect(schemas, p); err != nil {
	// 	return err
	// }
	// if err := s.Dependencies.Collect(schemas, p); err != nil {
	// 	return err
	// }
	// if err := s.AllOf.Collect(schemas, p); err != nil {
	// 	return err
	// }
	// if err := s.AnyOf.Collect(schemas, p); err != nil {
	// 	return err
	// }
	// if err := s.OneOf.Collect(schemas, p); err != nil {
	// 	return err
	// }
	// if s.Not != nil {
	// 	if err := s.Not.Collect(schemas, p); err != nil {
	// 		return err
	// 	}
	// }
	if err := s.Definitions.Collect(schemas, p); err != nil {
		return err
	}

	return nil
}

func (s *Schema) Resolve(schemas *map[string]*Schema) error {
	if s.Ref != "" {
		schema := (*schemas)[s.Ref]
		if schema == nil {
			return fmt.Errorf("undefined $ref: %s", s.Ref)
		}
		*s = *schema
		return nil
	}

	if err := s.AdditionalItems.Resolve(schemas); err != nil {
		return err
	}
	if err := s.Items.Resolve(schemas); err != nil {
		return err
	}
	if err := s.AdditionalProperties.Resolve(schemas); err != nil {
		return err
	}
	if err := s.Properties.Resolve(schemas); err != nil {
		return err
	}
	if err := s.Dependencies.Resolve(schemas); err != nil {
		return err
	}
	if err := s.AllOf.Resolve(schemas); err != nil {
		return err
	}
	if err := s.AnyOf.Resolve(schemas); err != nil {
		return err
	}
	if err := s.OneOf.Resolve(schemas); err != nil {
		return err
	}
	if s.Not != nil {
		if err := s.Not.Resolve(schemas); err != nil {
			return err
		}
	}
	if err := s.Definitions.Resolve(schemas); err != nil {
		return err
	}
	for _, link := range s.Links {
		link.Resolve(schemas)
	}

	return nil
}

func (s Schema) Validate(o interface{}) (err error) {
	if err = s.Type.Validate(o); err != nil {
		return err
	}
	return nil
}
