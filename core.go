package blood_contracts_go

type RefinementType interface {
	Pack(value interface{}) RefinementTypeBox
	IsValid(value interface{}) bool
	And(rType RefinementType) RefinementType
	Pipe(rType RefinementType) RefinementType
	Or(rType RefinementType) RefinementType
	Map(mapFunc MapFunction) RefinementType

	getMapFunction() MapFunction
}

type RefinementTypeBox interface {
	Unpack() (interface{}, error)
}