package repository

import (
	"context"
	"crypto/md5"
	"encoding/xml"
	"hash/adler32"
	"io/ioutil"
	"os"
	"time"

	"github.com/bwmarrin/snowflake"

	"github.com/deemount/gobpmn/models"
	"github.com/deemount/gobpmn/spec/process_instance"
	"github.com/deemount/gobpmn/utils"
)

// ProcessInfo ...
type ProcessInfo struct {
	BpmnProcessId string              // The ID as defined in the BPMN file
	Version       int32               // A version of the process, default=1, incremented, when another process with the same ID is loaded
	ProcessKey    int64               // The engines key for this given process with version
	ModelType     string              // The type of the model. Can be human-driven or system-driven
	Def           models.TDefinitions // parsed file content
	checksumBytes [16]byte            // internal checksum to identify different versions
}

// ProcessInstanceInfo ...
type ProcessInstanceInfo struct {
	processInfo     *ProcessInfo
	instanceKey     int64
	variableContext map[string]interface{}
	createdAt       time.Time
	state           process_instance.State
}

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

// CreateInstance creates a new instance for a process with given processKey
func (engine *BpmnEngine) CreateInstance(processKey int64, variableContext map[string]interface{}) (*ProcessInstanceInfo, error) {
	if variableContext == nil {
		variableContext = map[string]interface{}{}
	}
	for _, process := range engine.Processes {
		if process.ProcessKey == processKey {
			processInstanceInfo := ProcessInstanceInfo{
				processInfo:     &process,
				instanceKey:     engine.generateKey(),
				variableContext: variableContext,
				createdAt:       time.Now(),
				state:           process_instance.READY,
			}
			engine.ProcessInstances = append(engine.ProcessInstances, &processInstanceInfo)
			return &processInstanceInfo, nil
		}
	}
	return nil, nil
}

// fileExist ... function to check if file exists
func (engine *BpmnEngine) fileExist(fileName string) bool {
	_, error := os.Stat("files/" + fileName + ".bpmn")
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
	xmlData, err := ioutil.ReadFile("files/" + filename + ".bpmn")
	if err != nil {
		return nil, err
	}
	return engine.LoadFromBytes(xmlData)
}

// LoadFromBytes loads a given BPMN file by xmlData byte array into the engine
// and returns ProcessInfo details for the deployed workflow
func (engine *BpmnEngine) LoadFromBytes(xmlData []byte) (*ProcessInfo, error) {

	// create checksum
	md5sum := md5.Sum(xmlData)

	// assign definitions
	var definitions models.TDefinitions

	// unmarshal xml data to parsed definitions
	err := xml.Unmarshal(xmlData, &definitions)
	if err != nil {
		return nil, err
	}

	// fill process informations
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

// GetProcessInfo ...
func (engine *BpmnEngine) GetProcessInfo(ctx context.Context, b BpmnFactory) (*ProcessInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	done := make(chan func() (*ProcessInfo, error))
	go engine.requestFile(done, b)
	processInfo, err := (<-done)()
	return processInfo, err
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
