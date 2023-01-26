package impl

type IFBaseID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type IFBaseListenerID interface {
	SetListenerID(listenerID string)
	GetListenerID() *string
}

type IFBaseName interface {
	SetName(name string)
	GetName() STR_PTR
}

type IFBaseElement interface {
	SetElement(typ string, suffix interface{})
	GetElement() STR_PTR
}

type IFBaseValue interface {
	SetValue(value string)
	GetValue() STR_PTR
}

type IFBaseEvent interface {
	SetEvent(event string)
	GetEvent() STR_PTR
}

type IFBaseClass interface {
	SetClass(class string)
	GetClass() STR_PTR
}

type IFBaseLabel interface {
	SetLabel(label string)
	GetLabel() *string
}

type IFBaseType interface {
	SetType(typ string)
	GetType() *string
}

type IFBaseConfig interface {
	SetConfig(config string)
	GetConfig() *string
}

type IFBaseKey interface {
	SetKey(key string)
	GetKey() *string
}

type IFBaseDefaultValue interface {
	SetDefaultValue(defaultValue string)
	GetDefaultValue() *string
}

type IFBaseScriptFormat interface {
	SetScriptFormat(format string)
	GetScriptFormat() *string
}

type IFBaseLocalVariableName interface {
	SetLocalVariableName(variable string)
	GetLocalVariableName() *string
}

type IFBaseVariableAssignmentValue interface {
	SetVariableAssignmentValue(value string)
	GetVariableAssignmentValue() *string
}

type IFBaseCoordinates interface {
	SetCoordinates(x, y int)
	GetCoordinates() (*int, *int)
	SetX(x int)
	GetX() *int
	SetY(y int)
	GetY() *int
}
