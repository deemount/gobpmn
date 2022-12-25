package core

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deemount/gobpmn/models/canvas"
	"github.com/deemount/gobpmn/models/marker"
	"github.com/deemount/gobpmn/models/pool"
	"github.com/deemount/gobpmn/models/process"
	"github.com/deemount/gobpmn/utils"
)

var (
	schemaBpmnModel           = "http://www.omg.org/spec/BPMN/20100524/MODEL"
	schemaBpmn20_1            = "http://www.omg.org/spec/BPMN/2.0/20100501/BPMN20.xsd"
	schemaBpmn20_2            = "https://www.omg.org/spec/BPMN/20100501/BPMN20.xsd"
	schemaBpmnDI              = "http://www.omg.org/spec/BPMN/20100524/DI"
	schemaOMGDI               = "http://www.omg.org/spec/DD/20100524/DI"
	schemaOMGDC               = "http://www.omg.org/spec/DD/20100524/DC"
	schemaBpmnIOBioColor      = "http://bpmn.io/schema/bpmn/biocolor/1.0"
	schemaW3XmlSchema         = "http://www.w3.org/2001/XMLSchema"
	schemaW3XmlSchemaInstance = "http://www.w3.org/2001/XMLSchema-instance"
)

// DefinitionsRepository ...
type DefinitionsRepository interface {
	CoreBase

	SetBpmn()
	SetBpmnDI()
	SetOmgDI()
	SetDC()
	SetOmgDC()
	SetBioc()
	SetXSD()
	SetXSI()
	SetXsiSchemaLocation1()
	SetXsiSchemaLocation2()

	SetTargetNamespace()
	SetCamundaSchema()
	SetZeebeSchema()
	SetModelerSchema()
	SetModelerExecutionPlatform()
	SetModelerExecutionPlatformVersion(version string)
	SetExporter()
	SetExporterVersion(version string)

	SetCollaboration()
	SetProcess(num int)
	SetCategory(num int)
	SetMessage(num int)
	SetSignal(num int)
	SetDiagram(num int)

	SetDefaultAttributes()

	GetCollaboration() *pool.Collaboration
	GetProcess(num int) *process.Process
	GetCategory(num int) *marker.Category
	GetMessage(num int) *marker.Message
	GetSignal(num int) *marker.Signal
	GetDiagram(num int) *canvas.Diagram
}

// Definitions represents the root element
type Definitions struct {
	XMLName                         xml.Name             `xml:"bpmn:definitions" json:"-"`
	Bpmn                            string               `xml:"xmlns:bpmn,attr" json:"-"`
	Xsd                             string               `xml:"xmlns:xsd,attr,omitempty" json:"-"`
	Xsi                             string               `xml:"xmlns:xsi,omitempty" json:"-"`
	XsiSchemaLocation               string               `xml:"xsi:schemaLocation,attr,omitempty" json:"-"`
	BpmnDI                          string               `xml:"xmlns:bpmndi,attr" json:"-"`
	OmgDI                           string               `xml:"xmlns:omgdi,attr,omitempty" json:"-"`
	DC                              string               `xml:"xmlns:dc,attr,omitempty" json:"-"`
	OmgDC                           string               `xml:"xmlns:omgdc,attr,omitempty" json:"-"`
	Bioc                            string               `xml:"xmlns:bioc,attr,omitempty" json:"-"`
	CamundaSchema                   string               `xml:"xmlns:camunda,attr,omitempty" json:"-"`
	Zeebe                           string               `xml:"xmlns:zeebe,omitempty" json:"-"`
	Modeler                         string               `xml:"xmlns:modeler,omitempty" json:"-"`
	ModelerExecutionPlatform        string               `xml:"modeler:executionPlatform,omitempty" json:"-"`
	ModelerExecutionPlatformVersion string               `xml:"modeler:executionPlatformVersion,omitempty" json:"-"`
	ID                              string               `xml:"id,attr" json:"id"`
	TargetNamespace                 string               `xml:"targetNamespace,attr" json:"-"`
	Exporter                        string               `xml:"exporter,attr,omitempty" json:"-"`
	ExporterVersion                 string               `xml:"exporterVersion,attr,omitempty" json:"-"`
	Collaboration                   []pool.Collaboration `xml:"bpmn:collaboration,omitempty" json:"collaboration"`
	Process                         []process.Process    `xml:"bpmn:process,omitempty" json:"process"`
	Category                        []marker.Category    `xml:"bpmn:category,omitempty" json:"category,omitempty"`
	Msg                             []marker.Message     `xml:"bpmn:message,omitempty" json:"message,omitempty"`
	Signal                          []marker.Signal      `xml:"bpmn:signal,omitempty" json:"signal,omitempty"`
	Diagram                         []canvas.Diagram     `xml:"bpmndi:BPMNDiagram,omitempty" json:"-"`
}

// TDefinitions ...
type TDefinitions struct {
	XMLName       xml.Name              `xml:"definitions" json:"-"`
	ID            string                `xml:"id,attr" json:"id"`
	Collaboration []pool.TCollaboration `xml:"collaboration,omitempty" json:"collaboration"`
	Process       []process.TProcess    `xml:"process,omitempty" json:"process"`
	Category      []marker.TCategory    `xml:"category,omitempty" json:"category,omitempty"`
	Msg           []marker.TMessage     `xml:"message,omitempty" json:"message,omitempty"`
	Signal        []marker.TSignal      `xml:"signal,omitempty" json:"signal,omitempty"`
}

func NewDefinitions() DefinitionsRepository {
	return &Definitions{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetBpmn ...
func (definitions *Definitions) SetBpmn() {
	definitions.Bpmn = schemaBpmnModel
}

// SetBpmnDI ...
func (definitions *Definitions) SetBpmnDI() {
	definitions.BpmnDI = schemaBpmnDI
}

// SetOmgDI ...
func (definitions *Definitions) SetOmgDI() {
	definitions.OmgDI = schemaOMGDI
}

// SetDC ...
func (definitions *Definitions) SetDC() {
	definitions.DC = schemaOMGDC
}

// SetOmgDC ...
func (definitions *Definitions) SetOmgDC() {
	definitions.OmgDC = schemaOMGDC
}

// SetBioc ...
func (definitions *Definitions) SetBioc() {
	definitions.Bioc = schemaBpmnIOBioColor
}

// SetXSD ...
func (definitions *Definitions) SetXSD() {
	definitions.Xsd = schemaW3XmlSchema
}

// SetXSI ...
func (definitions *Definitions) SetXSI() {
	definitions.Xsi = schemaW3XmlSchemaInstance
}

// SetXsiSchemaLocation ...
func (definitions *Definitions) SetXsiSchemaLocation1() {
	definitions.XsiSchemaLocation = fmt.Sprintf("%s %s", schemaBpmnModel, schemaBpmn20_1)
}

// SetXsiSchemaLocation ...
func (definitions *Definitions) SetXsiSchemaLocation2() {
	definitions.XsiSchemaLocation = fmt.Sprintf("%s %s", schemaBpmnModel, schemaBpmn20_2)
}

// SetID ...
func (definitions *Definitions) SetID(typ string, suffix interface{}) {
	switch typ {
	case "definitions":
		definitions.ID = fmt.Sprintf("Definitions_%v", suffix)
		break
	case "id":
		definitions.ID = fmt.Sprintf("%s", suffix)
	}
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

/*** Make Elements ***/

/** BPMN **/

// SetCollaboration ...
func (definitions *Definitions) SetCollaboration() {
	definitions.Collaboration = make([]pool.Collaboration, 1)
}

// SetProcess ...
func (definitions *Definitions) SetProcess(num int) {
	definitions.Process = make([]process.Process, num)
}

// SetCategory ...
func (definitions *Definitions) SetCategory(num int) {
	definitions.Category = make([]marker.Category, num)
}

// SetMessage ...
func (definitions *Definitions) SetMessage(num int) {
	definitions.Msg = make([]marker.Message, num)
}

// SetSignal ...
func (definitions *Definitions) SetSignal(num int) {
	definitions.Signal = make([]marker.Signal, num)
}

/** BPMNDI **/

// SetDiagram ...
func (definitions *Definitions) SetDiagram(num int) {
	definitions.Diagram = make([]canvas.Diagram, num)
}

/*
 * Default Settings
 */

// SetDefinitionsAttributes ...
func (definitions *Definitions) SetDefaultAttributes() {
	definitionsHash := utils.GenerateHash()
	definitions.SetBpmn()
	definitions.SetBpmnDI()
	definitions.SetDC()
	definitions.SetID("definitions", definitionsHash)
	definitions.SetTargetNamespace()
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (definitions Definitions) GetID() STR_PTR {
	return &definitions.ID
}

/* Elements */

/** BPMN **/

// GetCollaboration ...
func (definitions Definitions) GetCollaboration() *pool.Collaboration {
	return &definitions.Collaboration[0]
}

// GetProcess ...
func (definitions Definitions) GetProcess(num int) *process.Process {
	return &definitions.Process[num]
}

// GetCategory ...
func (definitions Definitions) GetCategory(num int) *marker.Category {
	return &definitions.Category[num]
}

// GetMessage ...
func (definitions Definitions) GetMessage(num int) *marker.Message {
	return &definitions.Msg[num]
}

// GetSignal ...
func (definitions Definitions) GetSignal(num int) *marker.Signal {
	return &definitions.Signal[num]
}

/** BPMNDI **/

// SetDiagram ...
func (definitions Definitions) GetDiagram(num int) *canvas.Diagram {
	return &definitions.Diagram[num]
}

/* Schema Fetcher */
func (definitions Definitions) GetCamundaSchema() {
	schema, err := utils.FetchSchema(definitions.CamundaSchema)
	if err != nil {
		log.Print(err)
	}
	log.Print(string(schema))
}
