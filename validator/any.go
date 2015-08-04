package validator

import "fmt"

func (v Validator) Type(actual interface{}, expected string) {
	var a string
	switch actual := actual.(type) {
	default:
		return fmt.Errorf("InvalidType: %+v", actual)
	case nil:
		a = "null"
	case bool:
		a = "boolean"
	case int:
		a = "integer"
	case float64:
		a = "number"
	case []interface{}:
		a = "array"
	case map[string]interface{}:
		a = "object"
	}
	if a != expected {
		return TypeError{v, a, expected}
	}
}

type TypeError struct {
	Validator
	Actual, Expected string
}

func (err TypeError) Error() string {
	return fmt.Sprintf("the properties %s in %s should be %s, but %s is specified", err.PropertyName, err.SchemaName, err.Expected, err.Actual)
}
