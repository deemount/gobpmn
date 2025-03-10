package common

// initializeIndices initializes indices for different types of elements
func initializeIndices() map[processElement]int {
	return map[processElement]int{
		startEvent:             0,
		endEvent:               0,
		intermediateCatchEvent: 0,
		intermediateThrowEvent: 0,
		task:                   0,
		userTask:               0,
		scriptTask:             0,
		sequenceFlow:           0,
		inclusiveGateway:       0,
		exclusiveGateway:       0,
		parallelGateway:        0,
	}
}
