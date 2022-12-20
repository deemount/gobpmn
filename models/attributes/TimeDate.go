package attributes

import "fmt"

// TimeDateRepository ...
type TimeDateRepository interface {
	SetTimerDefinitionType()
	SetTimerDefinition(timerDefinition string)

	GetTimerDefinitionType() *string
	GetTimerDefinition() *string
}

// TimeDate ...
type TimeDate struct {
	TimerDefType string `xml:"xsi:type,attr,omitempty" json:"timerDefType,omitempty"`
	TimerDef     string `xml:",innerxml" json:"timerDef,omitempty"`
}

func NewTimeDate() TimeDateRepository {
	return &TimeDate{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetTimerDefinitionType ...
func (timedate *TimeDate) SetTimerDefinitionType() {
	timedate.TimerDefType = fmt.Sprint("bpmn:tFormalExpression")
}

// SetTimerDefinition ...
func (timedate *TimeDate) SetTimerDefinition(timerDefinition string) {
	timedate.TimerDef = timerDefinition
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetTimerDefinitionType ...
func (timedate TimeDate) GetTimerDefinitionType() *string {
	return &timedate.TimerDefType
}

// GetTimerDefinition ...
func (timedate TimeDate) GetTimerDefinition() *string {
	return &timedate.TimerDef
}
