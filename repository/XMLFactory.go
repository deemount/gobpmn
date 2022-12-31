package repository

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

// create .xml
/*err = bpmnFactory.toXML()
if err != nil {
	return err
}*/

// toXML ...
func (factory *bpmnFactory) toXML() error {

	var err error

	// marshal xml to byte slice
	b, _ := xml.MarshalIndent(&factory.Def, " ", "  ")

	// create .bpmn file
	f, err := os.Create("files/xml/" + factory.Options.CurrentFile + ".xml")
	if err != nil {
		return err
	}
	defer f.Close()

	// add xml header
	w := []byte(fmt.Sprintf("%v", xml.Header+string(b)))
	p := factory.toPlainXML(w)

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
func (factory *bpmnFactory) toPlainXML(b []byte) []byte {
	s := strings.Replace(string(b), "bpmn2:", "", -1)
	s = strings.Replace(s, "bpmn:", "", -1)
	s = strings.Replace(s, "bpmndi:", "", -1)
	s = strings.Replace(s, "camunda:", "", -1)
	s = strings.Replace(s, "di:", "", -1)
	s = strings.Replace(s, "dc:", "", -1)
	return []byte(fmt.Sprintf("%v", string(s)))
}
