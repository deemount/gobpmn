// gobpmn.go
package gobpmn

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/deemount/gobpmn/internal/parser"
	"github.com/deemount/gobpmn/pkg/core"
	"github.com/deemount/gobpmn/pkg/core/common"
	"github.com/deemount/gobpmn/pkg/core/foundation"
	"github.com/deemount/gobpmn/pkg/types"
)

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
// This is used to create a BPMN model from a static struct definition.
type HasDefinitions interface {
	GetDefinitions() Repository
}

// Options defines global optional parameters for model creation
type Options struct {
	Timeout  time.Duration
	Validate bool
}

// DefaultOptions provides sane defaults (e.g., 60s timeout)
func DefaultOptions() Options {
	return Options{
		Timeout:  time.Second * 60,
		Validate: true,
	}
}

// FromStruct creates a BPMN model from a static struct definition
func FromStruct[T HasDefinitions](model T, opts ...Options) (parser.BPMNParserRepository, error) {

	start := time.Now()

	var opt Options
	if len(opts) > 0 {
		opt = opts[0]
	} else {
		opt = DefaultOptions()
	}

	ctx, cancel := context.WithTimeout(context.Background(), opt.Timeout)
	defer cancel()

	core, err := core.Core(ctx, model, []reflect.StructField{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize core: %w", err)
	}

	def := core.GetDefinitions()

	bpmn, err := parser.NewBPMNParser(
		parser.WithCounter(),
		parser.WithDefinitions(def),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create parser: %w", err)
	}

	// marshal bpmn
	if err := bpmn.Marshal(); err != nil {
		return nil, fmt.Errorf("failed to marshal BPMN model: %w", err)
	}

	// validate bpmn
	if opt.Validate {
		if err := bpmn.Validate(); err != nil {
			return nil, fmt.Errorf("BPMN model validation failed: %w", err)
		}
	}

	log.Println("total time:", time.Since(start))

	return bpmn, nil
}

// FromMap creates a BPMN model from a generic or typed map definition
func FromMap[T any](model T, opts ...Options) (parser.BPMNParserRepository, error) {

	start := time.Now()

	var opt Options
	if len(opts) > 0 {
		opt = opts[0]
	} else {
		opt = DefaultOptions()
	}

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

	bpmn, err := parser.NewBPMNParser(
		parser.WithCounter(),
		parser.WithDefinitions(def),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create parser: %w", err)
	}

	// marshal bpmn
	if err := bpmn.Marshal(); err != nil {
		return nil, fmt.Errorf("failed to marshal BPMN model: %w", err)
	}

	// validate bpmn
	if opt.Validate {
		if err := bpmn.Validate(); err != nil {
			return nil, fmt.Errorf("bpmn model validation failed: %w", err)
		}
	}

	log.Println("total time:", time.Since(start))

	return bpmn, nil
}
