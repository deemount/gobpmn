package search

import (
	"github.com/deemount/gobpmnModels/pkg/core"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/process"
)

// FindActivitiesBySourceRef searches for activities by their source reference
func FindActivitiesBySourceRef(definitions *core.TDefinitions, sourceRef string) []interface{} {

	activities := make([]interface{}, 0)
	activityChan := make(chan interface{})

	for _, proc := range definitions.Process {
		go func(proc process.TProcess) {
			for _, seqFlow := range proc.SequenceFlow {
				if seqFlow.SourceTargetRef.SourceRef == sourceRef {
					go func(seqFlow flow.TSequenceFlow) {
						activity := FindActivityByID(definitions, seqFlow.SourceTargetRef.TargetRef)
						if activity != nil {
							activityChan <- activity
						}
					}(seqFlow)
				}
			}
		}(proc)
	}

	go func() {
		for activity := range activityChan {
			activities = append(activities, activity)
		}
	}()

	return activities
}

// FindActivitiesByTargetRef searches for activities by their target reference
func FindActivitiesByTargetRef(definitions *core.TDefinitions, targetRef string) []interface{} {

	activities := make([]interface{}, 0)
	activityChan := make(chan interface{})

	for _, proc := range definitions.Process {
		go func(proc process.TProcess) {
			for _, seqFlow := range proc.SequenceFlow {
				if seqFlow.SourceTargetRef.TargetRef == targetRef {
					go func(seqFlow flow.TSequenceFlow) {
						activity := FindActivityByID(definitions, seqFlow.SourceTargetRef.SourceRef)
						if activity != nil {
							activityChan <- activity
						}
					}(seqFlow)
				}
			}
		}(proc)
	}

	go func() {
		for activity := range activityChan {
			activities = append(activities, activity)
		}
	}()

	return activities
}

// FindActivityByID searches for an activity by its ID in the given definitions
func FindActivityByID(definitions *core.TDefinitions, activityID string) interface{} {
	activityChan := make(chan interface{})

	for _, proc := range definitions.TDefinitionsBaseElements.Process {
		go func(proc process.TProcess) {

			for _, scriptTask := range proc.TTasks.ScriptTask {
				if scriptTask.BaseAttributes.ID == activityID {
					activityChan <- scriptTask
				}
			}

			for _, task := range proc.TTasks.Task {
				if task.BaseAttributes.ID == activityID {
					activityChan <- task
				}
			}

			for _, userTask := range proc.TTasks.UserTask {
				if userTask.BaseAttributes.ID == activityID {
					activityChan <- userTask
				}
			}

			for _, startEvent := range proc.TProcessEvents.StartEvent {
				if startEvent.BaseAttributes.ID == activityID {
					activityChan <- startEvent
				}
			}

			for _, endEvent := range proc.TProcessEvents.EndEvent {
				if endEvent.BaseAttributes.ID == activityID {
					activityChan <- endEvent
				}
			}
		}(proc)
	}

	select {
	case activity := <-activityChan:
		return activity
	default:
		return nil
	}
}
