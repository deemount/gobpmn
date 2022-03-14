package models

import (
	"encoding/xml"

	"github.com/deemount/gobpmn/utils"
)

// DefinitionsRepository ...
type DefinitionsRepository interface {
	SetID(suffix string)
	GetID() string
}

// Definitions represents the root element
type Definitions struct {
	XMLName                         xml.Name        `xml:"bpmn:definitions" json:"-"`
	Bpmn                            string          `xml:"xmlns:bpmn,attr" json:"-"`
	Xsd                             string          `xml:"xmlns:xsd,attr,omitempty" json:"-"`
	Xsi                             string          `xml:"xmlns:xsi,omitempty" json:"-"`
	XsiSchemaLocation               string          `xml:"xsi:schemaLocation,attr,omitempty" json:"-"`
	BpmnDI                          string          `xml:"xmlns:bpmndi,attr" json:"-"`
	OmgDI                           string          `xml:"xmlns:omgdi,attr,omitempty" json:"-"`
	DC                              string          `xml:"xmlns:dc,attr,omitempty" json:"-"`
	OmgDC                           string          `xml:"xmlns:omgdc,attr,omitempty" json:"-"`
	Bioc                            string          `xml:"xmlns:bioc,attr,omitempty" json:"-"`
	CamundaSchema                   string          `xml:"xmlns:camunda,attr,omitempty" json:"-"`
	Zeebe                           string          `xml:"xmlns:zeebe,omitempty" json:"-"`
	Modeler                         string          `xml:"xmlns:modeler,omitempty" json:"-"`
	ModelerExecutionPlatform        string          `xml:"modeler:executionPlatform,omitempty" json:"-"`
	ModelerExecutionPlatformVersion string          `xml:"modeler:executionPlatformVersion,omitempty" json:"-"`
	ID                              string          `xml:"id,attr" json:"id"`
	TargetNamespace                 string          `xml:"targetNamespace,attr" json:"-"`
	Exporter                        string          `xml:"exporter,attr,omitempty" json:"-"`
	ExporterVersion                 string          `xml:"exporterVersion,attr,omitempty" json:"-"`
	Collaboration                   []Collaboration `xml:"bpmn:collaboration,omitempty" json:"collaboration"`
	Process                         []Process       `xml:"bpmn:process,omitempty" json:"process"`
	Category                        []Category      `xml:"bpmn:category,omitempty" json:"category,omitempty"`
	Msg                             []Message       `xml:"bpmn:message,omitempty" json:"message,omitempty"`
	Signal                          []Signal        `xml:"bpmn:signal,omitempty" json:"signal,omitempty"`
	Diagram                         []Diagram       `xml:"bpmndi:BPMNDiagram,omitempty" json:"-"`
}

type TDefinitions struct {
	XMLName       xml.Name         `xml:"definitions" json:"-"`
	ID            string           `xml:"id,attr" json:"id"`
	Collaboration []TCollaboration `xml:"collaboration,omitempty" json:"collaboration"`
	Process       []TProcess       `xml:"process,omitempty" json:"process"`
	Category      []Category       `xml:"category,omitempty" json:"category,omitempty"`
	Msg           []Message        `xml:"message,omitempty" json:"message,omitempty"`
	Signal        []Signal         `xml:"signal,omitempty" json:"signal,omitempty"`
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetBpmn ...
func (definitions *Definitions) SetBpmn() {
	definitions.Bpmn = "http://www.omg.org/spec/BPMN/20100524/MODEL"
}

// SetBpmnDI ...
func (definitions *Definitions) SetBpmnDI() {
	definitions.BpmnDI = "http://www.omg.org/spec/BPMN/20100524/DI"
}

// SetOmgDI ...
func (definitions *Definitions) SetOmgDI() {
	definitions.OmgDI = "http://www.omg.org/spec/DD/20100524/DI"
}

// SetDC ...
func (definitions *Definitions) SetDC() {
	definitions.DC = "http://www.omg.org/spec/DD/20100524/DC"
}

// SetOmgDC ...
func (definitions *Definitions) SetOmgDC() {
	definitions.OmgDC = "http://www.omg.org/spec/DD/20100524/DC"
}

// SetBioc ...
func (definitions *Definitions) SetBioc() {
	definitions.Bioc = "http://bpmn.io/schema/bpmn/biocolor/1.0"
}

// SetXSD ...
func (definitions *Definitions) SetXSD() {
	definitions.Xsd = "http://www.w3.org/2001/XMLSchema"
}

// SetXSI ...
func (definitions *Definitions) SetXSI() {
	definitions.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
}

// SetXsiSchemaLocation ...
func (definitions *Definitions) SetXsiSchemaLocation() {
	definitions.XsiSchemaLocation = "http://www.omg.org/spec/BPMN/20100524/MODEL http://www.omg.org/spec/BPMN/2.0/20100501/BPMN20.xsd"
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

// SetZeebe ...
func (definitions *Definitions) SetZeebeSchema() {
	definitions.Zeebe = "http://camunda.org/schema/zeebe/1.0"
}

// SetModeler ...
func (definitions *Definitions) SetModelerSchema() {
	definitions.Modeler = "http://camunda.org/schema/modeler/1.0"
}

// SetModelerExecutionPlatform ...
func (definitions *Definitions) SetModelerExecutionPlatform() {
	definitions.Modeler = "Camunda Cloud"
}

// SetModelerExecutionPlatform ...
// can be ^1.1.0
func (definitions *Definitions) SetModelerExecutionPlatformVersion(version string) {
	definitions.Modeler = version
}

// SetExporter ...
func (definitions *Definitions) SetExporter() {
	definitions.Exporter = "Camunda Modeler"
}

// SetExporterVersion ...
// can be ^4.5.0 (start with this library)
func (definitions *Definitions) SetExporterVersion(version string) {
	definitions.ExporterVersion = version
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

/**
 * Default Settings
 */

// SetDefinitionsAttributes ...
func (definitions *Definitions) SetDefaultAttributes() {
	definitionsHash := utils.GenerateHash()
	definitions.SetBpmn()
	definitions.SetBpmnDI()
	definitions.SetDC()
	definitions.SetID(definitionsHash)
	definitions.SetTargetNamespace()
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (definitions Definitions) GetID() string {
	return definitions.ID
}
