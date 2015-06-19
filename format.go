package jsonschema

import (
	"encoding/json"
	"fmt"

	"github.com/minodisk/go-jsonschema/format"
)

type Format struct {
	value       string
	definitions *format.Definitions
}

func (f *Format) UnmarshalJSON(raw []byte) error {
	var value string
	if err := json.Unmarshal(raw, &value); err != nil {
		return err
	}
	f.value = value
	return nil
}

func (f *Format) Resolve(schemas *map[string]*Schema, root *Schema) error {
	f.definitions = root.definitions
	return nil
}

func (f *Format) String() string {
	return f.value
}

func (f *Format) ExampleData() (*Example, error) {
	d, ok := f.definitions.FindDefinition(f.value)
	if !ok {
		return nil, fmt.Errorf("no format definition: %s", f.value)
	}
	return &Example{d.Example}, nil
}
