package blood_contracts_go

import "errors"

var positiveMapFunc = func(value interface{}) (interface{}, error){
	return value, nil
}

var negativeMapErr = errors.New("negative")

var negativeMapFunc = func(value interface{}) (interface{}, error){
	return nil, negativeMapErr
}
