package conditional

import "github.com/deemount/gobpmn/models/bpmn/impl"

// NewConditionExpression ...
func NewConditionExpression() ConditionExpressionRepository {
	return &ConditionExpression{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetConditionType ...
func (conditionExpression *ConditionExpression) SetConditionType() {
	conditionExpression.ConditionType = "bpmn:tFormalExpression"
}

// SetScriptFormat ...
func (conditionExpression *ConditionExpression) SetScriptFormat(language string) {
	conditionExpression.ScriptFormat = language
}

/* Content */

// SetScript ...
func (conditionExpression *ConditionExpression) SetScript(script string) {
	conditionExpression.Script = script
}

// SetExpression ...
func (conditionExpression *ConditionExpression) SetExpression(expression string) {
	conditionExpression.Expression = expression
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetConditionType ...
func (conditionExpression ConditionExpression) GetConditionType() impl.STR_PTR {
	return &conditionExpression.ConditionType
}

// GetScriptFormat ...
func (conditionExpression ConditionExpression) GetScriptFormat() impl.STR_PTR {
	return &conditionExpression.ScriptFormat
}

// GetScript ...
func (conditionExpression ConditionExpression) GetScript() impl.STR_PTR {
	return &conditionExpression.Script
}

// GetExpression ...
func (conditionExpression ConditionExpression) GetExpression() impl.STR_PTR {
	return &conditionExpression.Expression
}
