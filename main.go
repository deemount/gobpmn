package main

import (
	"context"
	_ "embed"
	"log"

	"github.com/deemount/gobpmn/repository"
)

var messageBroker repository.MessageBroker
var bpmnFactory repository.BpmnFactory
var bpmnEngine repository.BpmnEngine
var processInfo *repository.ProcessInfo
var instance repository.ProcessInstance

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
	file, err := bpmnFactory.Build()
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
		// Do your stuff here
	}

	// TESTING MESSAGE BROKER
	/*
		messageBroker.Dial()
		defer messageBroker.Connection.Close()
		messageBroker.OpenChannel()
		defer messageBroker.Channel.Close()
		messageBroker.DeclareQueue()
		messageBroker.Publish()
	*/
}
