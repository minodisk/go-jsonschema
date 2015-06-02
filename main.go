package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/minodisk/go-json-schema/doc"
	"github.com/minodisk/go-json-schema/parser"
)

func main() {
	file, err := os.Open("heroku.schema.json")
	if err != nil {
		panic(err)
	}
	r := parser.Parse(file)
	log.Printf("%+v", r)

	buf, err := ioutil.ReadFile("schema.tpl.md")
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("mytemplate").Parse(string(buf)))
	doc.Doc(os.Stdout, &r, t)
}
