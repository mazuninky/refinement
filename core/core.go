package core

type RefinementType interface {
	Pack(value interface{}) RefinementTypeBox
	IsValid(value interface{}) bool
	And(rType RefinementType) RefinementType
	Pipe(rType RefinementType) RefinementType
	Or(rType RefinementType) RefinementType

	GetMapFunction() MapFunction
}

type RefinementTypeBox interface {
	Unpack() (interface{}, error)
}