package gojsa_test

import (
	"bytes"
	"testing"

	"github.com/minodisk/gojsa"
)

func TestParse(t *testing.T) {
	json := `{
		"definitions": {
			"foo": {
				"type": ["object"],
				"id": "foo",
				"title": "This is Foo schema"
			}
		},
		"properties": {
			"foo": {
				"$ref": "#/definitions/foo"
			}
		}
	}`
	buf := bytes.NewBufferString(json)
	gojsa.Parse(buf)
	buf = bytes.NewBufferString(json)
	gojsa.Parse(buf)

	// btpl, err := ioutil.ReadFile("schema.tpl.md")
	// if err != nil {
	// 	panic(err)
	// }
	// tpl := template.Must(template.New("mytemplate").Parse(string(btpl)))
	// gojsa.Doc(os.Stdout, &r, tpl)
}
