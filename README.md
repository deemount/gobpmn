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

```go
package main

import (
 "github.com/deemount/gobpmn/repository"
)

var bpmnFactory repository.BpmnFactory
var bpmnEngine repository.BpmnEngine

// init ...
func init() {
  log.Println("main: init factory")
  bpmnFactory = repository.NewBpmnFactory()
  log.Println("main: init engine")
  bpmnEngine = repository.NewBpmnEngine("goBPMN")
  log.Println("main: init process instance")
  instance = repository.NewProcessInstance(bpmnEngine)
}

// main ...
func main() {
  log.Println("main: bpmnFactory.Create")
  file, err := bpmnFactory.Create()
  if err != nil {
    panic(err)
  }

  log.Println("main: instance.GetProcessInfo")
  processInfo, err = instance.GetProcessInfo(context.Background(), file)
  if err != nil {
    panic(err)
  }

  log.Println("main: instance.Create")
  instanceInfo, err := instance.Create(processInfo.ProcessKey, nil)
  if err != nil {
    panic(err)
  }

  log.Println("main: instance.Run")
  if err = instance.Run(instanceInfo); err != nil {
    panic(err)
  } else {
    log.Println("main: do your stuff here")
  }
}
```

### Docker ###

```bash
docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

### To Do's ###

* Naming conventions for ID's
* Add more comments
* Add testing and benchmarks
* Add more configuration options
* Build message queue with RabbitMQ
* Add behaviours, command
* Add BPMN, DMN and element Factory

### Links ###

#### OMG.org ####

* BPMN Specification [https://www.omg.org/spec/BPMN]

#### Camunda ####

* BPMN [https://camunda.com/bpmn/]
* Schema [https://camunda.org/schema/1.0/]

##### Practical #####

* Naming ID's [https://camunda.com/best-practices/naming-technically-relevant-ids/]

### History ###

* 2021-02-21: Added more elements and tried a message queue
* 2021-03-15: Added more elements and delved deeper into structure
* 2021-01-26: Update
* 2021-01-24: First Upload/ Stable Realease
* 2021-01-23: Initialize Repository
