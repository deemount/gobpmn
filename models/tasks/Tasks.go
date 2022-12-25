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
