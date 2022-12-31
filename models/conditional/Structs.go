package conditional

// CompletionCondition ...
type CompletionCondition struct{}

// Condition ...
type Condition struct {
	ConditionType string `xml:"xsi:type,attr,omitempty" json:"conditionType,omitempty"`
	ScriptFormat  string `xml:"language,attr,omitempty" json:"language,omitempty"`
	Script        string `xml:",innerxml,omitempty" json:"script,omitempty"`
}

// TCondition ...
type TCondition struct {
	ConditionType string `xml:"type,attr,omitempty" json:"conditionType,omitempty"`
	ScriptFormat  string `xml:"language,attr,omitempty" json:"language,omitempty"`
	Script        string `xml:",innerxml,omitempty" json:"script,omitempty"`
}

// ConditionExpression ...
type ConditionExpression struct {
	ConditionType string `xml:"xsi:type,attr,omitempty" json:"conditionType,omitempty"`
	ScriptFormat  string `xml:"language,attr,omitempty" json:"language,omitempty"`
	Script        string `xml:",innerxml,omitempty" json:"script,omitempty"`
	Expression    string `xml:",innerxml,omitempty" json:"expression,omitempty"`
}

// TConditionExpression ...
type TConditionExpression struct {
	ConditionType string `xml:"type,attr,omitempty" json:"conditionType,omitempty"`
	ScriptFormat  string `xml:"language,attr,omitempty" json:"language,omitempty"`
	Script        string `xml:",innerxml,omitempty" json:"script,omitempty"`
	Expression    string `xml:",innerxml,omitempty" json:"expression,omitempty"`
}
