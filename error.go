package refiment

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