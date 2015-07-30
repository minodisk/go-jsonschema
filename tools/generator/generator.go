package generator

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"text/template"

	"github.com/minodisk/go-jsonschema"
	"github.com/serenize/snaker"
)

var (
	funcMap = template.FuncMap{
		"snakeToCamel": snaker.SnakeToCamel,
		"camelToSnake": snaker.CamelToSnake,
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
	t, err := getTemplate(o.Template)
	if err != nil {
		return err
	}
	code, err := Generate(s, t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(o.Output, code, 0644)
}

func getTemplate(src string) (buf []byte, err error) {
	if src == "" {
		log.Println("[generate] use defualt template")
		return Asset("fixtures/schema.go.tmpl")
	}
	return ioutil.ReadFile(src)
}

func Generate(schema, tmpl []byte) (code []byte, err error) {
	s, err := jsonschema.New(schema)
	if err != nil {
		return nil, err
	}
	t := template.Must(template.New("").Funcs(funcMap).Parse(string(tmpl)))
	buf := bytes.NewBuffer([]byte{})
	err = t.Execute(buf, s)
	if err != nil {
		return nil, err
	}
	return format.Source(buf.Bytes())
}
