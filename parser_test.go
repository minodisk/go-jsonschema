package gojsa_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/minodisk/gojsa"
)

func TestItems(t *testing.T) {
	json := `{
		"properties": {
			"foo": {
				"items": {
					"id": "foo-1"
				}
			},
			"bar": {
				"items": [
					{
						"id": "bar-1"
					},
					{
						"id": "bar-2"
					}
				]
			}
		}
	}`
	s, err := gojsa.Parse(bytes.NewBufferString(json))
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", s)
	if !s.Properties["foo"].Items.IsSignle() {
		t.Errorf("foo.Items should be single")
	}
	if a := s.Properties["foo"].Items[0].ID; a != "foo-1" {
		t.Error(a)
	}
	if s.Properties["bar"].Items.IsSignle() {
		t.Errorf("bar.Items shouldn't be single")
	}
	if a := s.Properties["bar"].Items[0].ID; a != "bar-1" {
		t.Error(a)
	}
	if a := s.Properties["bar"].Items[1].ID; a != "bar-2" {
		t.Error(a)
	}
}

func TestAdditionalItems(t *testing.T) {
	json := `{
		"properties": {
			"foo": {
				"additionalItems": {
					"id": "foo-1"
				}
			},
			"bar": {
				"additionalItems": false
			}
		}
	}`
	s, err := gojsa.Parse(bytes.NewBufferString(json))
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", s)
	if a := s.Properties["foo"].Items[0].ID; a != "foo-1" {
		t.Errorf("foo.Type[0] is'nt number but %+v", a)
	}
	if a := s.Properties["bar"].Items[0].ID; a != "bar-1" {
		t.Errorf("foo.Type[0] is'nt number but %+v", a)
	}
	if a := s.Properties["bar"].Items[1].ID; a != "bar-2" {
		t.Errorf("foo.Type[0] is'nt boolean but %+v", a)
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
	log.Printf("%+v", schema)
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
