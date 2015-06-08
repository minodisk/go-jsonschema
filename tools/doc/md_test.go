package doc_test

import (
	"encoding/json"
	"os"
	"testing"
	"text/template"

	"github.com/minodisk/gojsa"
	"github.com/minodisk/gojsa/tools/doc"
)

func TestMarkdown(t *testing.T) {
	text := `{{.Title}}
	{{.Description}}`
	tpl := template.Must(template.New("md").Parse(text))

	j := `{
		"title": "this is title",
		"description":"this is description",
	}`
	s := new(gojsa.Schema)
	if err := json.Unmarshal([]byte(j), s); err != nil {
		t.Fatal(err)
	}

	doc.Markdown(os.Stdout, s, tpl)
}
