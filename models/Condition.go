package models

import "fmt"

// Condition ...
type Condition struct {
	ConditionType string `xml:"xsi:type,attr,omitempty"`
	ScriptFormat  string `xml:"language,attr,omitempty"`
	Script        string `xml:",innerxml,omitempty"`
}

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
