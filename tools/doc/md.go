package doc

import (
	"fmt"
	"io"
	"regexp"
	"strings"
	"text/template"

	"github.com/minodisk/gojsa"
)

var (
	rHashIgnorance = regexp.MustCompile(`[:\/]`)
	funcMap        = template.FuncMap{
		"urlHash": urlHash,
	}
)

func Markdown(w io.Writer, s *gojsa.Schema, name, text string) error {
	tmpl := template.Must(template.New(name).Funcs(funcMap).Parse(text))
	// if err != nil {
	// 	return err
	// }
	return tmpl.Execute(w, s)
}

// func YamlToMarkdown(w io.Writer, y []byte, t *ego.Template) error {
// 	t.Write(w)
// }

func urlHash(str string) string {
	u := strings.Replace(string(str), " ", "-", -1)
	u = rHashIgnorance.ReplaceAllString(u, "")
	u = strings.ToLower(u)
	return fmt.Sprintf("#%s", u)
}
