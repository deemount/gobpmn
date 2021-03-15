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
