package gojsa

type Items struct {
	SchemaOrSchemas
}

func (i Items) ExampleData() ([]interface{}, error) {
	var arr []interface{}
	if i.IsSchema {
		d, err := i.Schema.ExampleData()
		if err != nil {
			return nil, err
		}
		arr = append(arr, d)
		return arr, nil
	}
	for _, schema := range i.Schemas {
		d, err := schema.ExampleData()
		if err != nil {
			return nil, err
		}
		arr = append(arr, d)
	}
	return arr, nil
}
