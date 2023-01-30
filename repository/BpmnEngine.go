package repository

import (
	"crypto/md5"
	"encoding/xml"
	"hash/adler32"
	"io/ioutil"
	"log"
	"os"

	"github.com/bwmarrin/snowflake"

	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/utils"
)

// BpmnEngine ...
type BpmnEngine struct {
	Name             string
	Processes        []ProcessInfo
	ProcessInstances []*ProcessInstanceInfo
	snowflake        *snowflake.Node
}

// NewEngine creates an engine with an arbitrary name of the engine;
// useful in case you have multiple ones, in order to distinguish them.
func NewBpmnEngine(name string) BpmnEngine {
	snowflakeIdGenerator := initializeSnowflakeIdGenerator()
	return BpmnEngine{
		Name:             name,
		Processes:        []ProcessInfo{},
		ProcessInstances: []*ProcessInstanceInfo{},
		snowflake:        snowflakeIdGenerator,
	}
}

// fileExist ... function to check if file exists
func (engine *BpmnEngine) fileExist(fileName string) bool {
	_, error := os.Stat("files/bpmn/" + fileName + ".bpmn")
	// check if error is "file not exists"
	if os.IsNotExist(error) {
		return false
	} else {
		return true
	}
}

// LoadFromFile loads a given BPMN file by filename into the engine
// and returns ProcessInfo details for the deployed workflow
func (engine *BpmnEngine) LoadFromFile(filename string) (*ProcessInfo, error) {
	var err error
	log.Printf("engine: load bpmn file from %s", "files/bpmn/"+filename+".bpmn")
	xmlData, err := ioutil.ReadFile("files/bpmn/" + filename + ".bpmn")
	if err != nil {
		return nil, err
	}
	return engine.LoadFromBytes(xmlData)
}

// LoadFromBytes loads a given BPMN file by xmlData byte array into the engine
// and returns a pointer to ProcessInfo, which stores details for the deployed workflow
func (engine *BpmnEngine) LoadFromBytes(xmlData []byte) (*ProcessInfo, error) {

	// create checksum
	md5sum := md5.Sum(xmlData)

	// assign tdefinitions
	// Notice:
	// If you use the models.Definitions, the unmarshal panics with
	// panic: expected element type <bpmn:definitions> but have <definitions>
	var definitions core.TDefinitions

	// unmarshal xml data to parsed definitions
	log.Println("engine: unmarshal file")
	err := xml.Unmarshal(xmlData, &definitions)
	if err != nil {
		return nil, err
	}

	// fill process informations into struct
	processInfo := ProcessInfo{
		Version: 1,
		Def:     definitions,
	}

	// set model type
	if len(definitions.Process) == 1 {
		processInfo.ModelType = "system-driven"
	} else {
		processInfo.ModelType = "human-driven"
	}

	for _, procInfo := range engine.Processes {
		for _, proc := range definitions.Process {
			if procInfo.BpmnProcessId == proc.ID {
				if utils.AreEqual(procInfo.checksumBytes, md5sum) {
					return &procInfo, nil
				} else {
					processInfo.Version = procInfo.Version + 1
				}
			}

		}
	}

	for _, proc := range definitions.Process {
		processInfo.BpmnProcessId = proc.ID
	}

	processInfo.ProcessKey = engine.generateKey()
	processInfo.checksumBytes = md5sum
	engine.Processes = append(engine.Processes, processInfo)

	return &processInfo, nil
}

// requestFile ...
func (engine *BpmnEngine) requestFile(done chan func() (*ProcessInfo, error), b BpmnFactory) {
	log.Println("engine: request created file")
	done <- (func() (*ProcessInfo, error) {
		filename := b.GetCurrentlyCreatedFile()
		if engine.fileExist(filename) {
			process, err := engine.LoadFromFile(filename)
			if err != nil {
				panic(err)
			}
			return process, err
		}
		return nil, nil
	})
}

// generateKey ...
func (engine *BpmnEngine) generateKey() int64 {
	return engine.snowflake.Generate().Int64()
}

// initializeSnowflakeIdGenerator ...
func initializeSnowflakeIdGenerator() *snowflake.Node {
	hash32 := adler32.New()
	for _, e := range os.Environ() {
		hash32.Sum([]byte(e))
	}
	snowflakeNode, err := snowflake.NewNode(int64(hash32.Sum32()))
	if err != nil {
		panic("Can't initialize snowflake ID generator. Message: " + err.Error())
	}
	return snowflakeNode
}
