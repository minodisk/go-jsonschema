package validator

import "fmt"

type Validator struct {
	SchemaName   string `json:"schema_name"`
	PropertyName string `json:"property_name"`
}

func (v Validator) String() string {
	return fmt.Sprintf("%s.%s", v.SchemaName, v.PropertyName)
}
