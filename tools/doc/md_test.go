package doc_test

import (
	"bytes"
	"io/ioutil"
	"testing"
	"text/template"

	"github.com/minodisk/gojsa"
	"github.com/minodisk/gojsa/tools/doc"
)

func TestMarkdown(t *testing.T) {
	// y, err := ioutil.ReadFile("../../fixtures/schema.yml")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// s, err := gojsa.NewYAML(y)
	j, err := ioutil.ReadFile("../../fixtures/schema.json")
	if err != nil {
		t.Fatal(err)
	}
	s, err := gojsa.NewJSON(j)
	if err != nil {
		t.Fatal(err)
	}

	tplSrc, err := ioutil.ReadFile("../../fixtures/schema.md.tmpl")
	if err != nil {
		t.Fatal(err)
	}
	tpl := template.Must(template.New("md").Parse(string(tplSrc)))

	buf := bytes.NewBuffer(nil)
	if err := doc.Markdown(buf, s, tpl); err != nil {
		t.Fatal(err)
	}

	// fmt.Println(string(buf.Bytes()))

	ioutil.WriteFile("../../fixtures/schema.go.md", buf.Bytes(), 0644)
}
