package impl

type IFBaseID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type IFBaseListenerID interface {
	SetListenerID(listenerID string)
	GetListenerID() STR_PTR
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
	GetLabel() STR_PTR
}

type IFBaseType interface {
	SetType(typ string)
	GetType() STR_PTR
}

type IFBaseConfig interface {
	SetConfig(config string)
	GetConfig() STR_PTR
}

type IFBaseKey interface {
	SetKey(key string)
	GetKey() STR_PTR
}

type IFBaseDefaultValue interface {
	SetDefaultValue(defaultValue string)
	GetDefaultValue() STR_PTR
}

type IFBaseScriptFormat interface {
	SetScriptFormat(format string)
	GetScriptFormat() STR_PTR
}

type IFBaseLocalVariableName interface {
	SetLocalVariableName(variable string)
	GetLocalVariableName() STR_PTR
}

type IFBaseVariableAssignmentValue interface {
	SetVariableAssignmentValue(value string)
	GetVariableAssignmentValue() STR_PTR
}

type IFBaseCoordinates interface {
	SetCoordinates(x, y int)
	GetCoordinates() (INT_PTR, INT_PTR)
	SetX(x int)
	GetX() INT_PTR
	SetY(y int)
	GetY() INT_PTR
}
