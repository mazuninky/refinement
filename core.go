package blood_contracts_go

type RefinementType interface {
	Pack(value interface{}) RefinementTypeBox
	IsValid(value interface{}) bool
	And(rType RefinementType) RefinementType
	Pipe(rType RefinementType) RefinementType
	Or(rType RefinementType) RefinementType

	getMapFunction() mapFunction
}

type RefinementTypeBox interface {
	Unpack() (interface{}, error)
}