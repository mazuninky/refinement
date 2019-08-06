package refinement

import (
	"errors"
	. "github.com/mazuninky/blood-contracts-go/core"
)

type RType struct {
	mapFunc MapFunction
	//Pack(value interface{}) RefinementTypeBox
	//And(rType RefinementType) RefinementType
	//Pipe(rType RefinementType) RefinementType
	//Or(rType RefinementType) RefinementType
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

func (base *RType) GetMapFunction() MapFunction {
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

		secondValue, secondError := rt.GetMapFunction()(value)
		if secondError != nil {
			return nil, secondError
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

		secondValue, secondError := rt.GetMapFunction()(value)
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

		return rt.GetMapFunction()(firstValue)
	}

	return NewType(mapFunc)
}
