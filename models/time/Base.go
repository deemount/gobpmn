package time

type STR_PTR *string

type TimeBaseDefintionType interface {
	SetTimerDefinitionType()
	GetTimerDefinitionType() *string
}

type TimeBaseDefinition interface {
	SetTimerDefinition(timerDefinition string)
	GetTimerDefinition() *string
}

type TimeBaseCoreElements interface {
	TimeBaseDefintionType
	TimeBaseDefinition
}
type TimeBase interface {
	TimeBaseCoreElements
}
