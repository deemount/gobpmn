package factory

// Factory ...
type Factory interface {
}

// factory ...
type factory struct {
}

// NewFactory ...
func NewFactory() Factory {
	return &factory{}
}
