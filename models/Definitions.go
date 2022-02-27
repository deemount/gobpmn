package models

import (
	"encoding/xml"
)

// Definitions represents the root element
type Definitions struct {
	XMLName         xml.Name        `xml:"bpmn:definitions"`
	Bpmn            string          `xml:"xmlns:bpmn,attr"`
	Xsd             string          `xml:"xmlns:xsd,attr,omitempty"`
	Bpmndi          string          `xml:"xmlns:bpmndi,attr"`
	DC              string          `xml:"xmlns:dc,attr,omitempty"`
	Bioc            string          `xml:"xmlns:bioc,attr,omitempty"`
	CamundaSchema   string          `xml:"xmlns:camunda,attr,omitempty"`
	ID              string          `xml:"id,attr"`
	TargetNamespace string          `xml:"targetNamespace,attr"`
	Exporter        string          `xml:"exporter,attr"`
	ExporterVersion string          `xml:"exporterVersion,attr"`
	Collaboration   []Collaboration `xml:"bpmn:collaboration,omitempty"`
	Process         []Process       `xml:"bpmn:process,omitempty"`
	Category        []Category      `xml:"bpmn:category,omitempty"`
	Msg             []Message       `xml:"bpmn:message,omitempty"`
	Signal          []Signal        `xml:"bpmn:signal,omitempty"`
	Diagram         []Diagram       `xml:"bpmndi:BPMNDiagram,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetBpmn ...
func (definitions *Definitions) SetBpmn() {
	definitions.Bpmn = "http://www.omg.org/spec/BPMN/20100524/MODEL"
}

// SetBpmndi ...
func (definitions *Definitions) SetBpmndi() {
	definitions.Bpmndi = "http://www.omg.org/spec/BPMN/20100524/DI"
}

// SetDC ...
func (definitions *Definitions) SetDC() {
	definitions.DC = "http://www.omg.org/spec/DD/20100524/DC"
}

// SetBioc ...
func (definitions *Definitions) SetBioc() {
	definitions.Bioc = "http://bpmn.io/schema/bpmn/biocolor/1.0"
}

// SetXSD ...
func (definitions *Definitions) SetXSD() {
	definitions.Xsd = "http://www.w3.org/2001/XMLSchema"
}

// SetDefinitionsID ...
func (definitions *Definitions) SetID(suffix string) {
	definitions.ID = "Definitions_" + suffix
}

// SetTargetNamespace ...
func (definitions *Definitions) SetTargetNamespace() {
	definitions.TargetNamespace = "http://bpmn.io/schema/bpmn"
}

/** Camunda **/

// SetCamundaSchema ...
func (definitions *Definitions) SetCamundaSchema() {
	definitions.CamundaSchema = "http://camunda.org/schema/1.0/bpmn"
}

// SetExporter ...
func (definitions *Definitions) SetExporter() {
	definitions.Exporter = "Camunda Modeler"
}

// SetExporterVersion ...
func (definitions *Definitions) SetExporterVersion() {
	definitions.ExporterVersion = "4.5.0"
}

/* Elements */

/** BPMN **/

// SetCollaboration ...
func (definitions *Definitions) SetCollaboration() {
	definitions.Collaboration = make([]Collaboration, 1)
}

// SetProcess ...
func (definitions *Definitions) SetProcess(num int) {
	definitions.Process = make([]Process, num)
}

// SetMessage ...
func (definitions *Definitions) SetCategory(num int) {
	definitions.Category = make([]Category, num)
}

// SetMessage ...
func (definitions *Definitions) SetMessage(num int) {
	definitions.Msg = make([]Message, num)
}

// SetSignal ...
func (definitions *Definitions) SetSignal(num int) {
	definitions.Signal = make([]Signal, num)
}

/** BPMNDI **/

// SetDiagram ...
func (definitions *Definitions) SetDiagram() {
	definitions.Diagram = make([]Diagram, 1)
}
