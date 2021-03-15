package models

import (
	"fmt"
	"strconv"
)

// Shape ...
type Shape struct {
	ID      string `xml:"id,attr"`
	Element string `xml:"bpmnElement,attr"`
	Stroke  string `xml:"bioc:stroke,attr,omitempty"`
	Fill    string `xml:"bioc:fill,attr,omitempty"`
	Bounds  Bounds `xml:"dc:Bounds"`
}

// SetStartEventID ...
func (sh *Shape) SetStartEventID(num int64) {
	sh.ID = "_BPMNShape_StartEvent_" + strconv.FormatInt(num, 16)
}

// SetStartEventElement ...
func (sh *Shape) SetStartEventElement(num int64) {
	sh.Element = "StartEvent_" + strconv.FormatInt(num, 16)
}

// SetID ...
func (sh *Shape) SetID(suffix string) {
	sh.ID = "Event_" + suffix + "_di"
}

// SetElement ...
func (sh *Shape) SetElement(suffix string) {
	sh.Element = "Event_" + suffix
}

/*
	Color Management
	stroke & fill for start event
	groups:
	own: use rgb colors
	default: black/white
	green, red, blue, orange, purple
*/

// SetStroke ...
func (sh *Shape) SetStroke(r, g, b int8) {
	sh.Stroke = fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}

// SetFill ...
func (sh *Shape) SetFill(r, g, b int8) {
	sh.Fill = fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}

// SetStrokeWhite ...
func (sh *Shape) SetStrokeBlack() {
	sh.Stroke = "black"
}

// SetFillGreen ...
func (sh *Shape) SetFillWhite() {
	sh.Fill = "white"
}

// SetStrokeGreen ...
func (sh *Shape) SetStrokeGreen() {
	sh.Stroke = "rgb(67, 160, 71)"
}

// SetFillGreen ...
func (sh *Shape) SetFillGreen() {
	sh.Fill = "rgb(200, 230, 201)"
}

// SetStrokeRed ...
func (sh *Shape) SetStrokeRed() {
	sh.Stroke = "rgb(229, 57, 53)"
}

// SetFillRed ...
func (sh *Shape) SetFillRed() {
	sh.Fill = "rgb(255, 205, 210)"
}

// SetStrokeOrange ...
func (sh *Shape) SetStrokeOrange() {
	sh.Stroke = "rgb(251, 140, 0)"
}

// SetFillOrange ...
func (sh *Shape) SetFillOrange() {
	sh.Fill = "rgb(255, 224, 178)"
}

// SetStrokeBlue ...
func (sh *Shape) SetStrokeBlue() {
	sh.Stroke = "rgb(30, 136, 229)"
}

// SetFillBlue ...
func (sh *Shape) SetFillBlue() {
	sh.Fill = "rgb(187, 222, 251)"
}

// SetStrokePurple ...
func (sh *Shape) SetStrokePurple() {
	sh.Stroke = "rgb(142, 36, 170)"
}

// SetFillPurple ...
func (sh *Shape) SetFillPurple() {
	sh.Fill = "rgb(225, 190, 231)"
}
