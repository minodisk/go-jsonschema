package jsonschema

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Example struct {
	raw interface{}
}

func (e *Example) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &e.raw); err != nil {
		return err
	}
	return nil
}

func (e *Example) MarshalJSON() ([]byte, error) {
	return []byte(e.String()), nil
}

func (e Example) String() string {
	if r, ok := e.raw.(string); ok {
		return fmt.Sprintf("\"%s\"", r)
	}
	return e.RawString()
}

func (e Example) RawString() string {
	switch v := e.raw.(type) {
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
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case nil:
		return "null"
	}
}
