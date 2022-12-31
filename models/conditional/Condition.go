package conditional

import (
	"fmt"

	"github.com/deemount/gobpmn/models/compulsion"
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
func (condition Condition) GetConditionType() compulsion.STR_PTR {
	return &condition.ConditionType
}

// GetScriptFormat ...
func (condition Condition) GetScriptFormat() compulsion.STR_PTR {
	return &condition.ScriptFormat
}

// GetScript ...
func (condition Condition) GetScript() compulsion.STR_PTR {
	return &condition.Script
}
