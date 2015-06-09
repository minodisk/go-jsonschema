package gojsa

import "log"

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
	Definitions Schemas //[]*Schema
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
	Example string
	// 4.4. readOnly
	ReadOnly bool
	// 4.5. pathStart
	PathStart string

	Schema string `json:"$schema"`
	Ref    string `json:"$ref"`
}

func (s *Schema) Resolve() (err error) {
	schemas := make(map[string]*Schema)
	s.Collect(&schemas)
	log.Println(schemas)

	return nil
}

func (s *Schema) Collect(schemas *map[string]*Schema) (err error) {
	s.AdditionalItems.Collect(schemas)
	s.Items.Collect(schemas)
	s.AdditionalProperties.Collect(schemas)
	s.Properties.Collect(schemas)
	s.Dependencies.Collect(schemas)
	s.AllOf.Collect(schemas)
	s.AnyOf.Collect(schemas)
	s.OneOf.Collect(schemas)
	s.Not.Collect(schemas)
	s.Definitions.Collect(schemas)
	return nil
}

func (s Schema) Validate(o interface{}) (err error) {
	if err = s.Type.Validate(o); err != nil {
		return err
	}
	return nil
}
