package definitions

import (
	"fmt"
)

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "cancel":
		r = fmt.Sprintf("CancelEventDefinition_%v", suffix)
		break
	case "conditional":
		r = fmt.Sprintf("ConditionalEventDefinition_%v", suffix)
		break
	case "error":
		r = fmt.Sprintf("ErrorEventDefinition_%v", suffix)
		break
	case "escalation":
		r = fmt.Sprintf("EscalationEventDefinition_%v", suffix)
		break
	case "link":
		r = fmt.Sprintf("LinkEventDefinition_%v", suffix)
		break
	case "message":
		r = fmt.Sprintf("MessageEventDefinition_%v", suffix)
		break
	case "signal":
		r = fmt.Sprintf("SignalEventDefinition_%v", suffix)
		break
	case "terminate":
		r = fmt.Sprintf("TerminateEventDefinition_%v", suffix)
		break
	case "timer":
		r = fmt.Sprintf("TimerEventDefinition_%v", suffix)
		break
	case "id":
		r = fmt.Sprintf("%s", suffix)
		break
	}
	return r
}

// SetTimerEventDefinition ...
// t *TimerEventDefinition, timerDefinition string, hash string
func SetTimerEventDefinition(p DelegateParameter) {
	p.TED.SetID("timer", p.H[0])
	p.TED.SetTimeDuration()
	p.TED.GetTimeDuration().SetTimerDefinitionType()
	p.TED.GetTimeDuration().SetTimerDefinition(p.TD)
}
