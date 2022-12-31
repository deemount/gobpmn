package marker

// SetSequenceFlow ...
func SetSequenceFlow(e *SequenceFlow, typ string, sourceType string, targetType string, hash ...string) {
	e.SetID(typ, hash[0])
	e.SetSourceRef(sourceType, hash[1])
	e.SetTargetRef(targetType, hash[2])
}
