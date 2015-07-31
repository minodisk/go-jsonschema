package dummy

import "fmt"

type MinimumError struct {
	Defined, Specified float64
}

func (err MinimumError) Error() string {
	return fmt.Sprintf("MinimumError: %f is defined as minimum, but specified %f", err.Defined, err.Specified)
}

func Minimum(defined float64, specified interface{}) error {
	if specified < defined {
		return MinimuError{defined, specified}
	}
	return nil
}

type MaxItemsError struct {
	Max    int
	Length int
}

func (err MaxItemsError) Error() string {
	return fmt.Sprintf("MaxItemError: %d is defined as max, but actual %d", err.Max, err.Length)
}

func MaxItems(max int, items []interface{}) error {
	length := len(items)
	if length > max {
		return MaxItemsError{max, length}
	}
	return nil
}

type MinItemError struct {
	Min    int
	Length int
}

func (err MinItemsError) Error() string {
	return fmt.Sprintf("MinItemError: %d is defined as max, but actual %d", err.Min, err.Length)
}

func MinItems(min int, items []interface{}) error {
	length := len(items)
	if length < min {
		return MinItemsError{min, length}
	}
	return nil
}

type UniqueItemError struct {
	Item interface{}
}

func (err UniqueItemError) Error() string {
	return fmt.Sprintf("UniqueItemError: %v is duplicated", err.Item)
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
