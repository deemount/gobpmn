package conditional

import "github.com/deemount/gobpmn/models/bpmn/impl"

// ConditionalScriptFormat ...
type ConditionalScriptFormat interface {
	SetScriptFormat(format string)
	GetScriptFormat() impl.STR_PTR
}

// ConditionalScript ...
type ConditionalScript interface {
	SetScript(script string)
	GetScript() impl.STR_PTR
}

// ConditionalType ...
type ConditionalType interface {
	SetConditionType()
	GetConditionType() impl.STR_PTR
}

type ConditionalExpression interface {
	SetExpression(expression string)
	GetExpression() impl.STR_PTR
}

// CompletionConditionRepository ...
type CompletionConditionRepository interface{}

// ConditionRepository ...
type ConditionRepository interface {
	ConditionalScriptFormat
	ConditionalScript
	ConditionalType
}

// ConditionExpressionRepository ...
type ConditionExpressionRepository interface {
	ConditionalScriptFormat
	ConditionalScript
	ConditionalType
	ConditionalExpression
}
