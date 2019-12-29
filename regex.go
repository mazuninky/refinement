package refiment

import (
	"errors"
	"reflect"
	"regexp"
)

// Error section

var RegexMismatchError = errors.New("regex mismatch")

// Type section

func NewRegexType(regex string) (RefinementType, error) {
	reg, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}
	regexMapFunc := createRegexMapFunc(reg)
	return NewType(regexMapFunc), nil
}

func MustNewRegexType(regex string) RefinementType {
	reg, err := regexp.Compile(regex)
	if err != nil {
		panic(err)
	}
	regexMapFunc := createRegexMapFunc(reg)
	return NewType(regexMapFunc)
}

// Map section

var regexExpectedTypes = []reflect.Kind{
	reflect.String,
}

func createRegexMapFunc(regexp *regexp.Regexp) MapFunction {
	return func(value interface{}) (interface{}, error) {
		toMatch, isString := value.(string)
		if !isString {
			return nil, NewMismatchTypeError(regexExpectedTypes, reflect.TypeOf(value).Kind())
		}
		if regexp.MatchString(toMatch) {
			return value, nil
		} else {
			return nil, RegexMismatchError
		}
	}
}
