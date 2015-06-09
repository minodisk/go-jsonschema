package doc_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/minodisk/gojsa"
	"github.com/minodisk/gojsa/tools/doc"
)

func TestMarkdown(t *testing.T) {
	// 	text := `## {{.Title}}
	// {{.Description}}
	//
	// ###
	//
	// {{}}`
	// 	tpl := template.Must(template.New("md").Parse(text))
	// if err := json.Unmarshal([]byte(j), s); err != nil {
	// 	t.Fatal(err)
	// }

	b, err := ioutil.ReadFile("../../fixtures/schema.yml")
	if err != nil {
		t.Fatal(err)
	}
	s := new(gojsa.Schema)
	if err := yaml.Unmarshal(b, s); err != nil {
		t.Fatal(err)
	}

	log.Printf("%+v", s)

	buf := bytes.NewBuffer(nil)
	doc.Markdown(buf, s, tpl)
	a := string(buf.Bytes())
	e := `this is title
	this is description`
	if a != e {
		t.Errorf("expected %s, but actual %s", e, a)
	}
}
