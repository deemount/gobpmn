package time

import "github.com/deemount/gobpmn/models/bpmn/impl"

type TimeBaseDefintionType interface {
	SetTimerDefinitionType()
	GetTimerDefinitionType() impl.STR_PTR
}

type TimeBaseDefinition interface {
	SetTimerDefinition(timerDefinition string)
	GetTimerDefinition() impl.STR_PTR
}

type TimeBaseCoreElements interface {
	TimeBaseDefintionType
	TimeBaseDefinition
}
type TimeBase interface {
	TimeBaseCoreElements
}

/*
 * @Repositories
 */

// TimeCycle ...
type TimeCycleRepository interface{ TimeBase }

// TimeDateRepository ...
type TimeDateRepository interface{ TimeBase }

// TimeDurationRepository ...
type TimeDurationRepository interface{ TimeBase }
