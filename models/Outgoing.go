package models

// Outgoing ...
type Outgoing struct {
	Flow string `xml:",innerxml"`
}

// SetFlow ...
func (og *Outgoing) SetFlow(suffix string) {
	og.Flow = "Flow_" + suffix
}
