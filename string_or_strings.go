package gojsa

import (
	"encoding/json"
	"fmt"
)

type StringOrStrings struct {
	IsString bool
	String   string
	Strings  []string
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
		s.IsString = true
		s.String = t
	case []interface{}:
		var strs []string
		for _, v := range t {
			str, ok := v.(string)
			if !ok {
				fmt.Errorf("unexpected type %T", t)
			}
			strs = append(strs, str)
		}
		s.Strings = strs
	}
	return nil
}

func (s *StringOrStrings) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var str string
	var strs []string
	if err = unmarshal(&str); err == nil {
		s.IsString = true
		s.String = str
		return nil
	}
	if err = unmarshal(&strs); err == nil {
		s.Strings = strs
		return nil
	}
	return fmt.Errorf("unexpected type")
}
