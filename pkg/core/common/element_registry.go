package common

import (
	"regexp"
	"strings"
)

const MaxPriority = 999 // Standard value for high priority

// Match is a struct that holds
// - the name of the element,
// - the element itself,
// - a boolean to check if the match is exact,
// - a priority value and
// - a regex
type Match struct {
	name     string
	element  processElement
	exact    bool
	priority int
	regex    *regexp.Regexp
}

var ElementTypeList = []Match{
	{
		name:     "StartEvent",
		element:  startEvent,
		exact:    true, // NOTE: maybe useless; scale the fuunctionality
		priority: 1,    // NOTE: maybe useless; scale the fuunctionality
		regex:    regexp.MustCompile(`^StartEvent$`),
	},
	{
		name:     "EndEvent",
		element:  endEvent,
		exact:    true, // NOTE: maybe useless; scale the fuunctionality
		priority: 1,    // NOTE: maybe useless; scale the fuunctionality
		regex:    regexp.MustCompile(`^EndEvent$`),
	},
	{
		name:     "IntermediateCatchEvent",
		element:  intermediateCatchEvent,
		exact:    true, // NOTE: maybe useless; scale the fuunctionality
		priority: 1,    // NOTE: maybe useless; scale the fuunctionality
		regex:    regexp.MustCompile(`^IntermediateCatchEvent$`),
	},
	{
		name:     "IntermediateThrowEvent",
		element:  intermediateThrowEvent,
		exact:    true, // NOTE: maybe useless; scale the fuunctionality
		priority: 1,    // NOTE: maybe useless; scale the fuunctionality
		regex:    regexp.MustCompile(`^IntermediateThrowEvent$`),
	},
	{
		name:     "ParallelGateway",
		element:  parallelGateway,
		exact:    true, // NOTE: maybe useless; scale the fuunctionality
		priority: 1,    // NOTE: maybe useless; scale the fuunctionality
		regex:    regexp.MustCompile(`^ParallelGateway$`),
	},
	{
		name:     "ExclusiveGateway",
		element:  exclusiveGateway,
		exact:    true, // NOTE: maybe useless; scale the fuunctionality
		priority: 1,    // NOTE: maybe useless; scale the fuunctionality
		regex:    regexp.MustCompile(`^ExclusiveGateway$`),
	},
	{
		name:     "InclusiveGateway",
		element:  inclusiveGateway,
		exact:    true, // NOTE: maybe useless; scale the fuunctionality
		priority: 1,    // NOTE: maybe useless; scale the fuunctionality
		regex:    regexp.MustCompile(`^InclusiveGateway$`),
	},
	{
		name:     "UserTask",
		element:  userTask,
		exact:    false, // NOTE: maybe useless; scale the fuunctionality
		priority: 2,     // NOTE: maybe useless; scale the fuunctionality
		regex:    regexp.MustCompile(`^UserTask$`),
	},
	{
		name:     "ScriptTask",
		element:  scriptTask,
		exact:    false, // NOTE: maybe useless scale the fuunctionality
		priority: 3,     // NOTE: maybe useless; scale the fuunctionality
		regex:    regexp.MustCompile(`^ScriptTask$`),
	},
	{
		name:     "Task",
		element:  task,
		exact:    false, // NOTE: maybe useless; scale the fuunctionality
		priority: 4,     // NOTE: maybe useless; scale the fuunctionality
		regex:    regexp.MustCompile(`^Task$`),
	},
	{
		name:     "From",
		element:  sequenceFlow,
		exact:    false, // NOTE: maybe useless; scale the fuunctionality
		priority: 1,     // NOTE: maybe useless; scale the fuunctionality
		regex:    regexp.MustCompile(`^From.*$`),
	},
}

// matchElementType returns the processElement type of the given name.
// It checks if the name is empty and trims the name.
// It checks if the name is an exact match and returns the element.
// If there is no exact match, it returns the best match.
// If there is no best match, it returns an empty processElement.
// NOTE:
// maybe the bestMatch option is useless,
// also the first pass with the exact match is may not be needed.
func matchElementType(name string) processElement {

	if name == "" {
		return processElement("")
	}

	normalizedName := strings.TrimSpace(name)

	for _, match := range ElementTypeList {
		if match.exact && normalizedName == match.name {
			return match.element
		}
	}

	var bestMatch Match
	bestMatch.priority = MaxPriority

	for _, match := range ElementTypeList {
		if !match.exact &&
			strings.Contains(strings.ToLower(normalizedName), strings.ToLower(match.name)) &&
			match.priority < bestMatch.priority {
			bestMatch = match
		}
	}

	// NOTE: maybe the bestMatch option is useless, also the first pass with the exact match is may not be needed.
	for _, entry := range ElementTypeList {
		if entry.regex.MatchString(bestMatch.name) {
			return entry.element
		}
	}

	// NOTE: maybe useless, since  regexp.MatchString seems to be the method to use
	if bestMatch.priority != MaxPriority {
		return bestMatch.element
	}

	return processElement("")
}
