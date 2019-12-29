package refiment

type RTBox struct {
	mapFunc  MapFunction
	isMapped bool
	value    interface{}
	// Calculated result
	result interface{}
	err    error
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

func (rTBox *RTBox) IsValid() bool {
	_, err := rTBox.Unpack()

	return err != nil
}

func (rTBox *RTBox) Map(mapFunc MapFunction) RefinementTypeBox {
	baseMapFunc := rTBox.mapFunc
	nextMapFunc := func(value interface{}) (interface{}, error) {
		baseResult, err := baseMapFunc(value)
		if err != nil {
			return nil, err
		}

		return mapFunc(baseResult)
	}

	return NewBox(nextMapFunc, rTBox.value)
}
