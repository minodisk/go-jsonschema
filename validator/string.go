package validator

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

func (v Validator) MaxLength(s string, max int) error {
	l := utf8.RuneCountInString(s)
	if l > max {
		return MaxLengthError{v, l, max}
	}
	return nil
}

func (v Validator) MinLength(s string, min int) error {
	l := utf8.RuneCountInString(s)
	if l < min {
		return MinLengthError{v, l, min}
	}
	return nil
}

func (v Validator) Pattern(s string, pattern string) error {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	if !r.MatchString(s) {
		return PatternError{v, s, pattern}
	}
	return nil
}

type MaxLengthError struct {
	Validator
	Length int
	Max    int
}

func (err MaxLengthError) Error() string {
	return fmt.Sprintf("the length of the string %s in %s should be less than or equal to %d, but has %d characters", err.PropertyName, err.SchemaName, err.Max, err.Length)
}

type MinLengthError struct {
	Validator
	Length int
	Min    int
}

func (err MinLengthError) Error() string {
	return fmt.Sprintf("the length of the string %s in %s should be greater than or equal to %d, but has %d characters", err.PropertyName, err.SchemaName, err.Min, err.Length)
}

type PatternError struct {
	Validator
	String  string
	Pattern string
}

func (err PatternError) Error() string {
	return fmt.Sprintf("the pattern of the string %s in %s should be matched to %s, but %s is specified", err.PropertyName, err.SchemaName, err.Pattern, err.String)
}
