package attributes

type ConditionExpressionRepository interface {
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
