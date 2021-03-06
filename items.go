package jsonschema

type Items struct {
	SchemaOrSchemas
}

func (i Items) ExampleData(includesReadOnly bool) ([]interface{}, error) {
	var arr []interface{}
	if i.IsSchema {
		d, err := i.Schema.ExampleData(includesReadOnly)
		if err != nil {
			return nil, err
		}
		arr = append(arr, d)
		return arr, nil
	}
	for _, schema := range i.Schemas {
		d, err := schema.ExampleData(includesReadOnly)
		if err != nil {
			return nil, err
		}
		arr = append(arr, d)
	}
	return arr, nil
}

func (i Items) GoType() string {
	if i.IsSchema {
		if i.Schema.Title != "" {
			return i.Schema.Title
		}
		return i.Schema.Type.Go
	}
	return "interface{}"
}
