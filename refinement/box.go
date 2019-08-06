package refinement

import(. "github.com/mazuninky/blood-contracts-go/core")

type RTBox struct {
	mapFunc  MapFunction
	isMapped bool
	value    interface{}
	result   interface{}
	err      error
}

func NewBox(mapFunc MapFunction, value interface{}) RefinementTypeBox {
	box := RTBox{
		mapFunc:  mapFunc,
		isMapped: false,
		value:    value,
		result:   nil,
		err:      nil,
	}

	return &box
}

func (rTBox *RTBox) Unpack() (interface{}, error) {
	if !rTBox.isMapped {
		rTBox.result, rTBox.err = rTBox.mapFunc(rTBox.value)
	}

	return rTBox.result, rTBox.err
}
