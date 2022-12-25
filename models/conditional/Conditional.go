package conditional

type ConditionalRepository interface {
	CompletionConditionRepository
	ConditionRepository
	ConditionExpressionRepository
}

type Conditional struct {
	ConditionalRepository
}

func NewConditional() ConditionalRepository {
	return &Conditional{}
}
