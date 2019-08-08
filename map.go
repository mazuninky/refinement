package blood_contracts_go

import (
	"errors"
	"reflect"
	"regexp"
)

type MapFunction func(interface{}) (interface{}, error)

var regexExpectedTypes = []reflect.Kind{
	reflect.String,
}

var RegexMismatchError = errors.New("regex mismatch")

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
