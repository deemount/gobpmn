package repository

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/deemount/gobpmn/examples"
	"github.com/deemount/gobpmn/models"
)

// BpmnFactory ...
type BpmnFactory interface {
	Create() (bpmnFactory, error)
	GetCurrentlyCreatedFile() string
}

// bpmnFactory ...
type bpmnFactory struct {
	// options
	Options BpmnOptions
	Def     *models.DefinitionsRepository
}

type BpmnFactoryOption func(o BpmnOptions) BpmnOptions

// NewBpmnFactory ...
func NewBpmnFactory(opt ...BpmnFactoryOption) BpmnFactory {
	options := BpmnOptions{}
	for _, o := range opt {
		options = o(options)
	}
	files, err := os.ReadDir("files/bpmn")
	if err != nil {
		log.Panic(err)
	}
	options.Counter = len(files)

	// count up
	if options.Counter == 0 {
		options.Counter = 1
	} else {
		options.Counter++
	}
	// set default name for bpmn-file
	options.CurrentFile = "diagram_" + fmt.Sprintf("%d", options.Counter)

	return &bpmnFactory{Options: options}
}

// run ...
func (factory *bpmnFactory) run() {

}

// set ...
func (factory *bpmnFactory) set() {

	model := examples.NewCollaborativeProcess()
	d := model.Create()
	factory.Def = d.Def()

}

// Create ...
func (factory bpmnFactory) Create() (bpmnFactory, error) {

	var err error

	factory.set()

	// create .bpmn
	err = factory.toBPMN()
	if err != nil {
		return bpmnFactory{}, err
	}

	return factory, nil

}

// GetCurrentlyCreatedFilename ...
func (factory bpmnFactory) GetCurrentlyCreatedFile() string {
	return factory.Options.CurrentFile
}

// toBPMN ...
func (factory *bpmnFactory) toBPMN() error {

	var err error

	// marshal xml to byte slice
	b, _ := xml.MarshalIndent(&factory.Def, " ", "  ")

	// create .bpmn file
	f, err := os.Create("files/bpmn/" + factory.Options.CurrentFile + ".bpmn")
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

	// create .json
	err = factory.toJSON()
	if err != nil {
		return err
	}

	return nil

}
