package repository

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/deemount/gobpmn/examples/collaborative_process"
	"github.com/deemount/gobpmn/models/bpmn/core"
)

// BpmnFactory ...
type BpmnFactory interface {
	Build() (bpmnFactory, error)
	GetCurrentlyCreatedFile() string
}

// bpmnFactory ...
type bpmnFactory struct {
	// options
	Options BpmnOptions
	// using pointer to interface and struct for more flexibel modelling
	Repo core.DefinitionsRepository
	Def  *core.Definitions
}

type BpmnFactoryOption func(o BpmnOptions) BpmnOptions

// NewBpmnFactory ...
func NewBpmnFactory(opt ...BpmnFactoryOption) BpmnFactory {
	//
	options := BpmnOptions{}
	for _, o := range opt {
		options = o(options)
	}
	// read the dir for created bpmn files
	log.Println("factory: read directory")
	files, err := os.ReadDir("files/bpmn")
	if err != nil {
		log.Panic(err)
	}
	// set number and count up for each created file
	if options.Counter == 0 {
		options.Counter = len(files)
	} else {
		options.Counter += 1
	}
	// set default name for bpmn-file
	options.CurrentFile = "diagram_" + fmt.Sprintf("%d", options.Counter)
	log.Printf("factory: created filename %s.bpmn", options.CurrentFile)

	return &bpmnFactory{Options: options}
}

// set is a private method and is called inside Create().
// It calls the Create() method in the written business model,
// when it fits to the correct expectations
func (factory *bpmnFactory) set() {

	log.Println("factory: set model")

	// e.g. use interface pointer (no argument)

	//def := new(models.Definitions)
	//repositoryModel := examples.NewModel(def)

	// e.g. use struct pointer (no argument)

	// 1
	//repositoryModel := small_process.New().Build().Call()
	factory.Repo = collaborative_process.New().Build().Call()

	// 2
	// repositoryModel := examples.NewBlackBoxModel()
	//d := repositoryModel.Create()
	//factory.Repo = d.Def()

	// 3
	//repositoryModel := examples.NewSimpleModel001()
	//d := repositoryModel.Create()
	//factory.Repo = d.Def()

}

// Build...
func (factory bpmnFactory) Build() (bpmnFactory, error) {

	var err error

	factory.set()

	// create .bpmn
	log.Println("factory: create bpmn")
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
	b, _ := xml.MarshalIndent(&factory.Repo, " ", "  ")
	//b, _ := xml.MarshalIndent(&factory.Def, " ", "  ")

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
