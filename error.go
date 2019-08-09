package blood_contracts_go

import (
	"fmt"
	"reflect"
)

type MismatchTypeError struct {
	Expected []reflect.Kind
	Got      reflect.Kind
}

func (e *MismatchTypeError) Error() string {
	return fmt.Sprintf("Mismatch type. Expected: %v; Got: %v", e.Expected, e.Got)
}

func NewMismatchTypeError(expected []reflect.Kind, got reflect.Kind) *MismatchTypeError {
	return &MismatchTypeError{
		Expected: expected,
		Got:      got,
	}
}

func NewMismatchTypeErrorFromInterface(expected []reflect.Kind, got interface{}) *MismatchTypeError {
	return NewMismatchTypeError(expected, reflect.TypeOf(got).Kind())
}

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
