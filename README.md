# gobpmn #

Business Process Model Notation 2 with Go

Author: Salvatore Gonda

## Introduction ##

This is part of my journey through BPMN. To teach myself, I opened the Camunda Modeler, switch to XML tab and modelled the xsd template in Go.

### Status ###

* no refactoring actually
* still in development
* creates a simple diagram_*.bpmn file

### Example ###

Create a simple bpmn-file with start and intermediate throw event

```
package main

import (
	"io/ioutil"

	"github.com/deemount/gobpmn/repository"
	"github.com/deemount/gobpmn/utils"
)

func main() {

	files, _ := ioutil.ReadDir("files/")
	var n int64 = 1
	var c int64 = int64(len(files))

	var dh string = utils.GenerateHash()
	var ph string = utils.GenerateHash()
	var fh string = utils.GenerateHash()
	var eh string = utils.GenerateHash()

	var cvt string = ""
	var name string = ""

	// start event option
	var hasStartEvent = true
	var hasFirstState = true
	var fk string = "create"

	// collaboration option
	var hasCollab = true
	var ch string = utils.GenerateHash()
	var pp string = utils.GenerateHash()

	bpmnf := repository.NewBPMNF(
		utils.N(n),
		utils.Counter(c),
		utils.DefHash(dh),
		utils.ProcHash(ph),
		utils.FlowHash(fh),
		utils.EventHash(eh),
		utils.CVersionTag(cvt),
		utils.Name(name),
		utils.FormKey(fk),
		utils.CollaboHash(ch),
		utils.PartHash(pp),
		utils.HasCollab(hasCollab),
		utils.HasStartEvent(hasStartEvent),
		utils.HasFirstState(hasFirstState))

	bpmnf.Set()
	err := bpmnf.Create()
	if err != nil {
		panic(err)
	}

}
```

### To Do's ###

* Naming conventions for ID's
* Add more comments
* Add testing and benchmarks
* Add configuration options
* Building architecture

### Links ###

#### OMG.org ####

* BPMN Specification [https://www.omg.org/spec/BPMN]

#### Camunda ####

* BPMN [https://camunda.com/bpmn/]
* Schema [https://camunda.org/schema/1.0/]

##### Practical #####

* Naming ID's [https://camunda.com/best-practices/naming-technically-relevant-ids/]

### History ###

* 2021-03-15: Added more elements and delved deeper into structure
* 2021-01-26: Update
* 2021-01-24: First Upload/ Stable Realease
* 2021-01-23: Initialize Repository
