//go:build go1.24
// +build go1.24

package main

import (
	"log"

	"github.com/deemount/gobpmn/pkg/transformer"
)

func main() {

	/*
		processor := transformer.NewXSLTProcessor("saxon")
		err := processor.Transform(
			"assets/document_structure/2/main.xml",
			"assets/xslt/BPMN20-ToXMI.xslt",
			"assets/xmi/ds2.xmi")
		if err != nil {
			log.Fatal(err)
		}
	*/

	processor := transformer.NewXSLTProcessor("saxon")
	err := processor.Transform(
		"blueprint/collaboration_process.bpmn",
		"assets/xslt/BPMN20-ToXMI_v2.xslt",
		"assets/xmi/output/collaboration_process.xmi")
	if err != nil {
		log.Fatal(err)
	}

}
