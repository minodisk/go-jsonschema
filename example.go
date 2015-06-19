package jsonschema

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Example struct {
	value interface{}
}

func NewDefaultExample(t *Type) (*Example, error) {
	if t.Is(TypeNull) {
		return &Example{nil}, nil
	} else if t.Is(TypeBoolean) {
		return &Example{false}, nil
	} else if t.Is(TypeNumber) || t.Is(TypeInteger) {
		return &Example{0.0}, nil
	} else if t.Is(TypeString) {
		return &Example{""}, nil
	}
	return nil, fmt.Errorf("no default example: %s", t.String())
}

func (e *Example) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &e.value); err != nil {
		return err
	}
	return nil
}

func (e *Example) MarshalJSON() ([]byte, error) {
	return []byte(e.String()), nil
}

func (e *Example) Initialize(def interface{}) {
	if !e.IsDefined() {
		e.value = def
	}
}

func (e *Example) IsDefined() bool {
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
	// case int:
	// 	i := int64(v)
	// 	log.Println("->", i)
	// 	return strconv.FormatInt(i, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case nil:
		return "null"
	}
}
