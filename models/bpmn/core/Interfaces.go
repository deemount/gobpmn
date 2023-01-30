package core

import (
	"github.com/deemount/gobpmn/models/bpmn/canvas"
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/marker"
	"github.com/deemount/gobpmn/models/bpmn/pool"
	"github.com/deemount/gobpmn/models/bpmn/process"
)

type DefinitionsElements interface {
	SetCollaboration()
	GetCollaboration() *pool.Collaboration
	SetProcess(num int)
	GetProcess(num int) *process.Process
	SetCategory(num int)
	GetCategory(num int) *marker.Category
	SetMessage(num int)
	GetMessage(num int) *elements.Message
	GetSignal(num int) *elements.Signal
	SetSignal(num int)
	GetDiagram(num int) *canvas.Diagram
	SetDiagram(num int)
}

// DefinitionsRepository ...
type DefinitionsRepository interface {
	impl.IFBaseID
	DefinitionsElements
	SetBpmn()
	SetBpmnDI()
	SetOmgDI()
	SetDC()
	SetOmgDC()
	SetBioc()
	SetXSD()
	SetXSI()
	SetXsiSchemaLocation()
	SetXsiSchemaLocationHTTPS()

	SetTargetNamespace()
	SetCamundaSchema()
	SetZeebeSchema()
	SetModelerSchema()
	SetModelerExecutionPlatform()
	SetModelerExecutionPlatformVersion(version string)
	SetExporter()
	SetExporterVersion(version string)

	SetDefaultAttributes()
}
