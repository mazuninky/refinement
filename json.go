package refiment

import (
	"encoding/json"
	"reflect"
)

// Map section

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

// Type section

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

