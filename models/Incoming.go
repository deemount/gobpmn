package models

// Incoming ...
type Incoming struct {
	Flow string `xml:",innerxml"`
}

// SetFlow ...
func (ig *Incoming) SetFlow(suffix string) {
	ig.Flow = "Flow_" + suffix
}
