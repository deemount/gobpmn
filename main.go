package main

import (
	"fmt"

	"github.com/deemount/gobpmn/repository"
)

func main() {

	bpmnf := repository.NewBPMNF()
	s := bpmnf.Create()
	fmt.Printf("%v", s)

}
