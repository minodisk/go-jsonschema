package gojsa

type SchemaMap struct {
	Schemas map[string]*Schema
}

func (s *SchemaMap) Collect(schemas *map[string]*Schema) (err error) {
	for _, schema := range s.Schemas {
		if err := schema.Collect(schemas); err != nil {
			return err
		}
	}
	return nil
}
