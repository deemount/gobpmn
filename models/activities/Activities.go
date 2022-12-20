package activities

type Type_Task TaskRepository
type Type_User_Task UserTaskRepository
type Type_Receive_Task ReceiveTaskRepository
type Type_Manual_Task ManualTaskRepository
type Type_Script_Task ScriptTaskRepository
type Type_Send_Task SendTaskRepository
type Type_Business_Rule_Task BusinessRuleTaskRepository

type ActivitiesRepository[
	T Type_Task,
	U Type_User_Task,
	R Type_Receive_Task,
	M Type_Manual_Task,
	SC Type_Script_Task,
	SE Type_Send_Task,
	B Type_Business_Rule_Task] interface {
	TaskRepository() T
	UserTaskRepository() U
	ReceiveTaskRepository() R
	ManualTaskRepository() M
	ScriptTaskRepository() SC
	SendTaskRepository() SE
	BusinessRuleTaskRepository() B
}

type Activities[
	T Type_Task,
	U Type_User_Task,
	R Type_Receive_Task,
	M Type_Manual_Task,
	SC Type_Script_Task,
	SE Type_Send_Task,
	B Type_Business_Rule_Task] struct {
	ActivitiesRepository[T, U, R, M, SC, SE, B]
}

func NewActivities() ActivitiesRepository[
	Type_Task,
	Type_User_Task,
	Type_Receive_Task,
	Type_Manual_Task,
	Type_Script_Task,
	Type_Send_Task,
	Type_Business_Rule_Task] {
	return &Activities[
		Type_Task,
		Type_User_Task,
		Type_Receive_Task,
		Type_Manual_Task,
		Type_Script_Task,
		Type_Send_Task,
		Type_Business_Rule_Task]{}
}
