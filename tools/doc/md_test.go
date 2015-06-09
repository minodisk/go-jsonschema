package doc_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
	"text/template"

	"gopkg.in/yaml.v2"

	"github.com/minodisk/gojsa"
	"github.com/minodisk/gojsa/tools/doc"
)

func TestMarkdown(t *testing.T) {
	y, err := ioutil.ReadFile("../../fixtures/schema.yml")
	if err != nil {
		t.Fatal(err)
	}
	s := new(gojsa.Schema)
	if err := yaml.Unmarshal(y, s); err != nil {
		t.Fatal(err)
	}

	tplSrc, err := ioutil.ReadFile("../../fixtures/schema.md.tmpl")
	if err != nil {
		t.Fatal(err)
	}
	tpl := template.Must(template.New("md").Parse(string(tplSrc)))

	buf := bytes.NewBuffer(nil)
	doc.Markdown(buf, s, tpl)

	fmt.Println(string(buf.Bytes()))
	// a := string(buf.Bytes())
	// e := `this is title
	// this is description`
	// if a != e {
	// 	t.Errorf("expected %s, but actual %s", e, a)
	// }

	// b, err := ioutil.ReadFile("../../fixtures/schema.yml")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// s := new(gojsa.Schema)
	// if err := yaml.Unmarshal(b, s); err != nil {
	// 	t.Fatal(err)
	// }
	//
	// log.Printf("%+v", s)
}
