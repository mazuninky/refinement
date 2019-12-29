package _old

import (
	"errors"
	"regexp"
)

type RType struct {
	mapFunc MapFunction
}

func NewType(mapFunc MapFunction) RefinementType {
	rType := RType{
		mapFunc: mapFunc,
	}

	return &rType
}

func (base *RType) IsValid(value interface{}) bool {
	_, err := base.mapFunc(value)
	if err != nil {
		return false
	}
	return true
}

func (base *RType) getMapFunction() MapFunction {
	return base.mapFunc
}

func (base *RType) Pack(value interface{}) RefinementTypeBox {
	return NewBox(base.mapFunc, value)
}

func (base *RType) And(rt RefinementType) RefinementType {
	mapFunc := func(value interface{}) (interface{}, error) {
		firstValue, firstError := base.mapFunc(value)
		if firstError != nil {
			return nil, firstError
		}

		secondValue, secondError := rt.getMapFunction()(value)
		if secondError != nil {
			return nil, secondError
		}

		if firstValue == secondValue {
			return firstValue, nil
		}

		return []interface{}{
			firstValue,
			secondValue,
		}, nil
	}

	return NewType(mapFunc)
}

func (base *RType) Or(rt RefinementType) RefinementType {
	mapFunc := func(value interface{}) (interface{}, error) {
		firstValue, firstError := base.mapFunc(value)
		if firstError == nil {
			return firstValue, nil
		}

		secondValue, secondError := rt.getMapFunction()(value)
		if secondError == nil {
			return secondValue, nil
		}

		// TODO Choose exception
		return nil, errors.New("can't find type for both")
	}

	return NewType(mapFunc)
}

func (base *RType) Pipe(rt RefinementType) RefinementType {
	mapFunc := func(value interface{}) (interface{}, error) {
		firstValue, firstError := base.mapFunc(value)
		if firstError != nil {
			return nil, firstError
		}

		return rt.getMapFunction()(firstValue)
	}

	return NewType(mapFunc)
}

func (base *RType) Map(mapFunc MapFunction) RefinementType {
	baseMapFunc := base.mapFunc
	nextMapFunc := func(value interface{}) (interface{}, error) {
		baseResult, err := baseMapFunc(value)
		if err != nil {
			return nil, err
		}

		return mapFunc(baseResult)
	}

	return NewType(nextMapFunc)
}

// String

// Regex

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

// Number

func NewNumberMin(minValue interface{}) (RefinementType, error) {
	container, err := NewContainer(minValue)
	if err != nil {
		return nil, err
	}

	mapFunc := createMinMapFunc(container, true)
	return NewType(mapFunc), nil
}

func MustNewNumberMin(minValue interface{}) RefinementType {
	container, err := NewContainer(minValue)
	if err != nil {
		panic(err)
	}

	mapFunc := createMinMapFunc(container, true)
	return NewType(mapFunc)
}

func NewNumberMinExclude(minValue interface{}) (RefinementType, error) {
	container, err := NewContainer(minValue)
	if err != nil {
		return nil, err
	}

	mapFunc := createMinMapFunc(container, false)
	return NewType(mapFunc), nil
}

func MustNewNumberMinExclude(minValue interface{}) RefinementType {
	container, err := NewContainer(minValue)
	if err != nil {
		panic(err)
	}

	mapFunc := createMinMapFunc(container, false)
	return NewType(mapFunc)
}

func NewNumberMax(maxValue interface{}) (RefinementType, error) {
	container, err := NewContainer(maxValue)
	if err != nil {
		return nil, err
	}

	mapFunc := createMaxMapFunc(container, true)
	return NewType(mapFunc), nil
}

func MustNewNumberMax(maxValue interface{}) RefinementType {
	container, err := NewContainer(maxValue)
	if err != nil {
		panic(err)
	}

	mapFunc := createMaxMapFunc(container, true)
	return NewType(mapFunc)
}

func NewNumberMaxExclude(maxValue interface{}) (RefinementType, error) {
	container, err := NewContainer(maxValue)
	if err != nil {
		return nil, err
	}

	mapFunc := createMinMapFunc(container, false)
	return NewType(mapFunc), nil
}

func MustNewNumberMaxExclude(maxValue interface{}) RefinementType {
	container, err := NewContainer(maxValue)
	if err != nil {
		panic(err)
	}

	mapFunc := createMinMapFunc(container, false)
	return NewType(mapFunc)
}

func NewNumberEqual(equal interface{}) (RefinementType, error) {
	container, err := NewContainer(equal)
	if err != nil {
		return nil, err
	}

	mapFunc := createEqualMapFunc(container)
	return NewType(mapFunc), nil
}

func MustNewNumberEqual(equal interface{}) RefinementType {
	container, err := NewContainer(equal)
	if err != nil {
		panic(err)
	}

	mapFunc := createEqualMapFunc(container)
	return NewType(mapFunc)
}

// Struct

func NewStructType(valueToTypeMap map[string]RefinementType) {

}

func NewStructTypeFromBasis(typeBasis []RefinementType) {

}

// Function

func NewMapType(mapFunc MapFunction) RefinementType {
	return NewType(mapFunc)
}
