package camunda

// CamundaField ...
type CamundaFieldRepository interface {
	CamundaBaseName
	SetCamundaExpression()
	GetCamundaExpression() *CamundaExpression
	SetCamundaString()
	GetCamundaString() *CamundaString
}

// CamundaField ...
type CamundaField struct {
	Name              string              `xml:"name,attr,omitempty" json:"name"`
	CamundaExpression []CamundaExpression `xml:"camunda:expression,omitempty" json:"expression,omitempty"`
	CamundaString     []CamundaString     `xml:"camunda:string,omitempty" json:"string,omitempty"`
}

// TCamundaField ...
type TCamundaField struct {
	Name       string              `xml:"name,attr,omitempty" json:"name"`
	Expression []CamundaExpression `xml:"expression,omitempty" json:"expression,omitempty"`
	String     []CamundaString     `xml:"string,omitempty" json:"string,omitempty"`
}

// NewCamundaField ...
func NewCamundaField() CamundaFieldRepository {
	return &CamundaField{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetName ...
func (cf *CamundaField) SetName(name string) {
	cf.Name = name
}

/*** Make Elements ***/

/** Camunda **/

// SetCamundaExpression ...
func (cf *CamundaField) SetCamundaExpression() {
	cf.CamundaExpression = make([]CamundaExpression, 1)
}

// SetCamundaString ...
func (cf *CamundaField) SetCamundaString() {
	cf.CamundaString = make([]CamundaString, 1)
}

/*
 * Default Getters
 */

/* Attributes */

// GetName ...
func (cf CamundaField) GetName() STR_PTR {
	return &cf.Name
}

/* Elements */

/** Camunda **/

// GetCamundaExpression ...
func (cf CamundaField) GetCamundaExpression() *CamundaExpression {
	return &cf.CamundaExpression[0]
}

// GetCamundaString ...
func (cf CamundaField) GetCamundaString() *CamundaString {
	return &cf.CamundaString[0]
}
