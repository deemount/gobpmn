package repository

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/deemount/gobpmn/models"
)

type DMNF struct {
	DMNFCounter int
	Def         models.DefinitionsDMN
}

// NewBDMNF ...
func NewDMNF() DMNF {
	files, _ := ioutil.ReadDir("files/")
	counter := len(files)
	if counter == 0 {
		counter++
	}
	return DMNF{DMNFCounter: counter}
}

// Set ...
func (dm *DMNF) Set() {

}

// Create ...
func (dm *DMNF) Create() error {

	var err error

	// marshal xml to byte slice
	b, _ := xml.MarshalIndent(dm.Def, " ", "  ")

	// create file
	f, err := os.Create("files/diagram_" + fmt.Sprintf("%d", dm.DMNFCounter) + ".dmn")
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
