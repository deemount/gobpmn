package main

import (
	"github.com/deemount/gobpmn/repository"
)

func main() {

	bpmnf := repository.NewBPMNF()
	bpmnf.Set()
	err := bpmnf.Create()
	if err != nil {
		panic(err)
	}

}
