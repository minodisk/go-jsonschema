package jsonschema

import (
	"log"
	"path"
)

type Definitions struct {
	SchemaMap
}

func (d *Definitions) Collect(schemas map[string]*Schema, p string) error {
	p = path.Join(p, "definitions")
	for name, schema := range d.Schemas {
		if schema == nil {
			log.Printf("schema doesn't exist at '%s' in definitions", name)
			continue
		}
		id := path.Join(p, name)
		schemas[id] = schema
		if err := schema.Collect(schemas, id); err != nil {
			return err
		}
	}
	return nil
}
