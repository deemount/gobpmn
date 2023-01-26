package core

import "testing"

func TestSetMainElements(t *testing.T) {
	type args struct {
		d            DefinitionsRepository
		numProcesses int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetMainElements(tt.args.d, tt.args.numProcesses)
		})
	}
}
