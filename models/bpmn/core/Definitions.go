package core

import (
	"fmt"
	"log"

	"github.com/deemount/gobpmn/models/bpmn/canvas"
<<<<<<< HEAD
	"github.com/deemount/gobpmn/models/bpmn/collaboration"
	"github.com/deemount/gobpmn/models/bpmn/events"
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/marker"
=======
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/marker"
	"github.com/deemount/gobpmn/models/bpmn/pool"
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
	"github.com/deemount/gobpmn/models/bpmn/process"
	"github.com/deemount/gobpmn/utils"
)

var (
	schemaBpmnModel           = "http://www.omg.org/spec/BPMN/20100524/MODEL"
	schemaBpmn20              = "http://www.omg.org/spec/BPMN/2.0/20100501/BPMN20.xsd"
	schemaBpmn20_HTTPS        = "https://www.omg.org/spec/BPMN/20100501/BPMN20.xsd"
	schemaBpmnDI              = "http://www.omg.org/spec/BPMN/20100524/DI"
	schemaOMGDI               = "http://www.omg.org/spec/DD/20100524/DI"
	schemaOMGDC               = "http://www.omg.org/spec/DD/20100524/DC"
	schemaBpmnIOSchema        = "http://bpmn.io/schema/bpmn"
	schemaBpmnIOBioColor      = "http://bpmn.io/schema/bpmn/biocolor/1.0"
	schemaCamunda             = "http://camunda.org/schema/1.0/bpmn"
	schemaCamundaZebee        = "http://camunda.org/schema/zeebe/1.0"
	schemaCamundaModeler      = "http://camunda.org/schema/modeler/1.0"
	schemaW3XmlSchema         = "http://www.w3.org/2001/XMLSchema"
	schemaW3XmlSchemaInstance = "http://www.w3.org/2001/XMLSchema-instance"
)

// NewDefinitions ...
func NewDefinitions(p interface{}) DefinitionsRepository {
	return &Definitions{}
}

/*
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
func (definitions *Definitions) SetXsiSchemaLocation() {
	definitions.XsiSchemaLocation = fmt.Sprintf("%s %s", schemaBpmnModel, schemaBpmn20)
}

// SetXsiSchemaLocation ...
func (definitions *Definitions) SetXsiSchemaLocationHTTPS() {
	definitions.XsiSchemaLocation = fmt.Sprintf("%s %s", schemaBpmnModel, schemaBpmn20_HTTPS)
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
	definitions.TargetNamespace = schemaBpmnIOSchema
}

/** Camunda **/

// SetCamundaSchema ...
func (definitions *Definitions) SetCamundaSchema() {
	definitions.CamundaSchema = schemaCamunda
}

// SetZeebe ...
func (definitions *Definitions) SetZeebeSchema() {
	definitions.Zeebe = schemaCamundaZebee
}

// SetModeler ...
func (definitions *Definitions) SetModelerSchema() {
	definitions.Modeler = schemaCamundaModeler
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
	definitions.Collaboration = make(collaboration.COLLABORATION_SLC, 1)
}

// SetProcess ...
func (definitions *Definitions) SetProcess(num int) {
	definitions.Process = make(process.PROCESS_SLC, num)
}

// SetCategory ...
func (definitions *Definitions) SetCategory(num int) {
	definitions.Category = make([]marker.Category, num)
}

/*** Events ***/

// SetMessage ...
func (definitions *Definitions) SetMessage(num int) {
<<<<<<< HEAD
	definitions.Message = make(events.MESSAGE_SLC, num)
=======
	definitions.Msg = make([]elements.Message, num)
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}

// SetSignal ...
func (definitions *Definitions) SetSignal(num int) {
<<<<<<< HEAD
	definitions.Signal = make(events.SIGNAL_SLC, num)
=======
	definitions.Signal = make([]elements.Signal, num)
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}

/** BPMNDI **/

// SetDiagram ...
func (definitions *Definitions) SetDiagram(num int) {
	definitions.Diagram = make(canvas.DIAGRAM_SLC, num)
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

// SetMainElements ...
func (definitions *Definitions) SetMainElements(num int) {
	definitions.SetCollaboration()
	definitions.SetProcess(num)
	definitions.SetDiagram(1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (definitions Definitions) GetID() impl.STR_PTR {
	return &definitions.ID
}

/* Elements */

/** BPMN **/

// GetCollaboration ...
func (definitions Definitions) GetCollaboration() collaboration.COLLABORATION_PTR {
	return &definitions.Collaboration[0]
}

// GetProcess ...
func (definitions Definitions) GetProcess(num int) process.PROCESS_PTR {
	return &definitions.Process[num]
}

// GetCategory ...
func (definitions Definitions) GetCategory(num int) *marker.Category {
	return &definitions.Category[num]
}

/*** Events ***/

// GetMessage ...
func (definitions Definitions) GetMessage(num int) *elements.Message {
<<<<<<< HEAD
	return &definitions.Message[num]
=======
	return &definitions.Msg[num]
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}

// GetSignal ...
func (definitions Definitions) GetSignal(num int) *elements.Signal {
	return &definitions.Signal[num]
}

/** BPMNDI **/

// SetDiagram ...
func (definitions Definitions) GetDiagram(num int) canvas.DIAGRAM_PTR {
	return &definitions.Diagram[num]
}

/*
 * Default String
 */

// String ...
func (definitions Definitions) String() string {
	return fmt.Sprintf("id=%v", definitions.ID)
}

// String ...
func (definitions TDefinitions) String() string {
	return fmt.Sprintf("id=%v", definitions.ID)
}

/* Schema Fetcher */
func (definitions Definitions) GetCamundaSchema() {
	schema, err := utils.FetchSchema(definitions.CamundaSchema)
	if err != nil {
		log.Print(err)
	}
	log.Print(string(schema))
}
