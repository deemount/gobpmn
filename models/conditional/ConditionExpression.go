package conditional

// ConditionExpressionRepository ...
type ConditionExpressionRepository interface {
	ConditionalScriptFormat
	ConditionalScript
	SetConditionType()
	GetConditionType() *string
	SetExpression(expression string)
	GetExpression() *string
}

// ConditionExpression ...
type ConditionExpression struct {
	ConditionType string `xml:"xsi:type,attr,omitempty" json:"conditionType,omitempty"`
	ScriptFormat  string `xml:"language,attr,omitempty" json:"language,omitempty"`
	Script        string `xml:",innerxml,omitempty" json:"script,omitempty"`
	Expression    string `xml:",innerxml,omitempty" json:"expression,omitempty"`
}

func NewConditionExpression() ConditionExpressionRepository {
	return &ConditionExpression{}
}

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
func (conditionExpression ConditionExpression) GetConditionType() *string {
	return &conditionExpression.ConditionType
}

// GetScriptFormat ...
func (conditionExpression ConditionExpression) GetScriptFormat() STR_PTR {
	return &conditionExpression.ScriptFormat
}

// GetScript ...
func (conditionExpression ConditionExpression) GetScript() STR_PTR {
	return &conditionExpression.Script
}

// GetExpression ...
func (conditionExpression ConditionExpression) GetExpression() *string {
	return &conditionExpression.Expression
}
