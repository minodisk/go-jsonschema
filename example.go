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

func (e Example) String() string {
	if v, ok := e.value.(string); ok {
		return string(v)
	}
	return ""
}

func (e Example) TypedString() string {
	switch t := e.value.(type) {
	default:
		return ""
	case nil:
		return "null"
	case string:
		return fmt.Sprintf("\"%s\"", t)
	case float64:
		return fmt.Sprintf("%f", t)
	case bool:
		return fmt.Sprintf("%t", t)
	}
}
