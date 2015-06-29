package jsonschema_test

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/minodisk/go-jsonschema"
)

func TestSchema(t *testing.T) {
	b, err := ioutil.ReadFile("fixtures/schema.json")
	if err != nil {
		t.Fatal(err)
	}

	schema, err := jsonschema.New(b)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%+v", schema.Properties.Schemas["album"])
}
