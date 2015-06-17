package jsonschema

type Title string

func (t Title) DocString() string {
	return string(t)
}
