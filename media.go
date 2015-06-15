package jsonschema

type Media struct {
	Type           string
	BinaryEncoding string
}

// func Parse(r io.Reader) (s Schema, err error) {
// 	dec := json.NewDecoder(r)
// 	if err := dec.Decode(&s); err != nil {
// 		return s, err
// 	}
// 	return s, nil
// }
