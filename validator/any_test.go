package validator_test

import (
	"testing"

	"github.com/minodisk/go-jsonschema/validator"
)

func TestEnum(t *testing.T) {
	var err error
	v := validator.Validator{"Foo", "Bar"}

	err = v.Enum("a", []interface{}{"a", "b", 3})
	if err != nil {
		t.Error(err)
	}
	err = v.Enum(3, []interface{}{"a", "b", 3})
	if err != nil {
		t.Error(err)
	}
	err = v.Enum("c", []interface{}{"a", "b", 3})
	if err == nil {
		t.Error(err)
	}
	err = v.Enum(3.14, []interface{}{"a", "b", 3})
	if err == nil {
		t.Error(err)
	}
	err = v.Enum(true, []interface{}{"a", "b", 3})
	if err == nil {
		t.Error(err)
	}
}

func TestType(t *testing.T) {
	var err error
	v := validator.Validator{}

	for _, val := range []interface{}{nil} {
		err = v.Type(val, []string{"null"})
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{nil} {
		err = v.Type(val, []string{"null", "boolean"})
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{nil} {
		err = v.Type(val, []string{"boolean"})
		if err == nil {
			t.Error(err)
		}
	}

	for _, val := range []interface{}{true, false} {
		err = v.Type(val, []string{"boolean"})
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{true, false} {
		err = v.Type(val, []string{"boolean", "integer"})
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{true, false} {
		err = v.Type(val, []string{"integer"})
		if err == nil {
			t.Error(err)
		}
	}

	for _, val := range []interface{}{3, int8(3), int16(3), int32(3), int64(3)} {
		err = v.Type(val, []string{"integer"})
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{3, int8(3), int16(3), int32(3), int64(3)} {
		err = v.Type(val, []string{"integer", "number"})
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{3, int8(3), int16(3), int32(3), int64(3)} {
		err = v.Type(val, []string{"number"})
		if err == nil {
			t.Error(err)
		}
	}

	for _, val := range []interface{}{3.14, float32(3.14), float64(3.14), float32(3), float64(3)} {
		err = v.Type(val, []string{"number"})
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{3.14, float32(3.14), float64(3.14), float32(3), float64(3)} {
		err = v.Type(val, []string{"number", "array"})
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{3.14, float32(3.14), float64(3.14), float32(3), float64(3)} {
		err = v.Type(val, []string{"array"})
		if err == nil {
			t.Error(err)
		}
	}

	for _, val := range []interface{}{[]int{1, 2, 3}, []string{"a", "b", "c"}} {
		err = v.Type(val, []string{"array"})
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{[]int{1, 2, 3}, []string{"a", "b", "c"}} {
		err = v.Type(val, []string{"array", "object"})
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{[]int{1, 2, 3}, []string{"a", "b", "c"}} {
		err = v.Type(val, []string{"object"})
		if err == nil {
			t.Error(err)
		}
	}

	for _, val := range []interface{}{map[string]int{"a": 1, "b": 2, "c": 3}} {
		err = v.Type(val, []string{"object"})
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{map[string]int{"a": 1, "b": 2, "c": 3}} {
		err = v.Type(val, []string{"object", "null"})
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{map[string]int{"a": 1, "b": 2, "c": 3}} {
		err = v.Type(val, []string{"null"})
		if err == nil {
			t.Error(err)
		}
	}
}
