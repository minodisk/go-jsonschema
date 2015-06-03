package gojsa

import (
	"io"
	"text/template"
)

func Doc(w io.Writer, root *Root, t *template.Template) error {
	return t.Execute(w, root)
}
