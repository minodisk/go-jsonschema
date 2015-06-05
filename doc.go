package gojsa

import (
	"io"
	"text/template"
)

func Doc(w io.Writer, s *Schema, t *template.Template) error {
	return t.Execute(w, s)
}
