package parser

import (
	"fmt"
	"strings"
)

func ToPlainXML(bytes []byte) []byte {
	s := strings.Replace(string(bytes), "bpmn2:", "", -1)
	s = strings.ReplaceAll(s, "xmlns:bpmn=\"http://www.omg.org/spec/BPMN/20100524/MODEL\"", "")
	s = strings.ReplaceAll(s, "xmlns:dc=\"http://www.omg.org/spec/DD/20100524/DC\"", "")
	s = strings.ReplaceAll(s, "targetNamespace=\"http://bpmn.io/schema/bpmn\"", "")
	s = strings.Replace(s, "bpmn:", "", -1)
	s = strings.Replace(s, "bpmndi:", "", -1)
	s = strings.Replace(s, "camunda:", "", -1)
	s = strings.Replace(s, "di:", "", -1)
	s = strings.Replace(s, "dc:", "", -1)
	s = strings.Replace(s, " >", ">", -1)
	s = strings.Replace(s, " />", "/>", -1)
	s = strings.Replace(s, "   ", " ", -1)
	return fmt.Appendf(nil, "%v", string(s))
}
