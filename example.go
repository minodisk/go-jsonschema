package gojsa

import (
	"encoding/json"
	"fmt"
)

type Example struct {
	value interface{}
}

func (e *Example) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &e.value); err != nil {
		return err
	}
	return nil
}

func (e Example) Value() interface{} {
	return e.value
}

func (e Example) String() string {
	switch v := e.value.(type) {
	default:
		return ""
	case nil:
		return "null"
	case string:
		return v
	case float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return fmt.Sprintf("%t", v)
	}
}

func (e Example) TypedString() string {
	switch v := e.value.(type) {
	default:
		return ""
	case nil:
		return "null"
	case string:
		return fmt.Sprintf("\"%s\"", v)
	case float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return fmt.Sprintf("%t", v)
	}
}
