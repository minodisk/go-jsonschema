package validator_test

import (
	"testing"

	"github.com/minodisk/go-jsonschema/validator"
)

func TestMaxLength(t *testing.T) {
	var err error
	v := validator.Validator{}

	err = v.MaxLength("ab", 3)
	if err != nil {
		t.Error(err)
	}
	err = v.MaxLength("abc", 3)
	if err != nil {
		t.Error(err)
	}
	err = v.MaxLength("abcd", 3)
	if err == nil {
		t.Errorf("should return max length error")
	}
	err = v.MaxLength("あい", 3)
	if err != nil {
		t.Error(err)
	}
	err = v.MaxLength("あいう", 3)
	if err != nil {
		t.Error(err)
	}
	err = v.MaxLength("あいうえ", 3)
	if err == nil {
		t.Errorf("should return max length error")
	}
}

func TestMinLength(t *testing.T) {
	var err error
	v := validator.Validator{}

	err = v.MinLength("ab", 3)
	if err == nil {
		t.Errorf("should return min length error")
	}
	err = v.MinLength("abc", 3)
	if err != nil {
		t.Error(err)
	}
	err = v.MinLength("abcd", 3)
	if err != nil {
		t.Error(err)
	}
	err = v.MinLength("あい", 3)
	if err == nil {
		t.Errorf("should return min length error")
	}
	err = v.MinLength("あいう", 3)
	if err != nil {
		t.Error(err)
	}
	err = v.MinLength("あいうえ", 3)
	if err != nil {
		t.Error(err)
	}
}

func TestPattern(t *testing.T) {
	var err error
	v := validator.Validator{}

	err = v.Pattern("555-1212", "^(\\([0-9]{3}\\))?[0-9]{3}-[0-9]{4}$")
	if err != nil {
		t.Error(err)
	}
	err = v.Pattern("(888)555-1212", "^(\\([0-9]{3}\\))?[0-9]{3}-[0-9]{4}$")
	if err != nil {
		t.Error(err)
	}
	err = v.Pattern("(888)555-1212 ext. 532", "^(\\([0-9]{3}\\))?[0-9]{3}-[0-9]{4}$")
	if err == nil {
		t.Errorf("should return pattern error")
	}
	err = v.Pattern("(800)FLOWERS", "^(\\([0-9]{3}\\))?[0-9]{3}-[0-9]{4}$")
	if err == nil {
		t.Errorf("should return pattern error")
	}
}
