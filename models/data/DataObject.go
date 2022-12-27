package data

import "fmt"

// DataObjectRepository ...
type DataObjectRepository interface {
	DataBaseID
}

// DataObject ...
type DataObject struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
}

// NewDataObject ...
func NewDataObject() DataObjectRepository {
	return &DataObject{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetID ...
func (do *DataObject) SetID(typ string, suffix interface{}) {
	switch typ {
	case "dataobject":
		do.ID = fmt.Sprintf("DataObject_%v", suffix)
		break
	case "id":
		do.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

/*
 * Default Getters
 */

/* Attributes */

// GetID ...
func (do DataObject) GetID() STR_PTR {
	return &do.ID
}
