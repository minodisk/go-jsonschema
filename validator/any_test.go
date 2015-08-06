package validator_test

import (
	"encoding/json"
	"log"
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
	buf, err := json.MarshalIndent(err, "", "  ")
	if err != nil {
		t.Error(err)
	}
	log.Println(string(buf))
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
		err = v.Type(val, "null")
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{true, false} {
		err = v.Type(val, "boolean")
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{3, int8(3), int16(3), int32(3), int64(3)} {
		err = v.Type(val, "integer")
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{3.14, float32(3.14), float64(3.14), float32(3), float64(3)} {
		err = v.Type(val, "number")
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{[]int{1, 2, 3}, []string{"a", "b", "c"}} {
		err = v.Type(val, "array")
		if err != nil {
			t.Error(err)
		}
	}
	for _, val := range []interface{}{map[string]int{"a": 1, "b": 2, "c": 3}} {
		err = v.Type(val, "object")
		if err != nil {
			t.Error(err)
		}
	}
}
