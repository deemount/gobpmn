package main

import (
	"log"

	"github.com/deemount/gobpmn"
)

type Process struct {
	Def            gobpmn.Repository
	IsExecutable   bool
	Process        gobpmn.BPMN
	StartEvent     gobpmn.BPMN
	FromStartEvent gobpmn.BPMN
	Task           gobpmn.BPMN
	FromTask       gobpmn.BPMN
	EndEvent       gobpmn.BPMN
}

func (p Process) GetDefinitions() gobpmn.Repository {
	return p.Def
}

func main() {
	_, err := gobpmn.FromStruct(Process{})
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return
	}
}
