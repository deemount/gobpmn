package models

// CamundaConstraint ...
type CamundaConstraint struct {
	Name   string `xml:"name,attr,omitempty" json:"name"`
	Config string `xml:"config,attr,omitempty" json:"config,omitempty"`
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
