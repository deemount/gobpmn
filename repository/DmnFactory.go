package repository

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/deemount/gobpmn/models"
)

type DmnFactory struct {
	DmnFactoryCounter int
	Def               models.DefinitionsDMN
}

// NewDmnFactory ...
func NewDmnFactory() DmnFactory {
	files, _ := ioutil.ReadDir("files/")
	counter := len(files)
	if counter == 0 {
		counter++
	}
	return DmnFactory{DmnFactoryCounter: counter}
}

// Set ...
func (dm *DmnFactory) Set() {

}

// Create ...
func (dm *DmnFactory) Create() error {

	var err error

	// marshal xml to byte slice
	b, _ := xml.MarshalIndent(dm.Def, " ", "  ")

	// create file
	f, err := os.Create("files/diagram_" + fmt.Sprintf("%d", dm.DmnFactoryCounter) + ".dmn")
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
