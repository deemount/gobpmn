package camunda

import "fmt"

// CamundaScript ...
type CamundaScript struct {
	ScriptFormat string `xml:"scriptFormat,attr,omitempty"`
}

// NewCamundaScript ...
func NewCamundaScript() CamundaScriptRepository {
	return &CamundaScript{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetScriptFormat ...
func (cscript *CamundaScript) SetScriptFormat(format string) {
	cscript.ScriptFormat = fmt.Sprintf("%s", format)
}

/*
 * Default Getters
 */

/* Attributes */

// GetScriptFormat ...
func (cscript CamundaScript) GetScriptFormat() *string {
	return &cscript.ScriptFormat
}
