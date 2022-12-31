package conditional

import "github.com/deemount/gobpmn/models/compulsion"

// ConditionalScriptFormat ...
type ConditionalScriptFormat interface {
	SetScriptFormat(format string)
	GetScriptFormat() compulsion.STR_PTR
}

// ConditionalScript ...
type ConditionalScript interface {
	SetScript(script string)
	GetScript() compulsion.STR_PTR
}

// ConditionalType ...
type ConditionalType interface {
	SetConditionType()
	GetConditionType() compulsion.STR_PTR
}

type ConditionalExpression interface {
	SetExpression(expression string)
	GetExpression() compulsion.STR_PTR
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
