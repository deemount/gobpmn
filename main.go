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

	var fpn string = "Person 1"

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
		utils.HasFirstState(hasFirstState),
		utils.FirstParticipantName(fpn))

	bpmnf.Set()
	err := bpmnf.Create()
	if err != nil {
		panic(err)
	}

}
