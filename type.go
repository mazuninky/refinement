package refiment

import (
	"errors"
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
		return nil, errors.New("can't find a type")
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