package core

var (
	repository = "DefinitionsRepository"
	fieldLong  = "definitions"
	fieldShort = "def"
)

// SetMainElements ...
func SetMainElements(d DefinitionsRepository, num int) {
	d.SetMainElements(num)
}
