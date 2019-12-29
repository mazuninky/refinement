package _old

import "reflect"

// Argument: Pointer of struct
func NewJsonType(jsonType interface{}) (RefinementType, error) {
	if reflect.TypeOf(jsonType).Kind() == reflect.Ptr {
		return nil, NewMismatchTypeErrorFromInterface([]reflect.Kind{reflect.Ptr}, jsonType)
	}

	mapFunc := createJsonUnmarshalMapFunc(jsonType)
	return NewType(mapFunc), nil
}

// Argument: Pointer of struct
func MustNewJsonType(jsonType interface{}) RefinementType {
	if reflect.TypeOf(jsonType).Kind() == reflect.Ptr {
		panic(NewMismatchTypeErrorFromInterface([]reflect.Kind{reflect.Ptr}, jsonType))
	}

	mapFunc := createJsonUnmarshalMapFunc(jsonType)
	return NewType(mapFunc)
}

