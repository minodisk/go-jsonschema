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
}

func (t *Type) UnmarshalJSON(data []byte) (err error) {
	var obj interface{}
	if err = json.Unmarshal(data, &obj); err != nil {
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

// func (t Type) Validate(o interface{}) error {
// 	if t.validate(o) {
// 		return nil
// 	}
// 	return TypeError{o}
// }
//
// func (t Type) validate(o interface{}) bool {
// 	if t.isString {
// 		return validateWith(t.string, o)
// 	}
// 	for _, s := range t.strings {
// 		if validateWith(s, o) {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// func validateWith(t string, o interface{}) bool {
// 	switch T(t) {
// 	case TypeInteger:
// 		switch t := o.(type) {
// 		default:
// 			return false
// 		case int, int8, int16, int32, int64:
// 			return true
// 		case float32:
// 			return t == float32(int64(t))
// 		case float64:
// 			return t == float64(int64(t))
// 		}
// 	}
// 	return false
// }
//
// type TypeError struct {
// 	value interface{}
// }
//
// func (e TypeError) Error() string {
// 	return fmt.Sprintf("unexpected type %T", e.value)
// }
