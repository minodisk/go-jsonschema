package jsonschema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type Example struct {
	value interface{}
}

func NewDefaultExample(t *Type) (Example, error) {
	if t.Is(TypeNull) {
		return Example{nil}, nil
	} else if t.Is(TypeBoolean) {
		return Example{false}, nil
	} else if t.Is(TypeNumber) || t.Is(TypeInteger) {
		return Example{0.0}, nil
	} else if t.Is(TypeString) {
		return Example{""}, nil
	}
	return Example{}, fmt.Errorf("no default example: %s", t.String())
}

func (e *Example) UnmarshalJSON(data []byte) error {
	d := json.NewDecoder(bytes.NewBuffer(data))
	d.UseNumber()
	return d.Decode(&e.value)
}

func (e Example) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.value)
}

func (e Example) HasValue() bool {
	return e.value != nil
}

func (e Example) IsDefined() bool {
	return e.value != nil
}

func (e Example) String() string {
	if r, ok := e.value.(string); ok {
		return fmt.Sprintf("\"%s\"", r)
	}
	return e.RawString()
}

func (e Example) RawString() string {
	switch v := e.value.(type) {
	default:
		return ""
	case []interface{}, map[string]interface{}:
		b, err := json.Marshal(v)
		if err != nil {
			return ""
		}
		return string(b)
	case string:
		return v
	case json.Number:
		return v.String()
	case bool:
		return strconv.FormatBool(v)
	case nil:
		return "null"
	}
}
