package validator_test

import (
	"testing"

	"github.com/minodisk/go-jsonschema/validator"
)

func TestMultipleOf(t *testing.T) {
	var err error
	v := validator.Validator{}

	err = v.MultipleOf(0, 3)
	if err != nil {
		t.Error(err)
	}
	err = v.MultipleOf(1, 3)
	if err == nil {
		t.Error(err)
	}
	err = v.MultipleOf(2, 3)
	if err == nil {
		t.Error(err)
	}
	err = v.MultipleOf(3, 3)
	if err != nil {
		t.Error(err)
	}
	err = v.MultipleOf(4, 3)
	if err == nil {
		t.Error(err)
	}
	err = v.MultipleOf(5, 3)
	if err == nil {
		t.Error(err)
	}
	err = v.MultipleOf(6, 3)
	if err != nil {
		t.Error(err)
	}
}

func TestMaximum(t *testing.T) {
	var err error
	v := validator.Validator{}

	err = v.Maximum(2, 3, false)
	if err != nil {
		t.Error(err)
	}
	err = v.Maximum(3, 3, false)
	if err != nil {
		t.Error(err)
	}
	err = v.Maximum(4, 3, false)
	if err == nil {
		t.Error(err)
	}

	err = v.Maximum(2, 3, true)
	if err != nil {
		t.Error(err)
	}
	err = v.Maximum(3, 3, true)
	if err == nil {
		t.Error(err)
	}
	err = v.Maximum(4, 3, true)
	if err == nil {
		t.Error(err)
	}
}

func TestMinimum(t *testing.T) {
	var err error
	v := validator.Validator{}

	err = v.Minimum(2, 3, false)
	if err == nil {
		t.Error(err)
	}
	err = v.Minimum(3, 3, false)
	if err != nil {
		t.Error(err)
	}
	err = v.Minimum(4, 3, false)
	if err != nil {
		t.Error(err)
	}

	err = v.Minimum(2, 3, true)
	if err == nil {
		t.Error(err)
	}
	err = v.Minimum(3, 3, true)
	if err == nil {
		t.Error(err)
	}
	err = v.Minimum(4, 3, true)
	if err != nil {
		t.Error(err)
	}
}
