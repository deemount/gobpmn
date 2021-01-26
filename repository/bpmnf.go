package repository

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/deemount/gobpmn/models"
	"github.com/deemount/gobpmn/utils"
)

var counter int = 0

// BPMNF ...
type BPMNF struct {
	Def models.Definitions
}

// NewBPMNF ...
func NewBPMNF() BPMNF {
	files, _ := ioutil.ReadDir("files/")
	counter = len(files)
	if counter == 0 {
		counter++
	}
	return BPMNF{}
}

// Set ...
func (bpm *BPMNF) Set() {

	defHash := utils.GenerateHash()

	// set namespaces
	bpm.Def.SetBpmn()
	//d.SetXSD()
	bpm.Def.SetBpmndi()

	//
	bpm.Def.SetDefinitionsID(defHash)
	bpm.Def.SetTargetNamespace()

	// set exporter
	bpm.Def.SetExporter()
	bpm.Def.SetExporterVersion()

	// set process
	procHash := utils.GenerateHash()
	cVersionTag := ""
	name := ""
	bpm.Def.Proc.SetID(procHash)
	bpm.Def.Proc.SetName(name)
	bpm.Def.Proc.SetIsExecutable()
	bpm.Def.Proc.SetCamundaVersionTag(cVersionTag)

	// set diagram
	var n int64 = 1
	bpm.Def.Diagram.SetID(n)

	// set plane
	bpm.Def.Diagram.Plane.SetID(n)
	bpm.Def.Diagram.Plane.SetElement(procHash)

}

// Create ...
func (bpm *BPMNF) Create() error {

	var err error

	// marshal xml to byte slice
	b, _ := xml.MarshalIndent(bpm.Def, " ", "  ")

	// create file
	f, err := os.Create("files/diagram_" + fmt.Sprintf("%d", counter) + ".bpmn")
	if err != nil {
		return err
	}
	defer f.Close()

	// add xml header
	w := []byte(fmt.Sprintf("%v", xml.Header+string(b)))

	// write bytes to file
	_, err = f.Write(w)
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}

	return nil

}
