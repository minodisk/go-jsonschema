package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/minodisk/go-jsonschema"
)

func (v Validator) MaxProperties(o interface{}, maxProperties int) error {
	if reflect.TypeOf(o).Kind() != reflect.Map {
		return fmt.Errorf("TypeError")
	}

	object := reflect.ValueOf(o)
	keys := object.MapKeys()
	l := len(keys)
	if l > maxProperties {
		return MaxPropertiesError{v, l, maxProperties}
	}
	return nil
}

func (v Validator) MinProperties(o interface{}, minProperties int) error {
	if reflect.TypeOf(o).Kind() != reflect.Map {
		return fmt.Errorf("TypeError")
	}

	object := reflect.ValueOf(o)
	keys := object.MapKeys()
	l := len(keys)
	if l < minProperties {
		return MinPropertiesError{v, l, minProperties}
	}
	return nil
}

func (v Validator) Required(o interface{}, required []string) error {
	if reflect.TypeOf(o).Kind() != reflect.Map {
		return fmt.Errorf("TypeError")
	}

	var lacks []string
	ks := keys(o)
outer:
	for _, r := range required {
		for _, k := range ks {
			if r == k {
				continue outer
			}
		}
		lacks = append(lacks, r)
	}
	if len(lacks) > 0 {
		return RequiredError{v, lacks}
	}
	return nil
}

// Can't support dynamic properties in Go.
func (v Validator) AdditionalProperties(o interface{}, a jsonschema.SchemaOrBool) error {
	return nil
}

// Can't support dynamic properties in Go.
func (v Validator) PatternProperties(o interface{}, p map[string]string) error {
	return nil
}

func (v Validator) DependenciesStrings(o interface{}, deps map[string][]string) error {
	if reflect.TypeOf(o).Kind() != reflect.Map {
		return fmt.Errorf("TypeError")
	}

	lacks := make(map[string][]string)
	ks := keys(o)
	for k, ds := range deps {
		if !in(ks, k) {
			continue
		}
		for _, d := range ds {
			if !in(ks, d) {
				lacks[k] = append(lacks[k], d)
			}
		}
	}
	if len(lacks) > 0 {
		return DependenciesError{v, lacks}
	}
	return nil
}

func keys(o interface{}) []string {
	object := reflect.ValueOf(o)
	var keys []string
	for _, key := range object.MapKeys() {
		keys = append(keys, key.String())
	}
	return keys
}

func in(keys []string, key string) bool {
	for _, k := range keys {
		if key == k {
			return true
		}
	}
	return false
}

type MaxPropertiesError struct {
	Validator
	Length int
	Max    int
}

func (err MaxPropertiesError) Error() string {
	return fmt.Sprintf("the length of the properties %s in %s should be less than or equal to %d, but has %d properties", err.PropertyName, err.SchemaName, err.Max, err.Length)
}

type MinPropertiesError struct {
	Validator
	Length int
	Min    int
}

func (err MinPropertiesError) Error() string {
	return fmt.Sprintf("the length of the properties %s in %s should be greater than or equal to %d, but has %d properties", err.PropertyName, err.SchemaName, err.Min, err.Length)
}

type RequiredError struct {
	Validator
	Lacks []string
}

func (err RequiredError) Error() string {
	return fmt.Sprintf("the properties %s in %s requires %s", err.PropertyName, err.SchemaName, strings.Join(err.Lacks, ", "))
}

type DependenciesError struct {
	Validator
	Lacks map[string][]string
}

func (err DependenciesError) Error() string {
	var deps []string
	for k, ds := range err.Lacks {
		deps = append(deps, fmt.Sprintf("%s depends on %s", k, strings.Join(ds, ", ")))
	}
	return fmt.Sprintf("the properties %s in %s is invalid dependencies: %s", err.PropertyName, err.SchemaName, strings.Join(deps, ", "))
}
