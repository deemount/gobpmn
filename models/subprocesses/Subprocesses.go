package subprocesses

type SubprocessesRepository interface {
	CallActivityRepository
	SubProcessRepository
	AdHocSubProcessRepository
	TransactionRepository
}

type Subprocesses struct {
	SubprocessesRepository
}

func NewSubprocesses() SubprocessesRepository {
	return &Subprocesses{}
}
