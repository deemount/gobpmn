package parser

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/deemount/gobpmn/pkg/core/foundation"
)

var (
	// Defaults
	DefaultCounter        = 1            // DefaultCounter is the default counter
	DefaultFilenamePrefix = "diagram"    // DefaultFilenamePrefix is the default filename prefix
	DefaultPathBPMN       = "files/bpmn" // DefaultPathBPMN is the default path to the bpmn files
	DefaultPathJSON       = "files/json" // DefaultPathJSON is the default path to the json files

	// Errors
	ErrPathNotFound      = errors.New("path not found")           // ErrPathNotFound is the error for the path not found
	ErrEmptyFilePathBPMN = errors.New("empty file path for bpmn") // ErrEmptyFilePathBPMN is the error for the empty file path for bpmn
)

type (

	// BPMNParserRepository ...
	BPMNParserRepository interface {
		Marshal() error
		Validate() error
	}

	// BPMNParser ...
	BPMNParser struct {
		Counter        int                              // Counter is the number of created files
		FilenamePrefix string                           // FilenamePrefix is a part of the name of the current file
		FilePathBPMN   string                           // FilePathBPMN is the path to the bpmn files
		FilePathJSON   string                           // FilePathJSON is the path to the json files
		Def            *foundation.Definitions          // Def is the pointer definition of the model
		Repo           foundation.DefinitionsRepository // Repo is the repository of the model
	}

	// Option is a functional option type for the Parser
	Option = func(parser *BPMNParser) error
)

// NewBPMNParser initializes the parser with the given options
// and returns the bpmn parser repository.
// The method sets the default values and applies the options
// to the parser.
func NewBPMNParser(opts ...Option) (BPMNParserRepository, error) {
	// Set the default values
	parser := &BPMNParser{
		Counter:        DefaultCounter,
		FilenamePrefix: DefaultFilenamePrefix,
		FilePathBPMN:   DefaultPathBPMN,
		FilePathJSON:   DefaultPathJSON,
		Def:            nil,
		Repo:           nil,
	}
	// Apply the options to the parser
	for _, opt := range opts {
		opt(parser)
	}
	if err := parser.validate(); err != nil {
		return nil, err
	}
	return parser, nil
}

// Marshal builds the bpmn and json files and returns the BPMNParser
// and an error if an error occurs during the process.
// The method calls the MarshalBPMN and MarshalJSON methods.
func (parser BPMNParser) Marshal() error {
	if err := parser.save(); err != nil {
		return err
	}
	_, err := parser.marshalJSON()
	if err != nil {
		return err
	}
	return nil
}

// Save ...
func (parser *BPMNParser) save() error {

	if parser.FilePathBPMN == "" {
		return ErrEmptyFilePathBPMN
	}
	bpmnData, err := xml.MarshalIndent(&parser.Repo, " ", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling BPMN data: %v", err)
	}

	return parser.writeFile(parser.FilePathBPMN, parser.GetFilename(), bpmnData, "bpmn")

}

// writeFile handles the file writing process and error handling.
func (parser *BPMNParser) writeFile(path, filename string, data []byte, extension string) error {
	fullPath := fmt.Sprintf("%s/%s.%s", path, filename, extension)

	f, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", fullPath, err)
	}
	defer f.Close()

	if extension == "json" {
		// Write the JSON data to the file
		if _, err := f.Write(data); err != nil {
			return fmt.Errorf("error writing to file %s: %v", fullPath, err)
		}
		return f.Sync()
	}

	if _, err := f.Write([]byte(xml.Header + string(data))); err != nil {
		return fmt.Errorf("error writing to file %s: %v", fullPath, err)
	}

	return f.Sync()
}

// MarshalJSON marshals the repository data to JSON and saves it as a file.
func (parser *BPMNParser) marshalJSON() ([]byte, error) {
	jsonData, err := json.MarshalIndent(parser.Repo, " ", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON data: %v", err)
	}

	err = parser.writeFile(parser.FilePathJSON, parser.GetFilename(), jsonData, "json")
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

// GetFilename returns the current bpmn filename
// It is a helper method to get the current bpmn filename, which
// relies on the same instance
func (parser BPMNParser) GetFilename() string {
	return fmt.Sprintf("%s_%d", parser.FilenamePrefix, parser.Counter)
}

// validate validates the file oath for the BPMN file
// and returns an error if the path is empty.
func (parser *BPMNParser) validate() error {
	if parser.FilePathBPMN == "" {
		return ErrEmptyFilePathBPMN
	}
	return nil
}

// WithPath validates and sets paths for BPMN and JSON files
func WithPath(paths ...string) Option {
	return func(parser *BPMNParser) error {
		if len(paths) == 0 || len(paths) > 2 {
			return fmt.Errorf("invalid number of paths provided; expected 1 or 2, got %d", len(paths))
		}

		for _, path := range paths {
			if !isValidPath(path) {
				return fmt.Errorf("invalid path provided: %s", path)
			}
		}

		parser.FilePathBPMN = paths[0]
		if len(paths) == 2 {
			parser.FilePathJSON = paths[1]
		}
		return nil
	}
}

// Helper function to check if a given path exists and is accessible
func isValidPath(path string) bool {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return false
	}
	_, err = os.Stat(absPath)
	return !os.IsNotExist(err)
}

// WithCounter counts .bpmn files in the provided path and updates the Counter
func WithCounter() Option {
	return func(parser *BPMNParser) error {
		if parser.FilePathBPMN == "" {
			return ErrEmptyFilePathBPMN
		}

		count, err := countBPMNFiles(parser.FilePathBPMN)
		if err != nil {
			return fmt.Errorf("error counting BPMN files: %v", err)
		}

		parser.Counter = count + 1 // Next file will use this counter
		return nil
	}
}

// Helper function to count .bpmn files in a given directory
func countBPMNFiles(dir string) (int, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".bpmn" {
			count++
		}
	}
	return count, nil
}

// WithFilenamePrefix sets the filename prefix for the parser
func WithFilenamePrefix(filenamePrefix string) Option {
	return func(parser *BPMNParser) error {
		if !isValidFilenamePrefix(filenamePrefix) {
			return fmt.Errorf("invalid filename prefix: %s", filenamePrefix)
		}
		parser.FilenamePrefix = filenamePrefix
		return nil
	}
}

// Helper function to validate the filename prefix.
func isValidFilenamePrefix(prefix string) bool {
	var validPrefix = regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
	return validPrefix.MatchString(prefix)
}

// WithDefinitions assigns a DefinitionsRepository to the parser.
func WithDefinitions(repo foundation.DefinitionsRepository) Option {
	return func(parser *BPMNParser) error {
		if repo == nil {
			return errors.New("invalid DefinitionsRepository: repository is nil")
		}
		parser.Repo = repo
		return nil
	}
}
