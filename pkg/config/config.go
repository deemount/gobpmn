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
