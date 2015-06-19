package format

type Definitions struct {
	storage map[string]Definition
}

func NewDefinitions() *Definitions {
	ds := &Definitions{}
	ds.storage = map[string]Definition{}
	ds.AddDefinition(DateTime)
	ds.AddDefinition(Email)
	ds.AddDefinition(Hostname)
	ds.AddDefinition(IPV4)
	ds.AddDefinition(IPV6)
	ds.AddDefinition(URI)
	return ds
}

// Implementations MAY add custom format attributes.
// http://json-schema.org/latest/json-schema-validation.html#anchor106
func (ds *Definitions) AddDefinition(d Definition) {
	ds.storage[d.Name] = d
}

func (ds *Definitions) FindDefinition(name string) (Definition, bool) {
	d, ok := ds.storage[name]
	return d, ok
}
