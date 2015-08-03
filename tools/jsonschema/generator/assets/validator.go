package validator

import (
	"fmt"
	"math"
)

type Validator struct {
	SchemaName, PropertyName string
}

func (v Validator) MultipleOf(t, s float64) error {
	d := t / s
	if math.Ceil(d)-d != 0 {
		return MultipleOfError{v, t, s}
	}
	return nil
}

func (v Validator) Maximum(t, s float64, e bool) error {
	if e {
		if t >= s {
			return MaximumError{v, t, s, e}
		}
	} else {
		if t > s {
			return MaximumError{v, t, s, e}
		}
	}
	return nil
}

func (v Validator) Minimum(t, s float64, e bool) error {
	if e {
		if t <= s {
			return MinimumError{v, t, s, e}
		}
	} else {
		if t < s {
			return MinimumError{v, t, s, e}
		}
	}
	return nil
}

func MaxItems(max int, items []interface{}) error {
	length := len(items)
	if length > max {
		return MaxItemsError{max, length}
	}
	return nil
}

func MinItems(min int, items []interface{}) error {
	length := len(items)
	if length < min {
		return MinItemsError{min, length}
	}
	return nil
}

func UniqueItems(items []interface{}) error {
	for i, item := range items {
		rests := items[i+1:]
		for _, rest := range rests {
			if item == rest {
				return UniqueItemError{item}
			}
		}
	}
	return nil
}

type MultipleOfError struct {
	Validator
	Specified, MultipleOf float64
}

func (err MultipleOfError) Error() string {
	return fmt.Sprintf("the value of %s in %s should be multiple of %f, but specified with %f", err.SchemaName, err.PropertyName, err.MultipleOf, err.Specified)
}

type MaximumError struct {
	Validator
	Specified, Maximum float64
	Exclusive          bool
}

func (err MaximumError) Error() string {
	return fmt.Sprintf("%f is defined as maximum, but specified %f", err.Maximum, err.Specified)
}

type MinimumError struct {
	Validator
	Defined, Specified float64
	Exclusive          bool
}

func (err MinimumError) Error() string {
	return fmt.Sprintf("MinimumError: %f is defined as minimum, but specified %f", err.Defined, err.Specified)
}

type MaxItemsError struct {
	Max    int
	Length int
}

func (err MaxItemsError) Error() string {
	return fmt.Sprintf("MaxItemError: %d is defined as max, but actual %d", err.Max, err.Length)
}

type MinItemsError struct {
	Min    int
	Length int
}

func (err MinItemsError) Error() string {
	return fmt.Sprintf("MinItemError: %d is defined as max, but actual %d", err.Min, err.Length)
}

type UniqueItemError struct {
	Item interface{}
}

func (err UniqueItemError) Error() string {
	return fmt.Sprintf("UniqueItemError: %v is duplicated", err.Item)
}
