package gojsa

type Schemas struct {
	Schemas []*Schema
}

func (s *Schemas) Collect(schemas *map[string]*Schema) (err error) {
	for _, schema := range s.Schemas {
		if err := schema.Collect(schemas); err != nil {
			return err
		}
	}
	return nil
}
