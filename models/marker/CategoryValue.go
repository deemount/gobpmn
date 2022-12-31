package marker

import (
	"fmt"

	"github.com/deemount/gobpmn/models/compulsion"
)

// NewCategoryValue ...
func NewCategoryValue() CategoryValueRepository {
	return &CategoryValue{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID
func (categoryValue *CategoryValue) SetID(typ string, suffix interface{}) {
	switch typ {
	case "":
		categoryValue.ID = fmt.Sprintf("_%s", suffix)
		break
	case "id":
		categoryValue.ID = fmt.Sprintf("_%s", suffix)
		break
	}
}

// SetValue
func (categoryValue *CategoryValue) SetValue(value string) {
	categoryValue.Value = value
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID
func (categoryValue CategoryValue) GetID() compulsion.STR_PTR {
	return &categoryValue.ID
}

// GetValue
func (categoryValue CategoryValue) GetValue() *string {
	return &categoryValue.Value
}
