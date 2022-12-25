package examples

/**************************************************************************************/
/**
 * @Import
 *
 * Import models and utils package to build the model
 **/

import (
	"github.com/deemount/gobpmn/models/canvas"
	"github.com/deemount/gobpmn/models/core"
	"github.com/deemount/gobpmn/models/marker"
	"github.com/deemount/gobpmn/models/pool"
	"github.com/deemount/gobpmn/utils"
)

/**************************************************************************************/
/**
 * @BASE CLASS
 *
 *
 **/

type BlackBoxModel interface {
	Create() blackBoxModel
}

/**************************************************************************************/
/**
 * @OBJECTS
 * - blackBoxPool
 * - blackBoxMessage
 **/

type blackBoxPool struct {
	CollaborationID     string
	CustomerID          string
	CustomerName        string
	CustomerSupportID   string
	CustomerSupportName string
	ManufacturerID      string
	ManufacturerName    string
}

type blackBoxMessage struct {
	OrderHash             string
	RequestSparePartsHash string
	ReplacementSupplyHash string
	ConfirmationHash      string
	ShipmentHash          string
}

type blackBoxModel struct {
	def core.DefinitionsRepository
	blackBoxPool
	blackBoxMessage
}

/**************************************************************************************/
/**
 * @BUILDER
 *
 *
 **/

func NewBlackBoxModel() BlackBoxModel {
	return &blackBoxModel{
		def: new(core.Definitions),
		blackBoxPool: blackBoxPool{
			CollaborationID:     utils.GenerateHash(),
			CustomerID:          utils.GenerateHash(),
			CustomerName:        "Customer Name",
			CustomerSupportID:   utils.GenerateHash(),
			CustomerSupportName: "Customer Support",
			ManufacturerID:      utils.GenerateHash(),
			ManufacturerName:    "Manufacturer",
		},
		blackBoxMessage: blackBoxMessage{
			OrderHash:             utils.GenerateHash(),
			RequestSparePartsHash: utils.GenerateHash(),
			ReplacementSupplyHash: utils.GenerateHash(),
			ConfirmationHash:      utils.GenerateHash(),
			ShipmentHash:          utils.GenerateHash(),
		},
	}
}

/**************************************************************************************/
/**
 * @Create
 * @@setMainElements
 * @@setInnerElements
 * @@setDefinitionsAttributes
 * @@setCollaboration
 * @@setDiagram
 * @Def
 **/

func (bb blackBoxModel) Create() blackBoxModel {
	bb.setMainElements()
	bb.setInnerElements()
	bb.setDefinitionsAttributes()
	bb.setCollaboration()
	bb.setDiagram()
	return bb
}

// Def ...
func (bb *blackBoxModel) Def() *core.DefinitionsRepository {
	return &bb.def
}

/**************************************************************************************/

/**
 * @Setters I
 * @setMainElements
 * @@models.Definitions: SetCollaboration
 * @@models.Definitions: SetProcess
 * @@models.Definitions: SetDiagram
 *
 * @setInnerElements
 * @@Model: GetCollaboration
 * @@Model: GetDiagram
 *
 * @setDefinitionsAttributes
 * @@models.Definitions: SetDefaultAttributes
 *
 * @setCollaboration
 * @@Model: GetParticipant
 * @@Model: GetCollaboration
 **/

func (bb *blackBoxModel) setMainElements() {
	bb.def.SetCollaboration()
	bb.def.SetDiagram(1)
}

func (bb *blackBoxModel) setInnerElements() {

	collaboration := bb.GetCollaboration()
	collaboration.SetParticipant(3)
	collaboration.SetMessageFlow(5)

	diagram := bb.GetDiagram()
	diagram.SetPlane()
	plane := bb.GetPlane()

	plane.SetShape(3)
	plane.SetEdge(5)
}

func (bb *blackBoxModel) setDefinitionsAttributes() {
	bb.def.SetDefaultAttributes()
}

func (bb *blackBoxModel) setCollaboration() {

	customerParticipant := bb.GetParticipantCustomer(bb.GetCollaboration())
	customerSupportParticipant := bb.GetParticipantCustomerSupport(bb.GetCollaboration())
	manufacturerParticipant := bb.GetParticipantManufacturer(bb.GetCollaboration())

	order := bb.GetMessageOrder()
	requestSpareParts := bb.GetMessageRequestSpareParts()
	replacementSupply := bb.GetMessageReplacementSupply()
	confirmation := bb.GetMessageConfirmation()
	shipment := bb.GetMessageShipment()

	bb.GetCollaboration().SetID("collaboration", bb.CollaborationID)

	customerParticipant.SetID("participant", bb.CustomerID)
	customerParticipant.SetName(bb.CustomerName)

	customerSupportParticipant.SetID("participant", bb.CustomerSupportID)
	customerSupportParticipant.SetName(bb.CustomerSupportName)

	manufacturerParticipant.SetID("participant", bb.ManufacturerID)
	manufacturerParticipant.SetName(bb.ManufacturerName)

	order.SetID("flow", bb.OrderHash)
	order.SetName("Order")
	order.SetSourceRef("participant", bb.CustomerID)
	order.SetTargetRef("participant", bb.CustomerSupportID)

	requestSpareParts.SetID("flow", bb.RequestSparePartsHash)
	requestSpareParts.SetName("Request Spare Parts")
	requestSpareParts.SetSourceRef("participant", bb.CustomerSupportID)
	requestSpareParts.SetTargetRef("participant", bb.ManufacturerID)

	replacementSupply.SetID("flow", bb.ReplacementSupplyHash)
	replacementSupply.SetName("Replacement Supply")
	replacementSupply.SetSourceRef("participant", bb.ManufacturerID)
	replacementSupply.SetTargetRef("participant", bb.CustomerSupportID)

	confirmation.SetID("flow", bb.ConfirmationHash)
	confirmation.SetName("Confirmation")
	confirmation.SetSourceRef("participant", bb.CustomerSupportID)
	confirmation.SetTargetRef("participant", bb.CustomerID)

	shipment.SetID("flow", bb.ShipmentHash)
	shipment.SetName("Shipment")
	shipment.SetSourceRef("participant", bb.CustomerSupportID)
	shipment.SetTargetRef("participant", bb.CustomerID)

	bb.setPoolCustomer()
	bb.setPoolCustomerSupport()
	bb.setPoolManufacturer()

	edgeMsgOrder := bb.GetEdgeMessageOrder(bb.GetPlane())
	edgeMsgOrder.SetID("flow", bb.OrderHash)
	edgeMsgOrder.SetElement("flow", bb.OrderHash)
	edgeMsgOrder.SetWaypoint()
	edgeMsgOrder.Waypoint[0].SetCoordinates(225, 140)
	edgeMsgOrder.Waypoint[1].SetCoordinates(225, 220)
	edgeMsgOrder.SetLabel()
	edgeMsgOrder.Label[0].SetBounds()
	edgeMsgOrder.Label[0].Bounds[0].SetCoordinates(240, 163)
	edgeMsgOrder.Label[0].Bounds[0].SetSize(29, 14)

	edgeMsgRequestSpareParts := bb.GetEdgeMessageRequestSpareParts(bb.GetPlane())
	edgeMsgRequestSpareParts.SetID("flow", bb.RequestSparePartsHash)
	edgeMsgRequestSpareParts.SetElement("flow", bb.RequestSparePartsHash)
	edgeMsgRequestSpareParts.SetWaypoint()
	edgeMsgRequestSpareParts.Waypoint[0].SetCoordinates(225, 280)
	edgeMsgRequestSpareParts.Waypoint[1].SetCoordinates(225, 370)
	edgeMsgRequestSpareParts.SetLabel()
	edgeMsgRequestSpareParts.Label[0].SetBounds()
	edgeMsgRequestSpareParts.Label[0].Bounds[0].SetCoordinates(240, 306)
	edgeMsgRequestSpareParts.Label[0].Bounds[0].SetSize(74, 27)

	edgeMsgReplacementSupply := bb.GetEdgeMessageReplacementSupply(bb.GetPlane())
	edgeMsgReplacementSupply.SetID("flow", bb.ReplacementSupplyHash)
	edgeMsgReplacementSupply.SetElement("flow", bb.ReplacementSupplyHash)
	edgeMsgReplacementSupply.SetWaypoint()
	edgeMsgReplacementSupply.Waypoint[0].SetCoordinates(388, 370)
	edgeMsgReplacementSupply.Waypoint[1].SetCoordinates(388, 280)
	edgeMsgReplacementSupply.SetLabel()
	edgeMsgReplacementSupply.Label[0].SetBounds()
	edgeMsgReplacementSupply.Label[0].Bounds[0].SetCoordinates(408, 306)
	edgeMsgReplacementSupply.Label[0].Bounds[0].SetSize(65, 27)

	edgeMsgConfirmation := bb.GetEdgeMessageConfirmation(bb.GetPlane())
	edgeMsgConfirmation.SetID("flow", bb.ConfirmationHash)
	edgeMsgConfirmation.SetElement("flow", bb.ConfirmationHash)
	edgeMsgConfirmation.SetWaypoint()
	edgeMsgConfirmation.Waypoint[0].SetCoordinates(385, 219)
	edgeMsgConfirmation.Waypoint[1].SetCoordinates(385, 140)
	edgeMsgConfirmation.SetLabel()
	edgeMsgConfirmation.Label[0].SetBounds()
	edgeMsgConfirmation.Label[0].Bounds[0].SetCoordinates(408, 163)
	edgeMsgConfirmation.Label[0].Bounds[0].SetSize(63, 14)

	edgeMsgShipment := bb.GetEdgeMessageShipment(bb.GetPlane())
	edgeMsgShipment.SetID("flow", bb.ShipmentHash)
	edgeMsgShipment.SetElement("flow", bb.ShipmentHash)
	edgeMsgShipment.SetWaypoint()
	edgeMsgShipment.Waypoint[0].SetCoordinates(545, 219)
	edgeMsgShipment.Waypoint[1].SetCoordinates(545, 140)
	edgeMsgShipment.SetLabel()
	edgeMsgShipment.Label[0].SetBounds()
	edgeMsgShipment.Label[0].Bounds[0].SetCoordinates(566, 163)
	edgeMsgShipment.Label[0].Bounds[0].SetSize(47, 14)
}

/**
 * @Setters II
 * @Process
 * @private setPoolCustomer
 * @private setPoolCustomerSupport
 * @private setPoolManufacturer
 * @private setDiagram
**/

func (bb *blackBoxModel) setPoolCustomer() {
	e := bb.GetShapePoolCustomer(bb.GetPlane())
	e.SetID("participant", bb.CustomerID)
	e.SetElement("participant", bb.CustomerID)
	e.SetIsHorizontal(true)
	e.SetBounds()
	e.Bounds[0].SetCoordinates(150, 80)
	e.Bounds[0].SetSize(650, 60)
}

func (bb *blackBoxModel) setPoolCustomerSupport() {
	e := bb.GetShapePoolCustomerSupport(bb.GetPlane())
	e.SetID("participant", bb.CustomerSupportID)
	e.SetElement("participant", bb.CustomerSupportID)
	e.SetIsHorizontal(true)
	e.SetBounds()
	e.Bounds[0].SetCoordinates(150, 220)
	e.Bounds[0].SetSize(650, 60)
}

func (bb *blackBoxModel) setPoolManufacturer() {
	e := bb.GetShapePoolManufacturer(bb.GetPlane())
	e.SetID("participant", bb.ManufacturerID)
	e.SetElement("participant", bb.ManufacturerID)
	e.SetIsHorizontal(true)
	e.SetBounds()
	e.Bounds[0].SetCoordinates(150, 370)
	e.Bounds[0].SetSize(650, 60)
}

func (bb *blackBoxModel) setDiagram() {
	var n int64 = 1
	diagram := bb.GetDiagram()
	diagram.SetID("diagram", n)
	p := bb.GetPlane()
	p.SetID("plane", n)
	p.SetElement("collaboration", bb.CollaborationID)
}

/**************************************************************************************/

/**
 * @Getters I
 * @GetCollaboration -> models.Collaboration
 * @GetParticipantCustomer -> models.Participant
 * @GetParticipantCustomerSupport -> models.Participant
 * @GetParticipantManufacturer -> models.Participant
 * @GetMessageOrder -> models.MessageFlow
 * @GetMessageRequestSpareParts -> models.MessageFlow
 * @GetMessageReplacementSupply -> models.MessageFlow
 * @GetMessageConfirmation -> models.MessageFlow
 * @GetMessageShipment -> models.MessageFlow
 * @GetDiagram -> canvas.Diagram
**/

func (bb blackBoxModel) GetCollaboration() *pool.Collaboration {
	return bb.def.GetCollaboration()
}

func (bb blackBoxModel) GetParticipantCustomer(e *pool.Collaboration) *pool.Participant {
	return e.GetParticipant(0)
}

func (bb blackBoxModel) GetParticipantCustomerSupport(e *pool.Collaboration) *pool.Participant {
	return e.GetParticipant(1)
}

func (bb blackBoxModel) GetParticipantManufacturer(e *pool.Collaboration) *pool.Participant {
	return e.GetParticipant(2)
}

func (bb blackBoxModel) GetMessageOrder() *marker.MessageFlow {
	return bb.GetCollaboration().GetMessageFlow(0)
}

func (bb blackBoxModel) GetMessageRequestSpareParts() *marker.MessageFlow {
	return bb.GetCollaboration().GetMessageFlow(1)
}

func (bb blackBoxModel) GetMessageReplacementSupply() *marker.MessageFlow {
	return bb.GetCollaboration().GetMessageFlow(2)
}

func (bb blackBoxModel) GetMessageConfirmation() *marker.MessageFlow {
	return bb.GetCollaboration().GetMessageFlow(3)
}

func (bb blackBoxModel) GetMessageShipment() *marker.MessageFlow {
	return bb.GetCollaboration().GetMessageFlow(4)
}

func (bb blackBoxModel) GetDiagram() *canvas.Diagram {
	return bb.def.GetDiagram(0)
}

func (bb blackBoxModel) GetPlane() *canvas.Plane {
	return bb.GetDiagram().GetPlane()
}

func (bb blackBoxModel) GetShapePoolCustomer(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(0)
}

func (bb blackBoxModel) GetShapePoolCustomerSupport(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(1)
}

func (bb blackBoxModel) GetShapePoolManufacturer(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(2)
}

func (bb blackBoxModel) GetEdgeMessageOrder(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(0)
}

func (bb blackBoxModel) GetEdgeMessageRequestSpareParts(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(1)
}

func (bb blackBoxModel) GetEdgeMessageReplacementSupply(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(2)
}

func (bb blackBoxModel) GetEdgeMessageConfirmation(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(3)
}

func (bb blackBoxModel) GetEdgeMessageShipment(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(4)
}
