package factory

// Options
type Options struct {
	Counter     int
	CurrentFile string
	ModelType   string
}

type FactoryConfiguration func(o Options) Options
