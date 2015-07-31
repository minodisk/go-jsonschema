//go:generate go-bindata -pkg generator -o assets.go assets

package generator

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"

	"github.com/minodisk/go-jsonschema"
	"github.com/serenize/snaker"
)

var (
	funcMap = template.FuncMap{
		"snakeToCamel": snaker.SnakeToCamel,
		"camelToSnake": snaker.CamelToSnake,
		"firstChar": func(s string) string {
			return strings.ToLower(string(s[0]))
		},
	}
)

type Options struct {
	Input    string
	Template string
	Output   string
}

func Run(o Options) (err error) {
	s, err := ioutil.ReadFile(o.Input)
	if err != nil {
		return err
	}

	var t []byte
	if o.Template == "" {
		log.Println("[generate] use defualt template")
		t, err = Asset("assets/routing.go.tmpl")
	} else {
		t, err = ioutil.ReadFile(o.Template)
	}
	if err != nil {
		return err
	}

	v, err := Asset("assets/validator.go")
	if err != nil {
		return err
	}

	code, err := Generate(s, t, v, o.Output)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(o.Output, code, 0644)
}

func Generate(schema, tmpl, v []byte, filename string) (code []byte, err error) {
	s, err := jsonschema.New(schema)
	if err != nil {
		return nil, err
	}
	t := template.Must(template.New("").Funcs(funcMap).Parse(string(tmpl)))
	buf := bytes.NewBuffer([]byte{})
	err = t.Execute(buf, map[string]interface{}{
		"package":       "main",
		"schema":        s,
		"validatorCode": strings.Replace(strings.Replace(string(v), "package dummy", "", 1), "import \"fmt\"", "", -1),
	})
	if err != nil {
		return nil, err
	}
	return imports.Process(filename, buf.Bytes(), nil)
}
