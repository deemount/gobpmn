package repository

import (
	"encoding/json"
	"os"
)

// create .json
/*err = bpmnFactory.toJSON()
if err != nil {
	return err
}*/

// toJSON ...
func (factory *bpmnFactory) toJSON() error {
	var err error

	// marshal json to byte slice
	b, _ := json.MarshalIndent(&factory.Repo, " ", "  ")

	// create .json file
	f, err := os.Create("files/json/" + factory.Options.CurrentFile + ".json")
	if err != nil {
		return err
	}
	defer f.Close()

	// write bytes to file
	_, err = f.Write(b)
	if err != nil {
		return err
	}

	err = f.Sync()
	if err != nil {
		return err
	}

	return nil

}
