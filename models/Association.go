package models

// Association ...
type Association struct {
	ID                string              `xml:"id,attr"`
	Text              []Text              `xml:"bpmn:text,omitempty"`
	Documentation     []Documentation     `xml:"bpmn:documentation,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
}

/* Attributes */

/** BPMN **/

/* Elements */

/** BPMN **/
