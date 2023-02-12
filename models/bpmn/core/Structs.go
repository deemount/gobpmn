package core

import (
	"encoding/xml"

	"github.com/deemount/gobpmn/models/bpmn/canvas"
<<<<<<< HEAD
	"github.com/deemount/gobpmn/models/bpmn/collaboration"
	"github.com/deemount/gobpmn/models/bpmn/events"
	"github.com/deemount/gobpmn/models/bpmn/marker"
=======
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/marker"
	"github.com/deemount/gobpmn/models/bpmn/pool"
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
	"github.com/deemount/gobpmn/models/bpmn/process"
)

// DefinitionsBaseElements ...
type DefinitionsBaseElements struct {
<<<<<<< HEAD
	Collaboration []collaboration.Collaboration `xml:"bpmn:collaboration,omitempty" json:"collaboration"`
	Process       []process.Process             `xml:"bpmn:process,omitempty" json:"process"`
	Category      []marker.Category             `xml:"bpmn:category,omitempty" json:"category,omitempty"`
	events.CoreEvents
	Diagram []canvas.Diagram `xml:"bpmndi:BPMNDiagram,omitempty" json:"diagram"`
=======
	Collaboration []pool.Collaboration `xml:"bpmn:collaboration,omitempty" json:"collaboration"`
	Process       []process.Process    `xml:"bpmn:process,omitempty" json:"process"`
	Category      []marker.Category    `xml:"bpmn:category,omitempty" json:"category,omitempty"`
	Msg           []elements.Message   `xml:"bpmn:message,omitempty" json:"message,omitempty"`
	Signal        []elements.Signal    `xml:"bpmn:signal,omitempty" json:"signal,omitempty"`
	Diagram       []canvas.Diagram     `xml:"bpmndi:BPMNDiagram,omitempty" json:"diagram"`
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}

// TDefinitionsBaseElements ...
type TDefinitionsBaseElements struct {
<<<<<<< HEAD
	Collaboration []collaboration.TCollaboration `xml:"collaboration,omitempty" json:"collaboration"`
	Process       []process.TProcess             `xml:"process,omitempty" json:"process"`
	Category      []marker.TCategory             `xml:"category,omitempty" json:"category,omitempty"`
	events.TCoreEvents
	Diagram []canvas.TDiagram `xml:"BPMNDiagram,omitempty" json:"diagram"`
=======
	Collaboration []pool.TCollaboration `xml:"collaboration,omitempty" json:"collaboration"`
	Process       []process.TProcess    `xml:"process,omitempty" json:"process"`
	Category      []marker.TCategory    `xml:"category,omitempty" json:"category,omitempty"`
	Msg           []elements.TMessage   `xml:"message,omitempty" json:"message,omitempty"`
	Signal        []elements.TSignal    `xml:"signal,omitempty" json:"signal,omitempty"`
	Diagram       []canvas.TDiagram     `xml:"BPMNDiagram,omitempty" json:"diagram"`
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}

// Definitions represents the root element
type Definitions struct {
	XMLName                         xml.Name `xml:"bpmn:definitions" json:"-"`
	Bpmn                            string   `xml:"xmlns:bpmn,attr" json:"-"`
	Xsd                             string   `xml:"xmlns:xsd,attr,omitempty" json:"-"`
	Xsi                             string   `xml:"xmlns:xsi,omitempty" json:"-"`
	XsiSchemaLocation               string   `xml:"xsi:schemaLocation,attr,omitempty" json:"-"`
	BpmnDI                          string   `xml:"xmlns:bpmndi,attr" json:"-"`
	OmgDI                           string   `xml:"xmlns:omgdi,attr,omitempty" json:"-"`
	DC                              string   `xml:"xmlns:dc,attr,omitempty" json:"-"`
	OmgDC                           string   `xml:"xmlns:omgdc,attr,omitempty" json:"-"`
	Bioc                            string   `xml:"xmlns:bioc,attr,omitempty" json:"-"`
	CamundaSchema                   string   `xml:"xmlns:camunda,attr,omitempty" json:"-"`
	Zeebe                           string   `xml:"xmlns:zeebe,omitempty" json:"-"`
	Modeler                         string   `xml:"xmlns:modeler,omitempty" json:"-"`
	ModelerExecutionPlatform        string   `xml:"modeler:executionPlatform,omitempty" json:"-"`
	ModelerExecutionPlatformVersion string   `xml:"modeler:executionPlatformVersion,omitempty" json:"-"`
	ID                              string   `xml:"id,attr" json:"id"`
	TargetNamespace                 string   `xml:"targetNamespace,attr" json:"-"`
	Exporter                        string   `xml:"exporter,attr,omitempty" json:"-"`
	ExporterVersion                 string   `xml:"exporterVersion,attr,omitempty" json:"-"`
	DefinitionsBaseElements
}

// TDefinitions ...
type TDefinitions struct {
	XMLName xml.Name `xml:"definitions" json:"-"`
	ID      string   `xml:"id,attr" json:"id"`
	TDefinitionsBaseElements
}
