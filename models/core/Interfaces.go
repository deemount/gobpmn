package core

import (
	"github.com/deemount/gobpmn/models/canvas"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/marker"
	"github.com/deemount/gobpmn/models/pool"
	"github.com/deemount/gobpmn/models/process"
)

type DefinitionsElements interface {
	SetCollaboration()
	GetCollaboration() *pool.Collaboration
	SetProcess(num int)
	GetProcess(num int) *process.Process
	SetCategory(num int)
	GetCategory(num int) *marker.Category
	SetMessage(num int)
	GetMessage(num int) *marker.Message
	GetSignal(num int) *marker.Signal
	SetSignal(num int)
	GetDiagram(num int) *canvas.Diagram
	SetDiagram(num int)
}

// DefinitionsRepository ...
type DefinitionsRepository interface {
	compulsion.IFBaseID
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
