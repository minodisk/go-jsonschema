package doc

import (
	"io"
	"text/template"

	"github.com/minodisk/gojsa"
)

func Markdown(w io.Writer, s *gojsa.Schema, t *template.Template) error {
	return t.Execute(w, s)
}
