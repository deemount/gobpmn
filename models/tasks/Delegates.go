package tasks

// SetTask ...
func SetTask(e *Task, name string, hash ...string) {
	e.SetID("activity", hash[0])
	e.SetName(name)
	e.SetIncoming(1)
	e.GetIncoming(0).SetFlow(hash[1])
	e.SetOutgoing(1)
	e.GetOutgoing(0).SetFlow(hash[2])
}
