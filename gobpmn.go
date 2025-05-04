// gobpmn.go
package gobpmn

import (
	"context"
	"fmt"
	"io"
	"os"
	"reflect"
	"time"

	"github.com/deemount/gobpmn/internal/logger"
	"github.com/deemount/gobpmn/internal/parser"
	"github.com/deemount/gobpmn/pkg/core"
	"github.com/deemount/gobpmn/pkg/core/common"
	"github.com/deemount/gobpmn/pkg/core/foundation"
	"github.com/deemount/gobpmn/pkg/types"
)

// Level is an alias to the internal log level
type Level = logger.Level

const (
	DebugLevel = logger.DebugLevel
	InfoLevel  = logger.InfoLevel
	WarnLevel  = logger.WarnLevel
	ErrorLevel = logger.ErrorLevel
)

// LoggerRepository is what already exists internally
type LoggerRepository = logger.LoggerRepository

// NewLogger builds a logger without having to import internal/logger
func NewLogger(name string, lvl Level, output io.Writer) LoggerRepository {
	return logger.NewLogger(name, lvl, output)
}

// BPMN is a re-exported type from the types package
type BPMN = types.BPMN

// Repository is a re-exported type from the foundation package
type Repository = foundation.DefinitionsRepository

// Definitions creates a new DefinitionsRepository
// with a new Definitions struct.
// This is a convenience function to create a new DefinitionsRepository
// without having to create a new Definitions struct manually.
func Definitions() Repository {
	return foundation.NewDefinitions()
}

// HasDefinitions is an interface for types that have a GetDefinitions method
// that returns a DefinitionsRepository.
// This is used to create a bpmn model from a static struct definition.
type HasDefinitions interface {
	GetDefinitions() Repository
}

// Options defines global optional parameters for model creation
type Options struct {
	Timeout  time.Duration
	Validate bool
	Logger   logger.LoggerRepository
}

// DefaultOptions provides sane defaults (e.g., 60s timeout)
func DefaultOptions() Options {
	return Options{
		Timeout:  time.Second * 60,
		Validate: true,
		Logger:   logger.NewLogger("gobpmn", logger.InfoLevel, os.Stdout),
	}
}

// FromStruct creates a BPMN model from a static struct definition
func FromStruct[T HasDefinitions](model T, opts ...Options) (parser.BPMNParserRepository, error) {

	var opt Options
	if len(opts) > 0 {
		opt = opts[0]
	} else {
		opt = DefaultOptions()
	}

	l := opt.Logger
	if l == nil {
		l = logger.NewLogger("gobpmn", logger.InfoLevel, os.Stdout)
	}

	l.Debugf("FromStruct called for model type %T with opts %+v", model, opt)

	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), opt.Timeout)
	defer cancel()
	ctx = logger.WithLogger(ctx, l)

	l.Debugf("calling core.Core")
	core, err := core.Core(ctx, model, []reflect.StructField{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize core: %w", err)
	}

	// get the definitions from the core
	// this is a pointer to the definitions struct
	// and is used to create the BPMN model
	// the definitions struct is a map of all definitions
	l.Debugf("calling core.GetDefinitions")
	def := core.GetDefinitions()

	// create a new BPMN parser
	// with the given definitions
	// and a counter to count the number of files
	l.Debugf("creating bpmn parser")
	bpmn, err := parser.NewBPMNParser(
		parser.WithCounter(),
		parser.WithDefinitions(def),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create parser: %w", err)
	}

	// marshal bpmn
	l.Debugf("marshalling bpmn model")
	if err := bpmn.Marshal(); err != nil {
		return nil, fmt.Errorf("failed to marshal bpmn model: %w", err)
	}

	// validate bpmn
	l.Debugf("validating bpmn model")
	if opt.Validate {
		if err := bpmn.Validate(); err != nil {
			return nil, fmt.Errorf("bpmn model validation failed: %w", err)
		}
	}
	l.Debugf("bpmn model created in %s", time.Since(start))

	return bpmn, nil
}

// FromMap creates a BPMN model from a generic or typed map definition
func FromMap[T any](model T, opts ...Options) (parser.BPMNParserRepository, error) {

	var opt Options
	if len(opts) > 0 {
		opt = opts[0]
	} else {
		opt = DefaultOptions()
	}

	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), opt.Timeout)
	defer cancel()

	core, err := core.Core(ctx, model, map[string]any{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize reflect DI: %w", err)
	}

	def, err := common.ConvertDynamicStructToDefinitions(ctx, core)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to definitions: %w", err)
	}

	// create a new BPMN parser
	// with the given definitions
	// and a counter to count the number of files
	bpmn, err := parser.NewBPMNParser(
		parser.WithCounter(),
		parser.WithDefinitions(def),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create parser: %w", err)
	}

	// marshal bpmn
	if err := bpmn.Marshal(); err != nil {
		return nil, fmt.Errorf("failed to marshal bpmn model: %w", err)
	}

	// validate bpmn
	if opt.Validate {
		if err := bpmn.Validate(); err != nil {
			return nil, fmt.Errorf("bpmn model validation failed: %w", err)
		}
	}

	fmt.Println("total time:", time.Since(start))

	return bpmn, nil
}
