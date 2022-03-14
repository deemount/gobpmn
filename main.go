package main

import (
	"context"
	_ "embed"

	"github.com/deemount/gobpmn/repository"
)

var bpmnFactory repository.BpmnFactory
var bpmnEngine repository.BpmnEngine
var process *repository.ProcessInfo

// init ...
func init() {
	bpmnFactory = repository.NewBpmnFactory()
	bpmnEngine = repository.NewBpmnEngine("goBPMN")
}

// main ...
func main() {

	err := bpmnFactory.Create()
	if err != nil {
		panic(err)
	}

	process, err = bpmnEngine.GetProcessInfo(context.Background(), bpmnFactory)
	if err != nil {
		panic(err)
	}

}
