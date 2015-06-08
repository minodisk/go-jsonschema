package gojsa_test

import (
	"bytes"
	"testing"

	"github.com/minodisk/gojsa"
)

func TestSchemaOrBool(t *testing.T) {
	var json string
	var s gojsa.Schema
	var err error

	json = `{
		"additionalItems": {
			"id": "foo"
		}
	}`
	s, err = gojsa.Parse(bytes.NewBufferString(json))
	if err != nil {
		t.Fatal(err)
	}
	if !s.AdditionalItems.IsSchema {
		t.Errorf("should be schema")
	}
	if a := s.AdditionalItems.Schema.ID; a != "foo" {
		t.Errorf("id should be expected foo, but actual %s", a)
	}

	json = `{
		"additionalItems": true
	}`
	s, err = gojsa.Parse(bytes.NewBufferString(json))
	if err != nil {
		t.Fatal(err)
	}
	if s.AdditionalItems.IsSchema {
		t.Errorf("shouldn't be Schema")
	}
	if !s.AdditionalItems.Bool {
		t.Errorf("should be true")
	}

	json = `{
		"additionalItems": false
	}`
	s, err = gojsa.Parse(bytes.NewBufferString(json))
	if err != nil {
		t.Fatal(err)
	}
	if s.AdditionalItems.IsSchema {
		t.Errorf("when bool, shouldn't be schema")
	}
	if s.AdditionalItems.Bool {
		t.Errorf("should be false")
	}
}

func TestSchemaOrSchemas(t *testing.T) {
	var json string
	var s gojsa.Schema
	var err error

	json = `{
		"items": {
			"id": "foo"
		}
	}`
	s, err = gojsa.Parse(bytes.NewBufferString(json))
	if err != nil {
		t.Fatal(err)
	}
	if !s.Items.IsSchema {
		t.Fatal("should be Schema")
	}
	if a := s.Items.Schema.ID; a != "foo" {
		t.Errorf("id is expected foo, but actual %s", a)
	}

	json = `{
		"items": [
			{
				"id": "foo"
			},
			{
				"id": "bar"
			}
		]
	}`
	s, err = gojsa.Parse(bytes.NewBufferString(json))
	if err != nil {
		t.Fatal(err)
	}
	if s.Items.IsSchema {
		t.Fatal("shouldn't be Schema")
	}
	if a := s.Items.Schemas[0].ID; a != "foo" {
		t.Errorf("id is expected foo, but actual %s", a)
	}
	if a := s.Items.Schemas[1].ID; a != "bar" {
		t.Errorf("id is expected foo, but actual %s", a)
	}
}

func TestStrings(t *testing.T) {
	json := `{
		"properties": {
			"foo": {
				"type": "number"
			},
			"bar": {
				"type": ["number", "boolean"]
			}
		}
	}`
	buf := bytes.NewBufferString(json)
	schema, err := gojsa.Parse(buf)
	if err != nil {
		t.Fatal(err)
	}
	if a := schema.Properties["foo"].Type[0]; a != "number" {
		t.Errorf("foo.Type[0] is'nt number but %+v", a)
	}
	if a := schema.Properties["bar"].Type[0]; a != "number" {
		t.Errorf("foo.Type[0] is'nt number but %+v", a)
	}
	if a := schema.Properties["bar"].Type[1]; a != "boolean" {
		t.Errorf("foo.Type[0] is'nt boolean but %+v", a)
	}
}

// func TestParse(t *testing.T) {
// 	json := `{
// 		"definitions": {
// 			"foo": {
// 				"type": ["object"],
// 				"id": "foo",
// 				"title": "This is Foo schema"
// 			}
// 		},
// 		"properties": {
// 			"foo": {
// 				"$ref": "#/definitions/foo"
// 			}
// 		}
// 	}`
// 	buf := bytes.NewBufferString(json)
// 	gojsa.Parse(buf)
// 	buf = bytes.NewBufferString(json)
// 	gojsa.Parse(buf)
//
// 	// btpl, err := ioutil.ReadFile("schema.tpl.md")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// tpl := template.Must(template.New("mytemplate").Parse(string(btpl)))
// 	// gojsa.Doc(os.Stdout, &r, tpl)
// }
