package main

import (
	"log"
	"os"

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
				watch,
				meta,
			},
			Action: func(c *cli.Context) {
				if err := doc.Generate(doc.Options{
					Input:    c.Args()[0],
					Template: c.String("template"),
					Engine:   doc.Engine(c.String("engine")),
					Output:   c.String("output"),
					Format:   c.String("format"),
					IsWatch:  c.Bool("watch"),
					Meta:     c.String("meta"),
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
					Input:    c.Args()[0],
					Meta:     c.String("meta"),
					Output:   c.String("output"),
					Encoding: encoding.Encoding(c.String("encoding")),
					IsWatch:  c.Bool("watch"),
				}); err != nil {
					log.Println(err)
				}
			},
		},
	}

	app.Run(os.Args)
}
