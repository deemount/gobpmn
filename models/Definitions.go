package models

import (
	"encoding/xml"
)

// Definitions ...
type Definitions struct {
	XMLName         xml.Name `xml:"bpmn:definitions"`
	Bpmn            string   `xml:"xmlns:bpmn,attr"`
	Xsd             string   `xml:"xmlns:xsd,attr,omitempty"`
	Bpmndi          string   `xml:"xmlns:bpmndi,attr"`
	CamundaSchema   string   `xml:"xmlns:camunda,attr,omitempty"`
	ID              string   `xml:"id,attr"`
	TargetNamespace string   `xml:"targetNamespace,attr"`
	Exporter        string   `xml:"exporter,attr"`
	ExporterVersion string   `xml:"exporterVersion,attr"`
	Proc            Process  `xml:"bpmn:process"`
	Diagram         Diagram  `xml:"bpmndi:BPMNDiagram"`
}

// SetBpmn ...
func (def *Definitions) SetBpmn() {
	def.Bpmn = "http://www.omg.org/spec/BPMN/20100524/MODEL"
}

// SetBpmndi ...
func (def *Definitions) SetBpmndi() {
	def.Bpmndi = "http://www.omg.org/spec/BPMN/20100524/DI"
}

// SetCamundaSchema ...
func (def *Definitions) SetCamundaSchema() {
	def.CamundaSchema = "http://camunda.org/schema/1.0/bpmn"
}

// SetXSD ...
func (def *Definitions) SetXSD() {
	def.Xsd = "http://www.w3.org/2001/XMLSchema"
}

// SetDefinitionsID ...
func (def *Definitions) SetDefinitionsID(suffix string) {
	def.ID = "Definitions_" + suffix
}

// SetTargetNamespace ...
func (def *Definitions) SetTargetNamespace() {
	def.TargetNamespace = "http://bpmn.io/schema/bpmn"
}

// SetExporter ...
func (def *Definitions) SetExporter() {
	def.Exporter = "Camunda Modeler"
}

// SetExporterVersion ...
func (def *Definitions) SetExporterVersion() {
	def.ExporterVersion = "4.5.0"
}
