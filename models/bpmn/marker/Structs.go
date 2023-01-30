package marker

import (
	"github.com/deemount/gobpmn/models/bpmn/attributes"
)

// marker
type DelegateParameter struct {
}

/*
 * Elementary
 */

// Notice: can be implemented in seperate package

type IncomingOutgoing struct {
	Incoming []Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing []Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

type TIncomingOutgoing struct {
	Incoming []Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing []Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

// Category ...
type Category struct {
	ID            string          `xml:"id,attr,omitempty" json:"id"`
	CategoryValue []CategoryValue `xml:"bpmn:categoryValue,omitempty" json:"categoryValue,omitempty"`
	attributes.Attributes
}

// TCategory ...
type TCategory struct {
	ID            string          `xml:"id,attr,omitempty" json:"id"`
	CategoryValue []CategoryValue `xml:"categoryValue,omitempty" json:"categoryValue,omitempty"`
	attributes.TAttributes
}

// CategoryValue ...
type CategoryValue struct {
	ID    string `xml:"id,attr,omitempty" json:"id"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty"`
}

// Group ...
type Group struct {
	ID               string `xml:"id,attr,omitempty" json:"id"`
	CategoryValueRef string `xml:"categoryValueRef,attr,omitempty" json:"categoryValueRef,omitempty"`
	attributes.Attributes
}

// TGroup ...
type TGroup struct {
	ID               string `xml:"id,attr,omitempty" json:"id"`
	CategoryValueRef string `xml:"categoryValueRef,attr,omitempty" json:"categoryValueRef,omitempty"`
	attributes.TAttributes
}

// Incoming ...
type Incoming struct {
	Flow string `xml:",innerxml" json:"flow"`
}

// Outgoing ...
type Outgoing struct {
	Flow string `xml:",innerxml" json:"flow"`
}
