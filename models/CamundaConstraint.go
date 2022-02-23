package models

// CamundaConstraint ...
type CamundaConstraint struct {
	Name   string `xml:"name,attr,omitempty"`
	Config string `xml:"config,attr,omitempty"`
}

/* Attributes */

// SetName ...
func (cconstraint *CamundaConstraint) SetName(name string) {
	cconstraint.Name = name
}

// SetConfig ...
func (cconstraint *CamundaConstraint) SetConfig(config string) {
	cconstraint.Config = config
}
