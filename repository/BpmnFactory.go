package repository

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gocarina/gocsv"

	"github.com/deemount/gobpmn/examples"
	"github.com/deemount/gobpmn/models"
)

// BpmnFactoryRepository ...
type BpmnFactoryRepository interface {
	Create() error
}

// BpmnFactory ...
type BpmnFactory struct {
	// options
	Options BpmnOptions
	Def     models.Definitions
}

type BpmnFactoryOption func(o BpmnOptions) BpmnOptions

// NewBpmnFactory ...
func NewBpmnFactory(opt ...BpmnFactoryOption) BpmnFactory {
	options := BpmnOptions{}
	for _, o := range opt {
		options = o(options)
	}
	files, _ := ioutil.ReadDir("files/bpmn")
	options.Counter = len(files)

	// count up
	if options.Counter == 0 {
		options.Counter = 1
	} else {
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
		modelName = fmt.Sprintf("00%d", 5)
	}

	/* set & create model */
	switch modelName {
	case "001":
		model := examples.NewSimpleModel001(def)
		model.Create()
	case "002":
		model := examples.NewSimpleModel002(def)
		model.Create()
	case "003":
		model := examples.NewSimpleModel003(def)
		model.Create()
	case "004":
		model := examples.NewSimpleModel004(def)
		model.Create()
	case "005":
		model := examples.NewCollaborativeProcess(def)
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

	// create .csv
	/*err = bpmnFactory.toCSV()
	if err != nil {
		return err
	}*/
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
	f, err := os.Create("files/bpmn/" + bpmnFactory.Options.CurrentFile + ".bpmn")
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
	f, err := os.Create("files/xml/" + bpmnFactory.Options.CurrentFile + ".xml")
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
	f, err := os.Create("files/json/" + bpmnFactory.Options.CurrentFile + ".json")
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

// toCSV ...
func (bpmnFactory *BpmnFactory) toCSV() error {
	var err error

	f, err := os.Create("files/csv/" + bpmnFactory.Options.CurrentFile + ".csv")
	if err != nil {
		return err
	}
	defer f.Close()

	err = gocsv.MarshalFile(&bpmnFactory.Def, f)
	if err != nil {
		panic(err)
	}

	return nil

}
