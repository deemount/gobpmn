package core

type CoreRepository interface{ DefinitionsRepository }
type Core struct{ CoreRepository }

// NewCore ...
func NewCore() CoreRepository {
	return &Core{NewDefinitions()}
}
