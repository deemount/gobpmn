package dmn

import "encoding/xml"

// DefinitionsDMN represents the root element
type DefinitionsDMN struct {
	XMLName   xml.Name   `xml:"definitions"`
	Xmlns     string     `xml:"xmlns,attr"`
	Dmndi     string     `xml:"xmlns:dmndi,attr"`
	DC        string     `xml:"xmlns:dc,attr,omitempty"`
	ID        string     `xml:"id,attr"`
	Name      string     `xml:"name,attr"`
	Namespace string     `xml:"namespace,attr"`
	Diagram   DiagramDMN `xml:"dmndi:DMNDI"`
}

// SetBpmn ...
func (defDMN *DefinitionsDMN) SetXmlns() {
	defDMN.Xmlns = "http://www.omg.org/spec/DMN/20191111/MODEL/"
}

// SetBpmndi ...
func (defDMN *DefinitionsDMN) SetBpmndi() {
	defDMN.Dmndi = "http://www.omg.org/spec/DMN/20191111/DMNDI/"
}

// SetDC ...
func (defDMN *DefinitionsDMN) SetDC() {
	defDMN.DC = "http://www.omg.org/spec/DMN/20180521/DC/"
}

// SetDefinitionsID ...
func (defDMN *DefinitionsDMN) SetDefinitionsID(suffix string) {
	defDMN.ID = "Definitions_" + suffix
}

// SetName ...
func (defDMN *DefinitionsDMN) SetName(name string) {
	defDMN.Name = name
}

// SetName ...
func (defDMN *DefinitionsDMN) SetNamespace(namespace string) {
	defDMN.Namespace = namespace
}
