package config

import (
	"reflect"
)

// NOTE: New Version (commented out)
/*
type ModelConfig struct {
	Name     string
	Type     reflect.Type
	Fields   []types.OrderedField
	Wrap     reflect.Value
	Instance reflect.Value
	Def      reflect.Value
	Pool     reflect.Value
}
*/

// ModelConfig holds the configuration for model generation
type ModelConfig[T any] struct {
	Name     string
	Type     reflect.Type
	Fields   T
	Wrap     reflect.Value
	Instance reflect.Value
	Def      reflect.Value
	Pool     reflect.Value
}

// ProcessConfig holds process-specific configuration
type ProcessConfig struct {
	Process     []reflect.Value
	ProcessName []string
	ProcessType []string
	ProcessHash []string
	ProcessExec []bool
}

// CollaborationConfig holds collaboration-specific configuration
type CollaborationConfig struct {
	Type string
	Hash string
}

// ParticipantConfig holds participant-specific configuration
type ParticipantConfig struct {
	ParticipantName []string
	ParticipantHash []string
}

// ParticipantDetails contains participant-specific information
type ParticipantDetails struct {
	ID          string
	Hash        string
	Name        string
	ProcessRef  string
	ProcessHash string
}
