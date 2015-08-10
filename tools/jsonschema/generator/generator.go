//go:generate go-bindata -pkg generator -o assets.go assets

package generator

import (
	"bytes"
	"io/ioutil"
	"path"
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

type IO struct {
	Input  string
	Output string
}

func Run(o Options) error {
	schema, err := ioutil.ReadFile(o.Input)
	if err != nil {
		return err
	}

	for _, io := range []IO{
		IO{"assets/routing.go.tmpl", "routing.go"},
		IO{"assets/struct.go.tmpl", "struct.go"},
		IO{"assets/validator.go.tmpl", "validator.go"},
	} {
		buf, err := Asset(io.Input)
		if err != nil {
			return err
		}
		filename := path.Join(o.Output, io.Output)
		code, err := Generate(schema, buf, filename)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filename, code, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func Generate(schema, tmpl []byte, filename string) ([]byte, error) {
	s, err := jsonschema.New(schema)
	if err != nil {
		return nil, err
	}

	t := template.Must(template.New("").Funcs(funcMap).Parse(string(tmpl)))
	buf := bytes.NewBuffer([]byte{})

	err = t.Execute(buf, map[string]interface{}{
		"package": "main",
		"schema":  s,
	})
	if err != nil {
		return nil, err
	}

	// return buf.Bytes(), nil
	return imports.Process(filename, buf.Bytes(), nil)
}
