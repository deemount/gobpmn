package time

type TimeBaseDefintionType interface {
	SetTimerDefinitionType()
	GetTimerDefinitionType() STR_PTR
}

type TimeBaseDefinition interface {
	SetTimerDefinition(timerDefinition string)
	GetTimerDefinition() STR_PTR
}

type TimeBaseCoreElements interface {
	TimeBaseDefintionType
	TimeBaseDefinition
}
type TimeBase interface {
	TimeBaseCoreElements
}
