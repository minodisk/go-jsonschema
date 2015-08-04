package validator_test

import (
	"testing"

	"github.com/minodisk/go-jsonschema/validator"
)

func TestMaxProperties(t *testing.T) {
	var err error
	v := validator.Validator{}

	err = v.MaxProperties(map[string]int{
		"a": 100,
		"b": 200,
	}, 3)
	if err != nil {
		t.Error(err)
	}
	err = v.MaxProperties(map[string]int{
		"a": 100,
		"b": 200,
		"c": 300,
	}, 3)
	if err != nil {
		t.Error(err)
	}
	err = v.MaxProperties(map[string]int{
		"a": 100,
		"b": 200,
		"c": 300,
		"d": 400,
	}, 3)
	if err == nil {
		t.Errorf("should return max properties error")
	}
}

func TestMinProperties(t *testing.T) {
	var err error
	v := validator.Validator{}

	err = v.MinProperties(map[string]int{
		"a": 100,
		"b": 200,
	}, 3)
	if err == nil {
		t.Errorf("should return min properties error")
	}
	err = v.MinProperties(map[string]int{
		"a": 100,
		"b": 200,
		"c": 300,
	}, 3)
	if err != nil {
		t.Error(err)
	}
	err = v.MinProperties(map[string]int{
		"a": 100,
		"b": 200,
		"c": 300,
		"d": 400,
	}, 3)
	if err != nil {
		t.Error(err)
	}
}

func TestRequired(t *testing.T) {
	var err error
	v := validator.Validator{}

	err = v.Required(map[string]int{
		"a": 100,
		"b": 200,
	}, []string{"a", "b", "c"})
	if err == nil {
		t.Error(err)
	}
	err = v.Required(map[string]int{
		"a": 100,
		"b": 200,
		"c": 300,
	}, []string{"a", "b", "c"})
	if err != nil {
		t.Error(err)
	}
	err = v.Required(map[string]int{
		"a": 100,
		"b": 200,
		"c": 300,
		"d": 400,
	}, []string{"a", "b", "c"})
	if err != nil {
		t.Error(err)
	}
}

func TestDependenciesStrings(t *testing.T) {
	var err error
	v := validator.Validator{}

	err = v.DependenciesStrings(map[string]int{
		"a": 100,
		"b": 200,
	}, map[string][]string{
		"c": []string{"d"},
	})
	if err != nil {
		t.Error(err)
	}
	err = v.DependenciesStrings(map[string]int{
		"a": 100,
		"b": 200,
		"c": 300,
	}, map[string][]string{
		"c": []string{"d"},
	})
	if err == nil {
		t.Errorf("should return dependencies error")
	}
	err = v.DependenciesStrings(map[string]int{
		"a": 100,
		"b": 200,
		"c": 300,
		"d": 400,
	}, map[string][]string{
		"c": []string{"d"},
	})
	if err != nil {
		t.Error(err)
	}
	err = v.DependenciesStrings(map[string]int{
		"a": 100,
		"b": 200,
		"d": 400,
	}, map[string][]string{
		"c": []string{"d"},
	})
	if err != nil {
		t.Error(err)
	}
}
