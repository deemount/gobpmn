package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

	// BPMNModellerRepository ...
	BPMNModellerRepository interface {
		Marshal() error
	}

	// BPMNModeller ...
	BPMNModeller struct {
		Counter        int                   // Counter is the number of created files
		FilenamePrefix string                // FilenamePrefix is a part of the name of the current file
		FilePathBPMN   string                // FilePathBPMN is the path to the bpmn files
		FilePathJSON   string                // FilePathJSON is the path to the json files
		Def            *Definitions          // Def is the pointer definition of the model
		Repo           DefinitionsRepository // Repo is the repository of the model
	}

	// Option is a functional option type for the Modeller
	Option = func(modeller *BPMNModeller) error
)

// NewBPMNModeller initializes the modeller with the given options
// and returns the bpmn modeller repository.
// The method sets the default values and applies the options
// to the modeller.
func NewBPMNModeller(opts ...Option) (BPMNModellerRepository, error) {
	// Set the default values
	modeller := &BPMNModeller{
		Counter:        DefaultCounter,
		FilenamePrefix: DefaultFilenamePrefix,
		FilePathBPMN:   DefaultPathBPMN,
		FilePathJSON:   DefaultPathJSON,
		Def:            nil,
		Repo:           nil,
	}
	// Apply the options to the modeller
	for _, opt := range opts {
		opt(modeller)
	}
	if err := modeller.validate(); err != nil {
		return nil, err
	}
	return modeller, nil
}

// Marshal builds the bpmn and json files and returns the BPMNModeller
// and an error if an error occurs during the process.
// The method calls the MarshalBPMN and MarshalJSON methods.
func (modeller BPMNModeller) Marshal() error {
	if err := modeller.save(); err != nil {
		return err
	}
	_, err := modeller.marshalJSON()
	if err != nil {
		return err
	}
	return nil
}

// Save ...
func (modeller *BPMNModeller) save() error {

	if modeller.FilePathBPMN == "" {
		return ErrEmptyFilePathBPMN
	}
	bpmnData, err := xml.MarshalIndent(&modeller.Repo, " ", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling BPMN data: %v", err)
	}

	return modeller.writeFile(modeller.FilePathBPMN, modeller.getFilename(), bpmnData, "bpmn")

}

// writeFile handles the file writing process and error handling.
func (modeller *BPMNModeller) writeFile(path, filename string, data []byte, extension string) error {
	fullPath := fmt.Sprintf("%s/%s.%s", path, filename, extension)

	f, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", fullPath, err)
	}
	defer f.Close()

	// Validate the BPMN data against the XSD schema
	/*err = clib.ValidateBPMNAgainstXSD(data, "xsd/BPMN20.xsd")
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("Validation successful.")
	}*/

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
// BUG: It marshals also the xml header.
func (modeller *BPMNModeller) marshalJSON() ([]byte, error) {
	jsonData, err := json.MarshalIndent(modeller.Repo, " ", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON data: %v", err)
	}

	err = modeller.writeFile(modeller.FilePathJSON, modeller.getFilename(), jsonData, "json")
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

// GetFilename returns the current bpmn filename
// It is a helper method to get the current bpmn filename, which
// relies on the same instance
func (modeller BPMNModeller) getFilename() string {
	return fmt.Sprintf("%s_%d", modeller.FilenamePrefix, modeller.Counter)
}

// validate validates the modeller fields
// and returns an error if the validation fails.
func (modeller *BPMNModeller) validate() error {
	if modeller.FilePathBPMN == "" {
		return ErrEmptyFilePathBPMN
	}
	return nil
}

// WithPath validates and sets paths for BPMN and JSON files
func WithPath(paths ...string) Option {
	return func(modeller *BPMNModeller) error {
		if len(paths) == 0 || len(paths) > 2 {
			return fmt.Errorf("invalid number of paths provided; expected 1 or 2, got %d", len(paths))
		}

		for _, path := range paths {
			if !isValidPath(path) {
				return fmt.Errorf("invalid path provided: %s", path)
			}
		}

		modeller.FilePathBPMN = paths[0]
		if len(paths) == 2 {
			modeller.FilePathJSON = paths[1]
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
	return func(modeller *BPMNModeller) error {
		if modeller.FilePathBPMN == "" {
			return ErrEmptyFilePathBPMN
		}

		count, err := countBPMNFiles(modeller.FilePathBPMN)
		if err != nil {
			return fmt.Errorf("error counting BPMN files: %v", err)
		}

		modeller.Counter = count + 1 // Next file will use this counter
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

// WithFilenamePrefix sets the filename prefix for the modeller
func WithFilenamePrefix(filenamePrefix string) Option {
	return func(modeller *BPMNModeller) error {
		if !isValidFilenamePrefix(filenamePrefix) {
			return fmt.Errorf("invalid filename prefix: %s", filenamePrefix)
		}
		modeller.FilenamePrefix = filenamePrefix
		return nil
	}
}

// Helper function to validate the filename prefix.
func isValidFilenamePrefix(prefix string) bool {
	var validPrefix = regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
	return validPrefix.MatchString(prefix)
}

// WithDefinitions assigns a DefinitionsRepository to the modeller.
func WithDefinitions(repo DefinitionsRepository) Option {
	return func(modeller *BPMNModeller) error {
		if repo == nil {
			return errors.New("invalid DefinitionsRepository: repository is nil")
		}
		modeller.Repo = repo
		return nil
	}
}
