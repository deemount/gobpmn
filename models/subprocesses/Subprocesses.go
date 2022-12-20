package subprocesses

type Type_Call_Activity CallActivityRepository

type SubprocessesRepository[C Type_Call_Activity] interface {
	CallActivityRepository() C
}

type Subprocesses[C Type_Call_Activity] struct {
	SubprocessesRepository[C]
}

func NewSubprocesses() SubprocessesRepository[Type_Call_Activity] {
	return &Subprocesses[Type_Call_Activity]{}
}
