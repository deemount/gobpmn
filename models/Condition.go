package models

import "fmt"

// Condition ...
type Condition struct {
	ConditionType string `xml:"xsi:type,attr,omitempty" json:"conditionType,omitempty"`
	ScriptFormat  string `xml:"language,attr,omitempty" json:"scriptFormat,omitempty"`
	Script        string `xml:",innerxml,omitempty" json:"script,omitempty"`
}

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
