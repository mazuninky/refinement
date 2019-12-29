package refiment

import (
	"fmt"
	"reflect"
)

// Error section

type ComparerType uint

const (
	EqualComparer ComparerType = iota
	LessOrEqualComparer
	LessComparer
	BiggerComparer
	BiggerOrEqualComparer
)

type NumberCompareError struct {
	Container     Container
	Comparer      ComparerType
	CompareResult CompareResult
	CompareWith   Container
}

func operationFromComparerType(comparer ComparerType) string {
	switch comparer {
	case EqualComparer:
		return "=="
	case LessOrEqualComparer:
		return "<="
	case LessComparer:
		return "<"
	case BiggerComparer:
		return "<"
	case BiggerOrEqualComparer:
		return ">"
	}

	panic("unknown comparer type")
}

func (e *NumberCompareError) Error() string {
	return fmt.Sprintf("Number comparing error. Expected %f %v %f", e.Container.Value(),
		operationFromComparerType(e.Comparer), e.CompareWith.Value())
}

func NewNumberCompareError(container Container, comparer ComparerType, compareResult CompareResult, compareWith Container) *NumberCompareError {
	return &NumberCompareError{
		Container:     container,
		Comparer:      comparer,
		CompareResult: compareResult,
		CompareWith:   compareWith,
	}
}

// Map section

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

// Type section

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

// Type section

type CompareResult int

const (
	LessValue   CompareResult = -1
	EqualValue  CompareResult = 0
	BiggerValue CompareResult = 1
)

type Container interface {
	Value() float64
	CompareInt(value int64) CompareResult
	CompareFloat(value float64) CompareResult
}

var expectedType = []reflect.Kind{
	reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
	reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
	reflect.Float32, reflect.Float64,
}

func NewContainer(value interface{}) (container Container, err error) {
	switch typedValue := value.(type) {
	case int:
		container = NewIntContainer(typedValue)
		break
	case int8:
		container = NewInt8Container(typedValue)
		break
	case int16:
		container = NewInt16Container(typedValue)
		break
	case int32:
		container = NewInt32Container(typedValue)
		break
	case int64:
		container = NewInt64Container(typedValue)
		break
	case uint:
		container = NewUIntContainer(typedValue)
		break
	case uint8:
		container = NewUInt8Container(typedValue)
		break
	case uint16:
		container = NewUInt16Container(typedValue)
		break
	case uint32:
		container = NewUInt32Container(typedValue)
		break
	case uint64:
		container = NewUInt64Container(typedValue)
		break
	case float32:
		container = NewFloat32Container(typedValue)
		break
	case float64:
		container = NewFloatContainer(typedValue)
		break
	default:
		container = nil
	}

	if container == nil {
		err = NewMismatchTypeErrorFromInterface(expectedType, value)
	} else {
		err = nil
	}

	return
}

type IntContainer struct {
	value int64
}

func NewIntContainer(value int) Container {
	return &IntContainer{value: int64(value)}
}

func NewInt8Container(value int8) Container {
	return &IntContainer{value: int64(value)}
}

func NewInt16Container(value int16) Container {
	return &IntContainer{value: int64(value)}
}

func NewInt32Container(value int32) Container {
	return &IntContainer{value: int64(value)}
}

func NewInt64Container(value int64) Container {
	return &IntContainer{value: value}
}

func NewUIntContainer(value uint) Container {
	return &IntContainer{value: int64(value)}
}

func NewUInt8Container(value uint8) Container {
	return &IntContainer{value: int64(value)}
}

func NewUInt16Container(value uint16) Container {
	return &IntContainer{value: int64(value)}
}

func NewUInt32Container(value uint32) Container {
	return &IntContainer{value: int64(value)}
}

func NewUInt64Container(value uint64) Container {
	return &IntContainer{value: int64(value)}
}

func (container *IntContainer) CompareInt(value int64) CompareResult {
	if container.value == value {
		return EqualValue
	} else if container.value > value {
		return BiggerValue
	}

	return LessValue
}

func (container *IntContainer) CompareFloat(value float64) CompareResult {
	floatContainer := float64(container.value)
	if floatContainer == value {
		return EqualValue
	} else if floatContainer > value {
		return BiggerValue
	}

	return LessValue
}

func (container *IntContainer) Value() float64 {
	return float64(container.value)
}

type FloatContainer struct {
	value float64
}

func NewFloatContainer(value float64) Container {
	return &FloatContainer{value: value}
}

func NewFloat32Container(value float32) Container {
	return &FloatContainer{value: float64(value)}
}

func (container *FloatContainer) CompareInt(value int64) CompareResult {
	floatValue := float64(value)
	if container.value == floatValue {
		return EqualValue
	} else if container.value > floatValue {
		return BiggerValue
	}

	return LessValue
}

func (container *FloatContainer) CompareFloat(value float64) CompareResult {
	if container.value == value {
		return EqualValue
	} else if container.value > value {
		return BiggerValue
	}

	return LessValue
}

func (container *FloatContainer) Value() float64 {
	return container.value
}

func Compare(first Container, second Container) CompareResult {
	switch secondTyped := second.(type) {
	case *IntContainer:
		return first.CompareInt(secondTyped.value)
	case *FloatContainer:
		return first.CompareFloat(secondTyped.value)
	default:
		panic("unknown Container implementation")
	}
}
