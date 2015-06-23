package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/minodisk/go-jsonschema/tools/combine"
	"github.com/minodisk/go-jsonschema/tools/doc"
	"github.com/minodisk/go-jsonschema/tools/encoding"
)

const (
	version = "0.0.0"
)

func main() {
	app := cli.NewApp()
	app.Name = "jsonschema"
	app.Usage = "Tools for JSON Schema"

	watch := cli.BoolFlag{
		Name:  "watch, w",
		Usage: "watch the files and run when the files are changed",
	}
	meta := cli.StringFlag{
		Name:  "meta, m",
		Value: "meta.yml",
		Usage: "meta file",
	}

	app.Commands = []cli.Command{
		{
			Name:  "doc",
			Usage: "generate docuemnt from schema.json",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output, o",
					Value: "schema.md",
					Usage: "the path of the output file",
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
				watch,
				meta,
			},
			Action: func(c *cli.Context) {
				if err := doc.Generate(doc.Options{
					Template: filepath.Clean(c.String("template")),
					Engine:   doc.Engine(c.String("engine")),
					Output:   filepath.Clean(c.String("output")),
					Format:   c.String("format"),
					IsWatch:  c.Bool("watch"),
					Meta:     filepath.Clean(c.String("meta")),
					Input:    filepath.Clean(c.Args()[0]),
				}); err != nil {
					log.Println(err)
				}
			},
		},

		{
			Name:  "combine",
			Usage: "combine partial schema files into one schema file",
			Flags: []cli.Flag{
				meta,
				cli.StringFlag{
					Name:  "output, o",
					Value: "schema.json",
					Usage: "the path of the output file",
				},
				cli.StringFlag{
					Name:  "encoding, e",
					Value: "json",
					Usage: "the encoding of output file",
				},
				watch,
			},
			Action: func(c *cli.Context) {
				if err := combine.Run(combine.Options{
					Input:    filepath.Clean(c.Args()[0]),
					Meta:     filepath.Clean(c.String("meta")),
					Output:   filepath.Clean(c.String("output")),
					Encoding: encoding.Encoding(c.String("encoding")),
				}); err != nil {
					log.Println(err)
				}
			},
		},
	}

	app.Run(os.Args)
}
