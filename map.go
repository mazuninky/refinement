package blood_contracts_go

import (
	"encoding/json"
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

func createMinMapFunc(container Container, include bool) MapFunction {
	return func(value interface{}) (interface{}, error) {
		toCompare, err := NewContainer(value)
		if err != nil {
			return nil, err
		}

		comparator := Compare(container, toCompare)
		if include && comparator == EqualValue || comparator == LessValue {
			return value, nil
		}

		var comparatorType ComparerType
		if include {
			comparatorType = BiggerOrEqualComparer
		} else {
			comparatorType = BiggerComparer
		}

		return nil, NewNumberCompareError(container, comparatorType, comparator, toCompare)
	}
}

func createMaxMapFunc(container Container, include bool) MapFunction {
	return func(value interface{}) (interface{}, error) {
		toCompare, err := NewContainer(value)
		if err != nil {
			return nil, err
		}

		comparator := Compare(container, toCompare)
		if include && comparator == EqualValue || comparator == BiggerValue {
			return value, nil
		}

		var comparatorType ComparerType
		if include {
			comparatorType = LessOrEqualComparer
		} else {
			comparatorType = LessComparer
		}

		return nil, NewNumberCompareError(container, comparatorType, comparator, toCompare)
	}
}

func createEqualMapFunc(container Container) MapFunction {
	return func(value interface{}) (interface{}, error) {
		toCompare, err := NewContainer(value)
		if err != nil {
			return nil, err
		}

		comparator := Compare(container, toCompare)
		if comparator == EqualValue {
			return value, nil
		}

		return nil, NewNumberCompareError(container, EqualComparer, comparator, toCompare)
	}
}

var jsonExpectedTypes = []reflect.Kind{
	reflect.String, reflect.Array,
}

func createJsonUnmarshalMapFunc(jsonType interface{}) MapFunction {
	return func(value interface{}) (interface{}, error) {
		var data []byte
		switch typedValue := value.(type) {
		case string:
			data = []byte(typedValue)
			break
		case []byte:
			data = typedValue
			break
		default:
			return nil, NewMismatchTypeErrorFromInterface(jsonExpectedTypes, data)
		}

		clone := reflect.New(reflect.ValueOf(jsonType).Elem().Type()).Interface()
		err := json.Unmarshal(data, clone)

		// TODO Map to own error
		if err != nil {
			return nil, err
		}

		return clone, nil
	}
}
