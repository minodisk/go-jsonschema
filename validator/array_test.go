package validator_test

import (
	"testing"

	"github.com/minodisk/go-jsonschema/validator"
)

func TestMaxItems(t *testing.T) {
	var err error
	v := validator.Validator{}

	err = v.MaxItems([]string{"a", "b"}, 3)
	if err != nil {
		t.Error(err)
	}
	err = v.MaxItems([]string{"a", "b", "c"}, 3)
	if err != nil {
		t.Error(err)
	}
	err = v.MaxItems([]string{"a", "b", "c", "d"}, 3)
	if err == nil {
		t.Error(err)
	}
}

func TestMinItems(t *testing.T) {
	var err error
	v := validator.Validator{}

	err = v.MinItems([]string{"a", "b"}, 3)
	if err == nil {
		t.Error(err)
	}
	err = v.MinItems([]string{"a", "b", "c"}, 3)
	if err != nil {
		t.Error(err)
	}
	err = v.MinItems([]string{"a", "b", "c", "d"}, 3)
	if err != nil {
		t.Error(err)
	}
}

func TestUniqueItems(t *testing.T) {
	var err error
	v := validator.Validator{}

	err = v.UniqueItems([]string{})
	if err != nil {
		t.Error(err)
	}
	err = v.UniqueItems([]string{"a", "b", "c"})
	if err != nil {
		t.Error(err)
	}
	err = v.UniqueItems([]string{"a", "b", "c", "a"})
	if err == nil {
		t.Error(err)
	}
	err = v.UniqueItems([]string{"a", "b", "c", "a", "b"})
	if err == nil {
		t.Error(err)
	}
	err = v.UniqueItems([]string{"a", "b", "c", "b", "b"})
	if err == nil {
		t.Error(err)
	}
}
