package elements

import (
	"testing"
)

func TestSetID(t *testing.T) {
	type args struct {
		typ    string
		suffix interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetID(tt.args.typ, tt.args.suffix); got != tt.want {
				t.Errorf("SetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetIntermediateCatchEvent(t *testing.T) {
	type args struct {
		e    *IntermediateCatchEvent
		name string
		hash []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
<<<<<<< HEAD

=======
			SetIntermediateCatchEvent(tt.args.e, tt.args.name, tt.args.hash...)
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
		})
	}
}
