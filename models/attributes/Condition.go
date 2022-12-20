package attributes

import "fmt"

// ConditionRepository ...
type ConditionRepository interface {
	SetConditionType()
	SetScriptFormat(format string)
	SetScript(script string)

	GetConditionType() *string
	GetScriptFormat() *string
	GetScript() *string
}

// Condition ...
type Condition struct {
	ConditionType string `xml:"xsi:type,attr,omitempty" json:"conditionType,omitempty"`
	ScriptFormat  string `xml:"language,attr,omitempty" json:"language,omitempty"`
	Script        string `xml:",innerxml,omitempty" json:"script,omitempty"`
}

func NewCondition() ConditionRepository {
	return &Condition{}
}

/**
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

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetConditionType ...
func (condition *Condition) GetConditionType() *string {
	return &condition.ConditionType
}

// GetScriptFormat ...
func (condition *Condition) GetScriptFormat() *string {
	return &condition.ScriptFormat
}

// GetScript ...
func (condition *Condition) GetScript() *string {
	return &condition.Script
}
