package pool

type STR_PTR *string

type PoolID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type PoolName interface {
	SetName(name string)
	GetName() STR_PTR
}

type PoolBase interface {
	PoolID
	PoolName
}
