package jsonschema

import (
	"encoding/json"
	"fmt"
	"strings"
)

type PrimitiveType string

const (
	TypeArray   PrimitiveType = "array"
	TypeBoolean PrimitiveType = "boolean"
	TypeInteger PrimitiveType = "integer"
	TypeNumber  PrimitiveType = "number"
	TypeNull    PrimitiveType = "null"
	TypeObject  PrimitiveType = "object"
	TypeString  PrimitiveType = "string"
)

type Type struct {
	types []PrimitiveType
	Go    string
}

func (t *Type) Resolve(f *Format, i *Items, title string) {
	t.Go = map[string]string{
		"array":   "[]interface{}",
		"boolean": "bool",
		"integer": "int64",
		"number":  "float64",
		"null":    "nil",
		"object":  "map[string]interface{}",
		"string":  "string",
	}[string(t.types[0])]
	switch t.Go {
	case "string":
		if f != nil && f.String() == "date-time" {
			t.Go = "time.Time"
		}
	case "[]interface{}":
		if i != nil {
			it := i.GoType()
			if it != "" {
				t.Go = fmt.Sprintf("[]%s", it)
			}
		}
	case "map[string]interface{}":
		if title != "" {
			t.Go = title
		}
	}
}

func (t *Type) Contains(types ...string) bool {
	for _, c := range types {
		for _, d := range t.types {
			if string(d) == c {
				return true
			}
		}
	}
	return false
}

func (t *Type) UnmarshalJSON(data []byte) error {
	var obj interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	switch o := obj.(type) {
	default:
		return fmt.Errorf("unexpected type %T", o)
	case string:
		t.types = []PrimitiveType{PrimitiveType(o)}
	case []interface{}:
		t.types = make([]PrimitiveType, len(o))
		for i, str := range o {
			str, ok := str.(string)
			if !ok {
				fmt.Errorf("unexpected type %T", str)
			}
			t.types[i] = PrimitiveType(str)
		}
	}
	return nil
}

func (t *Type) String() string {
	strs := make([]string, len(t.types))
	for i, p := range t.types {
		strs[i] = string(p)
	}
	return strings.Join(strs, ", ")
}

func (t *Type) Is(p PrimitiveType) bool {
	for _, tp := range t.types {
		if p == tp {
			return true
		}
	}
	return false
}
