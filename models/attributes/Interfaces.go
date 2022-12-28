package attributes

type AttributesBaseID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type AttributesBaseName interface {
	SetName(name string)
	GetName() STR_PTR
}

type AttributesBase interface {
	AttributesBaseID
	AttributesBaseName
}
