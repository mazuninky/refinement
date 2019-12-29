package refiment

type MapFunction func(interface{}) (interface{}, error)

type RefinementType interface {
	// Pack value into the type container
	Pack(value interface{}) RefinementTypeBox
	Pipe(rType RefinementType) RefinementType
	Or(rType RefinementType) RefinementType
	Map(mapFunc MapFunction) RefinementType

	getMapFunction() MapFunction
}

type RefinementTypeBox interface {
	Unpack() (interface{}, error)
	Map(mapFunc MapFunction) RefinementTypeBox
	IsValid() bool
}