package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/minodisk/go-jsonschema/tools/combine"
	"github.com/minodisk/go-jsonschema/tools/doc"
	"github.com/minodisk/go-jsonschema/tools/encoding"
	"github.com/minodisk/go-jsonschema/tools/generator"
)

//go:generate go-bindata -pkg generator -o ../generator/schema.go.tmpl.go ../generator/schema.go.tmpl

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
		Usage: "meta file",
	}

	app.Commands = []cli.Command{
		{
			Name:  "doc",
			Usage: "generate docuemnt from schema.json",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output, o",
					Usage: "the path of the output file",
				},
				cli.StringFlag{
					Name:  "template, t",
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
				meta,
				watch,
			},
			Action: func(c *cli.Context) {
				if err := doc.Generate(doc.Options{
					Input:    c.Args()[0],
					Template: c.String("template"),
					Engine:   doc.Engine(c.String("engine")),
					Output:   c.String("output"),
					Format:   c.String("format"),
					Meta:     c.String("meta"),
					IsWatch:  c.Bool("watch"),
				}); err != nil {
					log.Println(err)
				}
			},
		},

		{
			Name:  "combine",
			Usage: "combine partial schema files into one schema file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output, o",
					Usage: "the path of the output file",
				},
				cli.StringFlag{
					Name:  "encoding, e",
					Value: "json",
					Usage: "the encoding of output file",
				},
				meta,
				watch,
			},
			Action: func(c *cli.Context) {
				if err := combine.Run(combine.Options{
					Input:    c.Args()[0],
					Output:   c.String("output"),
					Encoding: encoding.Encoding(c.String("encoding")),
					Meta:     c.String("meta"),
					IsWatch:  c.Bool("watch"),
				}); err != nil {
					log.Println(err)
				}
			},
		},

		{
			Name:  "generate",
			Usage: "generate go files of router and struct",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output, o",
					Usage: "the path of the output file",
				},
				cli.StringFlag{
					Name:  "template, t",
					Usage: "template filename",
				},
				watch,
			},
			Action: func(c *cli.Context) {
				args := c.Args()
				var input string
				if len(args) > 0 {
					input = args[0]
				} else {
					input = ""
				}
				if err := generator.Run(generator.Options{
					Input:    input,
					Output:   c.String("output"),
					Template: c.String("template"),
				}); err != nil {
					log.Println(err)
				}
			},
		},
	}

	app.Run(os.Args)
}
