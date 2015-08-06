package validator

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func (v Validator) Enum(obj interface{}, def []interface{}) error {
	for _, d := range def {
		if obj == d {
			return nil
		}
	}
	return EnumError{
		Validator: v,
		Specified: obj,
		Enum:      def,
	}
}

type EnumError struct {
	ValidationError
	Validator Validator     `json:"validator"`
	Specified interface{}   `json:"specified"`
	Enum      []interface{} `json:"enum"`
}

func (err EnumError) MarshalJSON() ([]byte, error) {
	type VError EnumError
	obj := make(map[string]interface{})
	obj["message"] = err.Error()
	obj["type"] = reflect.ValueOf(err).Type().String()
	obj["detail"] = VError(err)
	return json.Marshal(obj)
}

func (err EnumError) Error() string {
	return fmt.Sprintf("%s must be one of %v, but `%v`", err.Validator, err.Enum, err.Specified)
}

func (v Validator) Type(obj interface{}, def string) error {
	var a string
	switch obj.(type) {
	case nil:
		a = "null"
	case bool:
		a = "boolean"
	case int, int8, int16, int32, int64:
		a = "integer"
	case float32, float64:
		a = "number"
	default:
		switch reflect.ValueOf(obj).Kind() {
		case reflect.Slice:
			a = "array"
		case reflect.Map:
			a = "object"
		default:
			return fmt.Errorf("InvalidType: %+v", obj)
		}
	}
	if a != def {
		return TypeError{v, a, def}
	}
	return nil
}

type TypeError struct {
	Validator          Validator
	Object, Definition string
}

func (err TypeError) Error() string {
	return fmt.Sprintf("%s must be %s, but `%s`", err.Validator, err.Definition, err.Object)
}
