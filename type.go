package gojsa

import "fmt"

type TypeError struct {
	value interface{}
}

func (e TypeError) Error() string {
	return fmt.Sprintf("unexpected type %T", e.value)
}

type Type struct {
	StringOrStrings
}

func (t Type) Validate(o interface{}) error {
	if t.validate(o) {
		return nil
	}
	return TypeError{o}
}

func (t Type) validate(o interface{}) bool {
	if t.IsString {
		return validateWith(t.String, o)
	}
	for _, s := range t.Strings {
		if validateWith(s, o) {
			return true
		}
	}
	return false
}

func validateWith(t string, o interface{}) bool {
	switch t {
	case "integer":
		switch t := o.(type) {
		default:
			return false
		case int, int8, int16, int32, int64:
			return true
		case float32:
			return t == float32(int64(t))
		case float64:
			return t == float64(int64(t))
		}
	}
	return false
}
