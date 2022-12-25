package canvas

type STR_PTR *string

type CanvasID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type CanvasElement interface {
	SetElement(typ string, suffix interface{})
	GetElement() STR_PTR
}

type CanvasBase interface {
	CanvasID
	CanvasElement
}
