package camunda

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
