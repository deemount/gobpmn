package attributes

import "fmt"

// PropertyRepository ...
type PropertyRepository interface {
	SetID(typ string, suffix interface{})
	GetID() *string
	SetName(name string)
	GetName() *string
}

// Property ...
type Property struct {
	ID   string `xml:"id,attr,omitempty" json:"id"`
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

// NewProperty ...
func NewProperty() PropertyRepository {
	return &Property{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetID ...
func (prop *Property) SetID(typ string, suffix interface{}) {
	switch typ {
	case "property":
		prop.ID = fmt.Sprintf("Property_%v", suffix)
		break
	case "id":
		prop.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
// Notice: maybe put a placeholder like name="__targetRef_placeholder" in it
func (prop *Property) SetName(name string) {
	prop.Name = name
}

/*
 * Default Getters
 */

/* Attributes */

// GetID ...
func (prop Property) GetID() *string {
	return &prop.ID

}

// GetName ...
func (prop Property) GetName() *string {
	return &prop.Name
}
