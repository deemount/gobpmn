package repository

import (
	"os"

	"github.com/gocarina/gocsv"
)

// create .csv
/*err = bpmnFactory.toCSV()
if err != nil {
	return err
}*/

// toCSV ...
func (factory *bpmnFactory) toCSV() error {
	var err error

	f, err := os.Create("files/csv/" + factory.Options.CurrentFile + ".csv")
	if err != nil {
		return err
	}
	defer f.Close()

	err = gocsv.MarshalFile(&factory.Repo, f)
	if err != nil {
		panic(err)
	}

	return nil

}
