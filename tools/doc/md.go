package doc

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/minodisk/go-jsonschema"
	"github.com/minodisk/go-jsonschema/tools/combine"
	"github.com/minodisk/go-jsonschema/tools/encoding"
	"github.com/minodisk/go-jsonschema/tools/watcher"
)

type Engine string

const (
	TextTemplateEngine Engine = "text/template"
)

var (
	rHashIgnorance = regexp.MustCompile(`[:\/]`)
	funcMap        = template.FuncMap{
		"urlHash": urlHash,
	}
)

type Options struct {
	Template string
	Engine   Engine
	Output   string
	Format   string
	IsWatch  bool
	Meta     string
	Input    string
}

func Generate(o Options) error {
	if err := generate(o); err != nil {
		return err
	}
	if o.IsWatch {
		return watcher.Watch([]string{o.Input, o.Template}, func(filename string) {
			if err := generate(o); err != nil {
				log.Printf("fail to generate: %s", err)
			}
		})
	}
	return nil
}

func generate(o Options) error {
	mode, err := fileMode(o.Input)
	if err != nil {
		return err
	}
	var src []byte
	var enc encoding.Encoding
	if mode.IsDir() {
		enc = encoding.JSONEncoding
		src, err = combine.Combine(o.Input, o.Meta, enc)
		if err != nil {
			return err
		}
		log.Printf("[doc] read schema files in '%s'", o.Input)
	} else {
		enc, err = encoding.NewWithFilename(o.Input)
		if err != nil {
			return err
		}
		src, err = ioutil.ReadFile(o.Input)
		if err != nil {
			return err
		}
		log.Printf("[doc] read schema file '%s'", o.Input)
	}

	switch enc {
	default:
		return fmt.Errorf("unsupported encoding '%s'", enc)
	case encoding.JSONEncoding:
	case encoding.YAMLEncoding:
		src, err = encoding.YAMLToJSON(src)
		if err != nil {
			return err
		}
	}

	schema, err := jsonschema.New(src)
	if err != nil {
		return err
	}

	tmpl, err := ioutil.ReadFile(o.Template)
	if err != nil {
		return err
	}
	template := string(tmpl)
	log.Printf("[doc] read template file: %s", o.Template)

	var buf bytes.Buffer
	switch o.Engine {
	case TextTemplateEngine:
		if err := Markdown(&buf, schema, "tmp", template); err != nil {
			return err
		}
		if err := ioutil.WriteFile(o.Output, buf.Bytes(), 0644); err != nil {
			return err
		}
		log.Printf("[doc] write document file: %s", o.Output)
	default:
		return fmt.Errorf("unsupported engine %s", string(o.Engine))
	}

	return nil
}

func fileMode(p string) (*os.FileMode, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	s, err := f.Stat()
	if err != nil {
		return nil, err
	}
	m := s.Mode()
	return &m, nil
}

func Markdown(w io.Writer, s *jsonschema.Schema, name, text string) error {
	tmpl := template.Must(template.New(name).Funcs(funcMap).Parse(text))
	return tmpl.Execute(w, s)
}

func urlHash(str string) string {
	u := strings.Replace(string(str), " ", "-", -1)
	u = rHashIgnorance.ReplaceAllString(u, "")
	u = strings.ToLower(u)
	return fmt.Sprintf("#%s", u)
}
