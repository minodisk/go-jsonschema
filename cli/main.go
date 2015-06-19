package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/minodisk/go-jsonschema/cli/doc"
)

const (
	version = "0.0.0"
)

func main() {
	app := cli.NewApp()
	app.Name = "jsonschema"
	app.Usage = "Tools for JSON Schema"

	app.Commands = []cli.Command{
		{
			Name:  "doc",
			Usage: "generate docuemnt from schema.json",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "encoding",
					Usage: "input encoding",
				},
				cli.StringFlag{
					Name:  "output, o",
					Value: "schema.md",
					Usage: "output filename",
				},
				cli.StringFlag{
					Name:  "template, t",
					Value: "schema.md.tmpl",
					Usage: "template filename",
				},
				cli.StringFlag{
					Name:  "engine, e",
					Value: "text/template",
					Usage: "template engine",
				},
				cli.StringFlag{
					Name:  "format, f",
					Value: "markdown",
					Usage: "the format of output document",
				},
				cli.BoolFlag{
					Name:  "watch, w",
					Usage: "run in changing related files",
				},
			},
			Action: func(c *cli.Context) {
				if err := doc.Generate(doc.Options{
					Input:    filepath.Clean(c.Args()[0]),
					Encoding: doc.Encoding(c.String("encoding")),
					Template: filepath.Clean(c.String("template")),
					Engine:   doc.Engine(c.String("engine")),
					Output:   filepath.Clean(c.String("output")),
					Format:   c.String("format"),
					IsWatch:  c.Bool("watch"),
				}); err != nil {
					log.Println(err)
				}
			},
		},
	}

	app.Run(os.Args)
}
