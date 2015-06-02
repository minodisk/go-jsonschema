package parser

import (
	"encoding/json"
	"io"

	"github.com/minodisk/go-json-schema/schema"
)

func Parse(r io.Reader) (root schema.Root) {
	dec := json.NewDecoder(r)
	if err := dec.Decode(&root); err != nil {
		panic(err)
	}

	return root
}
