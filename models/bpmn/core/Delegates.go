package core

// SetMainElements ...
func SetMainElements(d DefinitionsRepository, numProcesses int) {
	d.SetCollaboration()
	d.SetProcess(numProcesses)
	d.SetDiagram(1)
}
