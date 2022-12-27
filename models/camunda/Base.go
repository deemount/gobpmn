package camunda

type STR_PTR *string

type CamundaBaseID interface {
	SetID(typ string, suffix string)
	GetID() STR_PTR
}

type CamundaBaseName interface {
	SetName(name string)
	GetName() STR_PTR
}

type CamundaBaseValue interface {
	SetValue(value string)
	GetValue() STR_PTR
}

type CamundaBaseLabel interface {
	SetLabel(label string)
	GetLabel() *string
}

type CamundaBaseType interface {
	SetType(typ string)
	GetType() *string
}

type CamundaBaseEvent interface {
	SetEvent(event string)
	GetEvent() STR_PTR
}

type CamundaBaseClass interface {
	SetClass(class string)
	GetClass() STR_PTR
}

type CamundaBaseScriptElements interface {
	SetCamundaScript()
	GetCamundaScript() *CamundaScript
}

type CamundaBase interface {
}
