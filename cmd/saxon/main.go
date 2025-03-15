//go:build go1.24
// +build go1.24

package main

import (
	"log"
)

func main() {
	processor := NewXSLTProcessor("saxon")
	err := processor.Transform(
		"assets/xmi_2.5.1/input.xmi",
		"assets/xmi_2.5.1/transform.xslt",
		"assets/xmi_2.5.1/output.xml")
	if err != nil {
		log.Fatal(err)
	}
}
