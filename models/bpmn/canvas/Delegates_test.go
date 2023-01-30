package canvas

import (
	"testing"
)

func TestSetShape(t *testing.T) {
	type args struct {
		p DelegateParameter
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetShape(tt.args.p)
		})
	}
}

func TestSetEdge(t *testing.T) {
	type args struct {
		p DelegateParameter
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetEdge(tt.args.p)
		})
	}
}

func TestSetLabel(t *testing.T) {
	type args struct {
		p DelegateParameter
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLabel(tt.args.p)
		})
	}
}

func TestSetPool(t *testing.T) {
	type args struct {
		p DelegateParameter
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetPool(tt.args.p)
		})
	}
}
