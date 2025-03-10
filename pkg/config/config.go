package config

import "reflect"

// ModelConfig holds the configuration for model generation
type ModelConfig struct {
	Name     string
	Fields   []reflect.StructField
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

// ParticipantConfig holds participant-specific configuration
type ParticipantConfig struct {
	ParticipantName []string
	ParticipantHash []string
}
