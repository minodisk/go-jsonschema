package validator

import (
	"fmt"
	"math"
)

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

type MultipleOfError struct {
	Validator
	Specified, MultipleOf float64
}

func (err MultipleOfError) Error() string {
	return fmt.Sprintf("the value of %s in %s should be multiple of %f, but specified %f", err.SchemaName, err.PropertyName, err.MultipleOf, err.Specified)
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
