package transformer

import (
	"fmt"
	"os/exec"
)

// XSLTProcessor handles XSLT transformations
type XSLTProcessor struct {
	XsltPath string
	Saxon    string // Path to Saxon processor
}

// Initialize XSLT processor with Saxon path
func NewXSLTProcessor(saxonPath string) *XSLTProcessor {
	return &XSLTProcessor{
		Saxon: saxonPath,
	}
}

// Transform applies XSLT transformation
// saxon: https://www.saxonica.com/documentation12/index.html#!using-xsl/commandline
func (p *XSLTProcessor) Transform(inputXML, xsltFile, outputFile string) error {
	cmd := exec.Command(p.Saxon,
		"-s:"+inputXML,     // Source XML file
		"-xsl:"+xsltFile,   // XSLT file
		"-o:"+outputFile,   // Output file
		"-explain:out.txt", // Explain the transformation
		"-t",               // Trace the transformation, useful for debugging
		"-warnings:silent", // Suppress warnings
		"-ext:on",          // Enable extension functions
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("XSLT transformation failed: %v\nOutput: %s", err, output)
	}
	return nil
}
