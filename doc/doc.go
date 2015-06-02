package doc

import (
	"io"
	"text/template"

	"github.com/minodisk/go-json-schema/schema"
)

func Doc(w io.Writer, root *schema.Root, t *template.Template) error {
	return t.Execute(w, root)
}
