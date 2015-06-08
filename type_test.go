package gojsa_test

import (
	"encoding/json"
	"testing"

	"github.com/minodisk/gojsa"
)

func TestTypeInteger(t *testing.T) {
	s := new(gojsa.Schema)
	js := `{
		"type": "integer"
	}`
	if err := json.Unmarshal([]byte(js), s); err != nil {
		t.Fatal(err)
	}

	var i interface{}

	if err := json.Unmarshal([]byte(`42`), &i); err != nil {
		t.Fatal(err)
	}
	if err := s.Validate(i); err != nil {
		t.Errorf("should be valid, but %s", err)
	}

	if err := json.Unmarshal([]byte(`-1`), &i); err != nil {
		t.Fatal(err)
	}
	if err := s.Validate(i); err != nil {
		t.Errorf("should be valid, but %s", err)
	}

	if err := json.Unmarshal([]byte(`3.1415926`), &i); err != nil {
		t.Fatal(err)
	}
	if err := s.Validate(i); err == nil {
		t.Errorf("should be invalid, but %s", err)
	}

	if err := json.Unmarshal([]byte(`"42"`), &i); err != nil {
		t.Fatal(err)
	}
	if err := s.Validate(i); err == nil {
		t.Errorf("should be invalid, but %s", err)
	}
}
