package marker

type STR_PTR *string

type MarkerID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type MarkerName interface {
	SetName(suffix string)
	GetName() STR_PTR
}

type MarkerBase interface {
	MarkerID
	MarkerName
}
