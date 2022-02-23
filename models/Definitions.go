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
func (def *Definitions) SetBpmn() {
	def.Bpmn = "http://www.omg.org/spec/BPMN/20100524/MODEL"
}

// SetBpmndi ...
func (def *Definitions) SetBpmndi() {
	def.Bpmndi = "http://www.omg.org/spec/BPMN/20100524/DI"
}

// SetDC ...
func (def *Definitions) SetDC() {
	def.DC = "http://www.omg.org/spec/DD/20100524/DC"
}

// SetBioc ...
func (def *Definitions) SetBioc() {
	def.Bioc = "http://bpmn.io/schema/bpmn/biocolor/1.0"
}

// SetXSD ...
func (def *Definitions) SetXSD() {
	def.Xsd = "http://www.w3.org/2001/XMLSchema"
}

// SetDefinitionsID ...
func (def *Definitions) SetID(suffix string) {
	def.ID = "Definitions_" + suffix
}

// SetTargetNamespace ...
func (def *Definitions) SetTargetNamespace() {
	def.TargetNamespace = "http://bpmn.io/schema/bpmn"
}

/** Camunda **/

// SetCamundaSchema ...
func (def *Definitions) SetCamundaSchema() {
	def.CamundaSchema = "http://camunda.org/schema/1.0/bpmn"
}

// SetExporter ...
func (def *Definitions) SetExporter() {
	def.Exporter = "Camunda Modeler"
}

// SetExporterVersion ...
func (def *Definitions) SetExporterVersion() {
	def.ExporterVersion = "4.5.0"
}

/* Elements */

/** BPMN **/
func (def *Definitions) SetCollaboration() {
	def.Collaboration = make([]Collaboration, 1)
}

// SetProcess ...
func (def *Definitions) SetProcess(num int) {
	def.Process = make([]Process, num)
}

// SetMessage ...
func (def *Definitions) SetMessage(num int) {
	def.Msg = make([]Message, num)
}

// SetSignal ...
func (def *Definitions) SetSignal(num int) {
	def.Signal = make([]Signal, num)
}

// SetDiagram ...
func (def *Definitions) SetDiagram() {
	def.Diagram = make([]Diagram, 1)
}
