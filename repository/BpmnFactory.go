package repository

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/deemount/gobpmn/examples"
	"github.com/deemount/gobpmn/models"
)

type Options struct {
	Counter     int
	CurrentFile string
}

// BpmnFactoryRepository ...
type BpmnFactoryRepository interface {
	Create() error
}

// BpmnFactory ...
type BpmnFactory struct {
	// options
	Options Options
	Def     models.Definitions
}

type BpmnFactoryOption func(o Options) Options

// NewBpmnFactory ...
func NewBpmnFactory(opt ...BpmnFactoryOption) BpmnFactory {
	options := Options{}
	for _, o := range opt {
		options = o(options)
	}
	files, _ := ioutil.ReadDir("files/")
	options.Counter = len(files)
	fmt.Println(options.Counter)
	// count up
	if options.Counter == 0 {
		options.Counter++
	}
	options.CurrentFile = "diagram_" + fmt.Sprintf("%d", options.Counter)
	return BpmnFactory{Options: options}
}

// Set ...
func (bpmnFactory *BpmnFactory) set(modelName interface{}) {

	/* set elements */
	def := &bpmnFactory.Def

	if modelName == nil {
		modelName = fmt.Sprintf("simple-model-00%d", 2)
	}

	/* set & create model */
	switch modelName {
	case "simple-model-001":
		model := examples.NewSimpleModel001(def)
		model.Create()
	case "simple-model-002":
		model := examples.NewSimpleModel002(def)
		model.Create()
	case "simple-model-003":
		model := examples.NewSimpleModel003(def)
		model.Create()
	case "simple-model-004":
		model := examples.NewSimpleModel003(def)
		model.Create()
	}

}

// Create ...
func (bpmnFactory *BpmnFactory) Create() error {

	var err error

	bpmnFactory.set(nil)

	// create .bpmn
	err = bpmnFactory.toBPMN()
	if err != nil {
		return err
	}

	// create .json
	err = bpmnFactory.toJSON()
	if err != nil {
		return err
	}

	// create .xml
	err = bpmnFactory.toXML()
	if err != nil {
		return err
	}

	return nil

}

// GetCurrentlyCreatedFilename ...
func (bpmnFactory BpmnFactory) GetCurrentlyCreatedFile() string {
	return bpmnFactory.Options.CurrentFile
}

// toBPMN ...
func (bpmnFactory *BpmnFactory) toBPMN() error {

	var err error

	// marshal xml to byte slice
	b, _ := xml.MarshalIndent(&bpmnFactory.Def, " ", "  ")

	// create .bpmn file
	f, err := os.Create("files/" + bpmnFactory.Options.CurrentFile + ".bpmn")
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

// toXML ...
func (bpmnFactory *BpmnFactory) toXML() error {

	var err error

	// marshal xml to byte slice
	b, _ := xml.MarshalIndent(&bpmnFactory.Def, " ", "  ")

	// create .bpmn file
	f, err := os.Create("xml/" + bpmnFactory.Options.CurrentFile + ".xml")
	if err != nil {
		return err
	}
	defer f.Close()

	// add xml header
	w := []byte(fmt.Sprintf("%v", xml.Header+string(b)))
	p := bpmnFactory.toPlainXML(w)

	// write bytes to file
	_, err = f.Write(p)
	if err != nil {
		return err
	}

	err = f.Sync()
	if err != nil {
		return err
	}

	return nil

}

// toPlainXML ...
func (bpmnFactory *BpmnFactory) toPlainXML(b []byte) []byte {
	s := strings.Replace(string(b), "bpmn2:", "", -1)
	s = strings.Replace(s, "bpmn:", "", -1)
	s = strings.Replace(s, "bpmndi:", "", -1)
	s = strings.Replace(s, "camunda:", "", -1)
	s = strings.Replace(s, "di:", "", -1)
	s = strings.Replace(s, "dc:", "", -1)
	return []byte(fmt.Sprintf("%v", string(s)))
}

// toJSON ...
func (bpmnFactory *BpmnFactory) toJSON() error {
	var err error

	// marshal json to byte slice
	b, _ := json.MarshalIndent(&bpmnFactory.Def, " ", "  ")

	// create .json file
	f, err := os.Create("json/" + bpmnFactory.Options.CurrentFile + ".json")
	if err != nil {
		return err
	}
	defer f.Close()

	// write bytes to file
	_, err = f.Write(b)
	if err != nil {
		return err
	}

	err = f.Sync()
	if err != nil {
		return err
	}

	return nil

}
