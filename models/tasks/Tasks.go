package tasks

type TasksRepository interface {
	TaskRepository
	UserTaskRepository
	ReceiveTaskRepository
	ManualTaskRepository
	ScriptTaskRepository
	SendTaskRepository
	BusinessRuleTaskRepository
}

type Tasks struct {
	TasksRepository
}

func NewTasks() TasksRepository {
	return &Tasks{}
}

func (tasks *Tasks) SetID(typ string, suffix interface{}) {
	tasks.TasksRepository.SetID(typ, suffix)
}

func (tasks *Tasks) SetName(name string) {
	tasks.TasksRepository.SetName(name)
}
