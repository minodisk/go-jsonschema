package encoding

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/yaml.v2"
)

func YAMLToJSON(y []byte) (j []byte, err error) {
	var iy interface{}
	err = yaml.Unmarshal(y, &iy)
	if err != nil {
		return nil, err
	}

	ij, err := KeyValueMap(iy)
	if err != nil {
		return nil, err
	}

	return json.Marshal(ij)
}

// func UnmarshalFile(filename string) (i interface{}, err error) {
// 	b, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	enc, err := NewWithFilename(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	switch enc {
// 	default:
// 		return nil, fmt.Errorf("unsupported encoding '%s'", enc)
// 	case JSONEncoding:
// 		if err := json.Unmarshal(b, &i); err != nil {
// 			return nil, err
// 		}
// 		return i, nil
// 	case YAMLEncoding:
// 		if err := yaml.Unmarshal(b, &i); err != nil {
// 			return nil, err
// 		}
// 		return KeyValueMap(i)
// 	}
// }

func KeyValueMap(input interface{}) (interface{}, error) {
	switch i := input.(type) {
	default:
		return i, nil
	case map[interface{}]interface{}:
		output := make(map[string]interface{})
		var oKey string
		for key, val := range i {
			switch k := key.(type) {
			default:
				return nil, fmt.Errorf("unsupported key type %T", k)
			case int:
				oKey = strconv.Itoa(k)
			case string:
				oKey = k
			}
			oVal, err := KeyValueMap(val)
			if err != nil {
				return nil, err
			}
			output[oKey] = oVal
		}
		return output, nil
	case []interface{}:
		output := make([]interface{}, len(i))
		for index, val := range i {
			oVal, err := KeyValueMap(val)
			if err != nil {
				return nil, err
			}
			output[index] = oVal
		}
		return output, nil
	}
}
