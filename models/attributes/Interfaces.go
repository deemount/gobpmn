package attributes

import (
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/compulsion"
)

type AttributesBaseElements interface {
	SetDocumentation()
	GetDocumentation() *Documentation
	SetExtensionElements()
	GetExtensionElements() *ExtensionElements
}

// DocumentationRepository ...
type DocumentationRepository interface {
	SetData(data string)
	GetData() *string
}

// ExtensionElementsRepository ...
type ExtensionElementsRepository interface {
	camunda.ExtensionElementsRepository
}

// PropertyRepository ...
type PropertyRepository interface {
	compulsion.IFBaseID
	compulsion.IFBaseName
}
