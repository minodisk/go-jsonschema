package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/minodisk/jsonschema/cli/doc"
)

const (
	version = "0.0.0"
)

func main() {
	// c := cli.NewCLI("jsonschema", version)
	// c.Args = os.Args[1:]
	// c.Commands = map[string]cli.CommandFactory{
	// 	"doc": func() (cli.Command, error) {
	// 		return &doc.Cli{}, nil
	// 	},
	// }
	//
	// exitStatus, err := c.Run()
	// if err != nil {
	// 	log.Println(err)
	// }
	// os.Exit(exitStatus)

	app := cli.NewApp()
	app.Name = "jsonschema"
	app.Usage = "fooooooooo"
	app.Action = func(c *cli.Context) {
		println("Hellow friends!")
	}

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
			},
			Action: func(c *cli.Context) {
				if err := doc.Generate(doc.Options{
					Input:    c.Args()[0],
					Encoding: doc.Encoding(c.String("encoding")),
					Template: c.String("template"),
					Engine:   doc.Engine(c.String("engine")),
					Output:   c.String("output"),
					Format:   c.String("format"),
				}); err != nil {
					log.Println(err)
				}
			},
		},
	}

	app.Run(os.Args)
}
