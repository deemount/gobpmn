//go:build go1.23
// +build go1.23

package modeler

import (
	_ "net/http/pprof"

	"github.com/deemount/gobpmn/internal/parser"
)

var bpmnParser parser.BPMNParser
