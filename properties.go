package jsonschema

import "net/url"

type Properties struct {
	SchemaMap
}

func (p *Properties) QueryString() string {
	v := url.Values{}
	for name, schema := range p.Schemas {
		v.Set(name, schema.Example.String())
	}
	return v.Encode()
}

func (p *Properties) ExampleData() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	for name, schema := range p.Schemas {
		if schema != nil {
			if schema.Example != nil {
				m[name] = schema.Example.Value()
			} else {
				v, err := schema.ExampleData()
				if err != nil {
					return nil, err
				}
				m[name] = v
			}
		}
	}
	return m, nil
}
