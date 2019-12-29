package refiment

import "errors"

var validMapFunc = func(value interface{}) (interface{}, error) {
	return value, nil
}

var mapErr = errors.New("negative")

var errorMapFunc = func(value interface{}) (interface{}, error) {
	return nil, mapErr
}
