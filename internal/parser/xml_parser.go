package parser

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Element struct {
	Name      string
	Attrs     map[string]string
	Children  []*Element
	Text      string
	Namespace string
}

func XMLParser(filePath string) (*Element, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %v", filePath, err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	var root *Element
	var stack []*Element

	for {
		tok, err := decoder.Token()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("error decoding XML: %v", err)
		}

		switch elem := tok.(type) {
		case xml.StartElement:
			node := &Element{
				Name:      elem.Name.Local,
				Namespace: elem.Name.Space,
				Attrs:     make(map[string]string),
			}
			for _, attr := range elem.Attr {
				node.Attrs[attr.Name.Local] = attr.Value
			}

			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, node)
			} else {
				root = node
			}

			stack = append(stack, node)

		case xml.CharData:
			if len(stack) > 0 {
				text := strings.TrimSpace(string(elem))
				if len(text) > 0 {
					stack[len(stack)-1].Text = text
				}
			}

		case xml.EndElement:
			stack = stack[:len(stack)-1]
		}
	}

	if root == nil {
		return nil, errors.New("error: no valid XML structure found")
	}

	return root, nil
}

func FindElementsByPath(root *Element, path string) ([]*Element, error) {
	if root == nil {
		return nil, errors.New("root element is nil")
	}

	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 0 {
		return nil, errors.New("invalid XPath")
	}
	return findElementsRecursive(root, parts, 0), nil
}

func findElementsRecursive(node *Element, path []string, depth int) []*Element {
	if node == nil || depth >= len(path) {
		return []*Element{node}
	}

	var result []*Element
	if node.Name == path[depth] {
		if depth == len(path)-1 {
			result = append(result, node)
		} else {
			for _, child := range node.Children {
				result = append(result, findElementsRecursive(child, path, depth+1)...)
			}
		}
	}
	return result
}

func ToPlainXML(bytes []byte) []byte {

	s := strings.ReplaceAll(string(bytes), "xmlns:bpmn=\"http://www.omg.org/spec/BPMN/20100524/MODEL\"", "")
	s = strings.ReplaceAll(s, "xmlns:bizagi=\"http://www.bizagi.com/schemas/2011/BizagiProcessModel\"", "")
	s = strings.ReplaceAll(s, "xmlns:semantic=\"http://www.omg.org/spec/BPMN/20100524/Semantic\"", "")
	s = strings.ReplaceAll(s, "xmlns:xmlns=\"http://www.w3.org/2001/XMLSchema-instance\"", "")
	s = strings.ReplaceAll(s, "xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"", "")
	s = strings.ReplaceAll(s, "xmlns:dc=\"http://www.omg.org/spec/DD/20100524/DC\"", "")
	s = strings.ReplaceAll(s, "xmlns:di=\"http://www.omg.org/spec/DD/20100524/DI\"", "")
	s = strings.ReplaceAll(s, "xmlns:bpmndi=\"http://www.omg.org/spec/BPMN/20100524/DI\"", "")
	s = strings.ReplaceAll(s, "xmlns:camunda=\"http://camunda.org/schema/1.0/bpmn\"", "")
	s = strings.ReplaceAll(s, "targetNamespace=\"http://bpmn.io/schema/bpmn\"", "")
	s = strings.ReplaceAll(s, "bizagi:", "")
	s = strings.ReplaceAll(s, "bpmn:", "")
	s = strings.ReplaceAll(s, "bpmn2:", "")
	s = strings.ReplaceAll(s, "bpmndi:", "")
	s = strings.ReplaceAll(s, "camunda:", "")
	s = strings.ReplaceAll(s, "semantic:", "")
	s = strings.ReplaceAll(s, "di:", "")
	s = strings.ReplaceAll(s, "dc:", "")
	s = strings.ReplaceAll(s, " >", ">")
	s = strings.ReplaceAll(s, " />", "/>")
	s = strings.ReplaceAll(s, "   ", " ")
	return fmt.Appendf(nil, "%v", string(s))
}
