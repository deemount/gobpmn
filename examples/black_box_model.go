package examples

/**************************************************************************************/
/**
 * @Import
 *
 * Import models and utils package to build the model
 **/

import (
	"github.com/deemount/gobpmn/models"
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
	def models.DefinitionsRepository
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
		def: new(models.Definitions),
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

func (bm blackBoxModel) Create() blackBoxModel {
	bm.setMainElements()
	bm.setInnerElements()
	bm.setDefinitionsAttributes()
	bm.setCollaboration()
	bm.setDiagram()
	return bm
}

// Def ...
func (bm *blackBoxModel) Def() *models.DefinitionsRepository {
	return &bm.def
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

func (bm *blackBoxModel) setMainElements() {
	bm.def.SetCollaboration()
	bm.def.SetDiagram(1)
}

func (bm *blackBoxModel) setInnerElements() {

	collaboration := bm.GetCollaboration()
	collaboration.SetParticipant(3)
	collaboration.SetMessageFlow(5)

	bm.GetDiagram().SetPlane()
	plane := bm.GetPlane()

	plane.SetShape(3)
	plane.SetEdge(5)
}

func (bm *blackBoxModel) setDefinitionsAttributes() {
	bm.def.SetDefaultAttributes()
}

func (bm *blackBoxModel) setCollaboration() {

	customerParticipant := bm.GetParticipantCustomer(bm.GetCollaboration())
	customerSupportParticipant := bm.GetParticipantCustomerSupport(bm.GetCollaboration())
	manufacturerParticipant := bm.GetParticipantManufacturer(bm.GetCollaboration())

	order := bm.GetMessageOrder()
	requestSpareParts := bm.GetMessageRequestSpareParts()
	replacementSupply := bm.GetMessageReplacementSupply()
	confirmation := bm.GetMessageConfirmation()
	shipment := bm.GetMessageShipment()

	bm.GetCollaboration().SetID("collaboration", bm.CollaborationID)

	customerParticipant.SetID("participant", bm.CustomerID)
	customerParticipant.SetName(bm.CustomerName)

	customerSupportParticipant.SetID("participant", bm.CustomerSupportID)
	customerSupportParticipant.SetName(bm.CustomerSupportName)

	manufacturerParticipant.SetID("participant", bm.ManufacturerID)
	manufacturerParticipant.SetName(bm.ManufacturerName)

	order.SetID("flow", bm.OrderHash)
	order.SetName("Order")
	order.SetSourceRef("participant", bm.CustomerID)
	order.SetTargetRef("participant", bm.CustomerSupportID)

	requestSpareParts.SetID("flow", bm.RequestSparePartsHash)
	requestSpareParts.SetName("Request Spare Parts")
	requestSpareParts.SetSourceRef("participant", bm.CustomerSupportID)
	requestSpareParts.SetTargetRef("participant", bm.ManufacturerID)

	replacementSupply.SetID("flow", bm.ReplacementSupplyHash)
	replacementSupply.SetName("Replacement Supply")
	replacementSupply.SetSourceRef("participant", bm.ManufacturerID)
	replacementSupply.SetTargetRef("participant", bm.CustomerSupportID)

	confirmation.SetID("flow", bm.ConfirmationHash)
	confirmation.SetName("Confirmation")
	confirmation.SetSourceRef("participant", bm.CustomerSupportID)
	confirmation.SetTargetRef("participant", bm.CustomerID)

	shipment.SetID("flow", bm.ShipmentHash)
	shipment.SetName("Shipment")
	shipment.SetSourceRef("participant", bm.CustomerSupportID)
	shipment.SetTargetRef("participant", bm.CustomerID)

	bm.setPoolCustomer()
	bm.setPoolCustomerSupport()
	bm.setPoolManufacturer()

	edgeMsgOrder := bm.GetEdgeMessageOrder(bm.GetPlane())
	edgeMsgOrder.SetID("flow", bm.OrderHash)
	edgeMsgOrder.SetElement("flow", bm.OrderHash)
	edgeMsgOrder.SetWaypoint()
	edgeMsgOrder.Waypoint[0].SetCoordinates(225, 140)
	edgeMsgOrder.Waypoint[1].SetCoordinates(225, 220)
	edgeMsgOrder.SetLabel()
	edgeMsgOrder.Label[0].SetBounds()
	edgeMsgOrder.Label[0].Bounds[0].SetCoordinates(240, 163)
	edgeMsgOrder.Label[0].Bounds[0].SetSize(29, 14)

	edgeMsgRequestSpareParts := bm.GetEdgeMessageRequestSpareParts(bm.GetPlane())
	edgeMsgRequestSpareParts.SetID("flow", bm.RequestSparePartsHash)
	edgeMsgRequestSpareParts.SetElement("flow", bm.RequestSparePartsHash)
	edgeMsgRequestSpareParts.SetWaypoint()
	edgeMsgRequestSpareParts.Waypoint[0].SetCoordinates(225, 280)
	edgeMsgRequestSpareParts.Waypoint[1].SetCoordinates(225, 370)
	edgeMsgRequestSpareParts.SetLabel()
	edgeMsgRequestSpareParts.Label[0].SetBounds()
	edgeMsgRequestSpareParts.Label[0].Bounds[0].SetCoordinates(240, 306)
	edgeMsgRequestSpareParts.Label[0].Bounds[0].SetSize(74, 27)

	edgeMsgReplacementSupply := bm.GetEdgeMessageReplacementSupply(bm.GetPlane())
	edgeMsgReplacementSupply.SetID("flow", bm.ReplacementSupplyHash)
	edgeMsgReplacementSupply.SetElement("flow", bm.ReplacementSupplyHash)
	edgeMsgReplacementSupply.SetWaypoint()
	edgeMsgReplacementSupply.Waypoint[0].SetCoordinates(388, 370)
	edgeMsgReplacementSupply.Waypoint[1].SetCoordinates(388, 280)
	edgeMsgReplacementSupply.SetLabel()
	edgeMsgReplacementSupply.Label[0].SetBounds()
	edgeMsgReplacementSupply.Label[0].Bounds[0].SetCoordinates(408, 306)
	edgeMsgReplacementSupply.Label[0].Bounds[0].SetSize(65, 27)

	edgeMsgConfirmation := bm.GetEdgeMessageConfirmation(bm.GetPlane())
	edgeMsgConfirmation.SetID("flow", bm.ConfirmationHash)
	edgeMsgConfirmation.SetElement("flow", bm.ConfirmationHash)
	edgeMsgConfirmation.SetWaypoint()
	edgeMsgConfirmation.Waypoint[0].SetCoordinates(385, 219)
	edgeMsgConfirmation.Waypoint[1].SetCoordinates(385, 140)
	edgeMsgConfirmation.SetLabel()
	edgeMsgConfirmation.Label[0].SetBounds()
	edgeMsgConfirmation.Label[0].Bounds[0].SetCoordinates(408, 163)
	edgeMsgConfirmation.Label[0].Bounds[0].SetSize(63, 14)

	edgeMsgShipment := bm.GetEdgeMessageShipment(bm.GetPlane())
	edgeMsgShipment.SetID("flow", bm.ShipmentHash)
	edgeMsgShipment.SetElement("flow", bm.ShipmentHash)
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

func (bm *blackBoxModel) setPoolCustomer() {
	e := bm.GetShapePoolCustomer(bm.GetPlane())
	e.SetID("participant", bm.CustomerID)
	e.SetElement("participant", bm.CustomerID)
	e.SetIsHorizontal(true)
	e.SetBounds()
	e.Bounds[0].SetCoordinates(150, 80)
	e.Bounds[0].SetSize(650, 60)
}

func (bm *blackBoxModel) setPoolCustomerSupport() {
	e := bm.GetShapePoolCustomerSupport(bm.GetPlane())
	e.SetID("participant", bm.CustomerSupportID)
	e.SetElement("participant", bm.CustomerSupportID)
	e.SetIsHorizontal(true)
	e.SetBounds()
	e.Bounds[0].SetCoordinates(150, 220)
	e.Bounds[0].SetSize(650, 60)
}

func (bm *blackBoxModel) setPoolManufacturer() {
	e := bm.GetShapePoolManufacturer(bm.GetPlane())
	e.SetID("participant", bm.ManufacturerID)
	e.SetElement("participant", bm.ManufacturerID)
	e.SetIsHorizontal(true)
	e.SetBounds()
	e.Bounds[0].SetCoordinates(150, 370)
	e.Bounds[0].SetSize(650, 60)
}

func (bm *blackBoxModel) setDiagram() {
	var n int64 = 1
	bm.GetDiagram().SetID(n)
	p := bm.GetPlane()
	p.SetID(n)
	p.SetElement("collaboration", bm.CollaborationID)
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
 * @GetDiagram -> models.Diagram
**/

func (bm blackBoxModel) GetCollaboration() *models.Collaboration {
	return bm.def.GetCollaboration()
}

func (bm blackBoxModel) GetParticipantCustomer(e *models.Collaboration) *models.Participant {
	return e.GetParticipant(0)
}

func (bm blackBoxModel) GetParticipantCustomerSupport(e *models.Collaboration) *models.Participant {
	return e.GetParticipant(1)
}

func (bm blackBoxModel) GetParticipantManufacturer(e *models.Collaboration) *models.Participant {
	return e.GetParticipant(2)
}

func (bm blackBoxModel) GetMessageOrder() *models.MessageFlow {
	return bm.GetCollaboration().GetMessageFlow(0)
}

func (bm blackBoxModel) GetMessageRequestSpareParts() *models.MessageFlow {
	return bm.GetCollaboration().GetMessageFlow(1)
}

func (bm blackBoxModel) GetMessageReplacementSupply() *models.MessageFlow {
	return bm.GetCollaboration().GetMessageFlow(2)
}

func (bm blackBoxModel) GetMessageConfirmation() *models.MessageFlow {
	return bm.GetCollaboration().GetMessageFlow(3)
}

func (bm blackBoxModel) GetMessageShipment() *models.MessageFlow {
	return bm.GetCollaboration().GetMessageFlow(4)
}

func (bm blackBoxModel) GetDiagram() *models.Diagram {
	return bm.def.GetDiagram(0)
}

func (bm blackBoxModel) GetPlane() *models.Plane {
	return bm.GetDiagram().GetPlane()
}

func (bm blackBoxModel) GetShapePoolCustomer(e *models.Plane) *models.Shape {
	return e.GetShape(0)
}

func (bm blackBoxModel) GetShapePoolCustomerSupport(e *models.Plane) *models.Shape {
	return e.GetShape(1)
}

func (bm blackBoxModel) GetShapePoolManufacturer(e *models.Plane) *models.Shape {
	return e.GetShape(2)
}

func (bm blackBoxModel) GetEdgeMessageOrder(e *models.Plane) *models.Edge {
	return e.GetEdge(0)
}

func (bm blackBoxModel) GetEdgeMessageRequestSpareParts(e *models.Plane) *models.Edge {
	return e.GetEdge(1)
}

func (bm blackBoxModel) GetEdgeMessageReplacementSupply(e *models.Plane) *models.Edge {
	return e.GetEdge(2)
}

func (bm blackBoxModel) GetEdgeMessageConfirmation(e *models.Plane) *models.Edge {
	return e.GetEdge(3)
}

func (bm blackBoxModel) GetEdgeMessageShipment(e *models.Plane) *models.Edge {
	return e.GetEdge(4)
}
