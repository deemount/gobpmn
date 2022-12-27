package camunda

// CamundaPropertiesRepository ...
type CamundaPropertiesRepository interface {
	SetCamundaProperty(num int)
	GetCamundaProperty(num int) *CamundaProperty
}

// CamundaProperties ...
type CamundaProperties struct {
	CamundaProperty []CamundaProperty `xml:"camunda:properties,omitempty" json:"properties,omitempty"`
}

// TCamundaProperties ...
type TCamundaProperties struct {
	Property []CamundaProperty `xml:"properties,omitempty" json:"properties,omitempty"`
}

// NewCamundaProperties ...
func NewCamundaProperties() CamundaPropertiesRepository {
	return &CamundaProperties{}
}

/*
 * Default Setters
 */

/*** Make Elements ***/

/** Camunda **/

// SetCamundaProperties ...
func (properties *CamundaProperties) SetCamundaProperty(num int) {
	properties.CamundaProperty = make([]CamundaProperty, num)
}

/*
 * Default Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaProperties ...
func (properties CamundaProperties) GetCamundaProperty(num int) *CamundaProperty {
	return &properties.CamundaProperty[num]
}
