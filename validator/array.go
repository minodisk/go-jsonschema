package validator

import (
	"fmt"
	"reflect"
)

func (v Validator) MaxItems(items interface{}, max int) error {
	if reflect.TypeOf(items).Kind() != reflect.Slice {
		return fmt.Errorf("TypeError")
	}

	array := reflect.ValueOf(items)
	l := array.Len()
	if l > max {
		return MaxItemsError{v, l, max}
	}
	return nil
}

func (v Validator) MinItems(items interface{}, min int) error {
	if reflect.TypeOf(items).Kind() != reflect.Slice {
		return fmt.Errorf("TypeError")
	}

	array := reflect.ValueOf(items)
	l := array.Len()
	if l < min {
		return MinItemsError{v, l, min}
	}
	return nil
}

func (v Validator) UniqueItems(items interface{}) error {
	if reflect.TypeOf(items).Kind() != reflect.Slice {
		return fmt.Errorf("TypeError")
	}

	s := reflect.ValueOf(items)
	l := s.Len()
	for i := 0; i < l; i++ {
		a := s.Index(i).Interface()
		for j := i + 1; j < l; j++ {
			b := s.Index(j).Interface()
			if a == b {
				return UniqueItemError{v, a}
			}
		}
	}
	return nil
}

type MaxItemsError struct {
	Validator
	Length int
	Max    int
}

func (err MaxItemsError) Error() string {
	return fmt.Sprintf("the length of %s in %s should be less than or equal to %d, but has %d items", err.PropertyName, err.SchemaName, err.Max, err.Length)
}

type MinItemsError struct {
	Validator
	Length int
	Min    int
}

func (err MinItemsError) Error() string {
	return fmt.Sprintf("the length of %s in %s should be greater than or equal to %d, but has %d items", err.PropertyName, err.SchemaName, err.Min, err.Length)
}

type UniqueItemError struct {
	Validator
	Item interface{}
}

func (err UniqueItemError) Error() string {
	return fmt.Sprintf("the items of %s in %s should contains unique items, but %v is duplicated", err.PropertyName, err.SchemaName, err.Item)
}
