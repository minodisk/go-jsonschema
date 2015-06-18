package jsonschema

import (
	"encoding/json"
	"fmt"
)

type StringOrStrings struct {
	isString bool
	string   string
	strings  []string
}

func (s *StringOrStrings) UnmarshalJSON(data []byte) (err error) {
	var i interface{}
	if err = json.Unmarshal(data, &i); err != nil {
		return err
	}
	switch t := i.(type) {
	default:
		return fmt.Errorf("unexpected type %T", t)
	case string:
		s.isString = true
		s.string = t
	case []interface{}:
		var strs []string
		for _, v := range t {
			str, ok := v.(string)
			if !ok {
				fmt.Errorf("unexpected type %T", t)
			}
			strs = append(strs, str)
		}
		s.strings = strs
	}
	return nil
}

// func (s *StringOrStrings) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
// 	var str string
// 	var strs []string
// 	if err = unmarshal(&str); err == nil {
// 		s.isString = true
// 		s.string = str
// 		return nil
// 	}
// 	if err = unmarshal(&strs); err == nil {
// 		s.strings = strs
// 		return nil
// 	}
// 	return fmt.Errorf("unexpected type")
// }
//
// func (t Type) String() string {
// 	if t.isString {
// 		return t.string
// 	}
// 	return strings.Join(t.strings, ", ")
// }
