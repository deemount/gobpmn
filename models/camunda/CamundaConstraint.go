package camunda

// CamundaConstraintRepository ...
type CamundaConstraintRepository interface {
	CamundaBaseName
	SetConfig(config string)
	GetConfig() *string
}

// CamundaConstraint ...
type CamundaConstraint struct {
	Name   string `xml:"name,attr,omitempty" json:"name"`
	Config string `xml:"config,attr,omitempty" json:"config,omitempty"`
}

// NewCamundaConstraint ...
func NewCamunadConstraint() CamundaConstraintRepository {
	return &CamundaConstraint{}
}

/*
 * Default Getters
 */

/* Attributes */

// SetName ...
func (cconstraint *CamundaConstraint) SetName(name string) {
	cconstraint.Name = name
}

// SetConfig ...
func (cconstraint *CamundaConstraint) SetConfig(config string) {
	cconstraint.Config = config
}

/*
 * Default Getters
 */

/* Attributes */

// GetName ...
func (cconstraint CamundaConstraint) GetName() STR_PTR {
	return &cconstraint.Name
}

// GetConfig ...
func (cconstraint CamundaConstraint) GetConfig() *string {
	return &cconstraint.Config
}
