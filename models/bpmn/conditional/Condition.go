package conditional

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/impl"
)

// NewCondition ...
func NewCondition() ConditionRepository {
	return &Condition{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetConditionType ...
func (condition *Condition) SetConditionType() {
	condition.ConditionType = fmt.Sprint("bpmn:tFormalExpression")
}

// SetScriptFormat ...
func (condition *Condition) SetScriptFormat(format string) {
	condition.ScriptFormat = format
}

// SetScript ...
func (condition *Condition) SetScript(script string) {
	condition.Script = script
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetConditionType ...
func (condition Condition) GetConditionType() impl.STR_PTR {
	return &condition.ConditionType
}

// GetScriptFormat ...
func (condition Condition) GetScriptFormat() impl.STR_PTR {
	return &condition.ScriptFormat
}

// GetScript ...
func (condition Condition) GetScript() impl.STR_PTR {
	return &condition.Script
}
