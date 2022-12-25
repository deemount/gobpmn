package core

type STR_PTR *string

type CoreID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type CoreBase interface {
	CoreID
}
