package gojsa

import "path"

type Definitions struct {
	SchemaMap
}

func (d *Definitions) Collect(schemas *map[string]*Schema, p string) error {
	p = path.Join(p, "definitions")
	for name, schema := range d.Schemas {
		id := path.Join(p, name)
		(*schemas)[id] = schema
		if err := schema.Collect(schemas, id); err != nil {
			return err
		}
	}
	return nil
}
