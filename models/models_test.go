package models

import (
	"testing"
)

func TestNewModels(t *testing.T) {

	def := NewModels()
	t.Logf("NewModels()=%+v", def)

}
