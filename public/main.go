package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/deemount/gobpmn/models/bpmn/artifacts"
	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/camunda"
	"github.com/deemount/gobpmn/models/bpmn/canvas"
	"github.com/deemount/gobpmn/models/bpmn/collaboration"
	"github.com/deemount/gobpmn/models/bpmn/conditional"
	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/models/bpmn/data"
	"github.com/deemount/gobpmn/models/bpmn/events/definitions"
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/gateways"
	"github.com/deemount/gobpmn/models/bpmn/loop"
	"github.com/deemount/gobpmn/models/bpmn/marker"
	"github.com/deemount/gobpmn/models/bpmn/pool"
	"github.com/deemount/gobpmn/models/bpmn/process"
	"github.com/deemount/gobpmn/models/bpmn/subprocesses"
	"github.com/deemount/gobpmn/models/bpmn/tasks"
	tm "github.com/deemount/gobpmn/models/bpmn/time"
	"github.com/deemount/gobpmn/public/views"
	"github.com/deemount/gobpmn/utils"
)

var index *views.View
var contact *views.View
var bpmnjs *views.View
var elmnts *views.View

func main() {

	index = views.NewView("bootstrap", "views/index.gohtml")
	elmnts = views.NewView("bootstrap", "views/elements.gohtml")
	bpmnjs = views.NewView("bpmnjs", "views/bpmnjs.gohtml")

	// http.Handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/elements", elementsHandler)
	mux.HandleFunc("/bpmnjs", bpmnJsHandler)
	mux.HandleFunc("/methods", receiveMethods)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

func elementsHandler(w http.ResponseWriter, r *http.Request) {

	userDir, _ := os.UserHomeDir()
	packageEnv := "/go/src/github.com/deemount/gobpmn/models/bpmn/"
	option := ``

	// get methods by elements
	err := filepath.Walk(userDir+packageEnv,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if !info.IsDir() &&
				info.Name() != ".DS_Store" &&
				info.Name() != "Reflect.go" &&
				info.Name() != "Delegates.go" &&
				info.Name() != "Structs.go" &&
				info.Name() != "Interfaces.go" &&
				info.Name() != "Types.go" {

				fileName := filepath.Base(path)
				fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

				option += `<option value="` + utils.ToSnakeCase(fileName) + `">` + filepath.Base(filepath.Dir(path)) + `.` + fileName + `</option>`

			}
			return err
		})

	if err != nil {
		log.Println(err)
	}

	elmnts.Render(w, option)
}

func receiveMethods(w http.ResponseWriter, r *http.Request) {

	var names string
	var m []string

	if r.Method == "POST" {

		methods := r.FormValue("methods")

		switch methods {

		// artifacts
		case "text":
			m = artifacts.ReflectTextMethodsToSlice()
			break
		case "text_annotation":
			m = artifacts.ReflectTextAnnotationMethodsToSlice()
			break

		// attributes
		case "documentation":
			m = attributes.ReflectDocumentationMethodsToSlice()
			break
		case "extension_elements":
			m = attributes.ReflectExtensionElementsMethodsToSlice()
			break
		case "property":
			m = attributes.ReflectPropertyMethodsToSlice()
			break

		// camunda
		case "camunda_connector":
			m = camunda.ReflectCamundaConnectorMethodsToSlice()
			break
		case "camunda_connector_id":
			m = camunda.ReflectCamundaConnectorIDMethodsToSlice()
			break
		case "camunda_constraint":
			m = camunda.ReflectCamundaConstraintMethodsToSlice()
			break
		case "camunda_execution_listener":
			m = camunda.ReflectCamundaExecutionListenerMethodsToSlice()
			break
		case "camunda_expression":
			m = camunda.ReflectCamundaExpressionMethodsToSlice()
			break
		case "camunda_failed_job_retry_cycle":
			m = camunda.ReflectCamundaFailedJobRetryCycleMethodsToSlice()
			break
		case "camunda_field":
			m = camunda.ReflectCamundaFieldMethodsToSlice()
			break
		case "camunda_form_data":
			m = camunda.ReflectCamundaFormDataMethodsToSlice()
			break
		case "camunda_form_field":
			m = camunda.ReflectCamundaFormFieldMethodsToSlice()
			break
		case "camunda_in":
			m = camunda.ReflectCamundaInMethodsToSlice()
			break
		case "camunda_input_output":
			m = camunda.ReflectCamundaInputOutputMethodsToSlice()
			break
		case "camunda_list":
			m = camunda.ReflectCamundaListMethodsToSlice()
			break
		case "camunda_map":
			m = camunda.ReflectCamundaMapMethodsToSlice()
			break
		case "camunda_out":
			m = camunda.ReflectCamundaOutMethodsToSlice()
			break
		case "camunda_output_parameter":
			m = camunda.ReflectCamundaOutputParameterMethodsToSlice()
			break
		case "camunda_properties":
			m = camunda.ReflectCamundaPropertiesMethodsToSlice()
			break
		case "camunda_property":
			m = camunda.ReflectCamundaPropertyMethodsToSlice()
			break
		case "camunda_script":
			m = camunda.ReflectCamundaScriptMethodsToSlice()
			break
		case "camunda_string":
			m = camunda.ReflectCamundaStringMethodsToSlice()
			break
		case "camunda_task_listener":
			m = camunda.ReflectCamundaTaskListenerMethodsToSlice()
			break
		case "camunda_validation":
			m = camunda.ReflectCamundaValidationMethodsToSlice()
			break
		case "camunda_value":
			m = camunda.ReflectCamundaValueMethodsToSlice()
			break

		// canvas
		case "bounds":
			m = canvas.ReflectBoundsMethodsToSlice()
			break
		case "diagram":
			m = canvas.ReflectDiagramMethodsToSlice()
			break
		case "edge":
			m = canvas.ReflectEdgeMethodsToSlice()
			break
		case "label":
			m = canvas.ReflectLabelMethodsToSlice()
			break
		case "plane":
			m = canvas.ReflectPlaneMethodsToSlice()
			break
		case "shape":
			m = canvas.ReflectShapeMethodsToSlice()
			break
		case "waypoint":
			m = canvas.ReflectWaypointMethodsToSlice()
			break

		// collaboration
		case "collaboration":
			m = collaboration.ReflectCollaborationMethodsToSlice()
			break
		case "participant":
			m = collaboration.ReflectParticipantMethodsToSlice()
			break

		// conditional
		case "completion_condition":
			m = conditional.ReflectCompletionConditionMethodsToSlice()
			break
		case "condition":
			m = conditional.ReflectConditionMethodsToSlice()
			break
		case "condition_expression":
			m = conditional.ReflectConditionExpressionMethodsToSlice()
			break

		// core
		case "definitions":
			m = core.ReflectDefinitionsMethodsToSlice()
			break

		// data
		case "data_object":
			m = data.ReflectDataObjectMethodsToSlice()
			break
		case "data_object_reference":
			m = data.ReflectDataObjectReferenceMethodsToSlice()
			break
		case "data_store_reference":
			m = data.ReflectDataStoreReferenceMethodsToSlice()
			break

		// event definitions
		case "cancel_event_definition":
			m = definitions.ReflectCancelEventDefinitionMethodsToSlice()
			break
		case "compensate_event_defintion":
			m = definitions.ReflectCompensateEventDefinitionMethodsToSlice()
			break
		case "conditional_event_definition":
			m = definitions.ReflectConditionalEventDefinitionMethodsToSlice()
			break
		case "escalation_event_definition":
			m = definitions.ReflectEscalationEventDefinitionMethodsToSlice()
			break
		case "link_event_definition":
			m = definitions.ReflectLinkEventDefinitionMethodsToSlice()
			break
		case "message_event_definition":
			m = definitions.ReflectMessageEventDefinitionMethodsToSlice()
			break
		case "signal_event_definition":
			m = definitions.ReflectSignalEventDefinitionMethodsToSlice()
			break
		case "terminate_event_definition":
			m = definitions.ReflectTerminateEventDefinitionMethodsToSlice()
			break
		case "timer_event_definition":
			m = definitions.ReflectTimerEventDefinitionMethodsToSlice()
			break

		// elements
		case "boundary_event":
			m = elements.ReflectBoundaryEventMethodsToSlice()
			break
		case "end_event":
			m = elements.ReflectEndEventMethodsToSlice()
			break
		case "intermediate_catch_event":
			m = elements.ReflectIntermediateCatchEventMethodsToSlice()
			break
		case "intermediate_throw_event":
			m = elements.ReflectIntermediateThrowEventMethodsToSlice()
			break
		case "message":
			m = elements.ReflectMessageMethodsToSlice()
			break
		case "signal":
			m = elements.ReflectSignalMethodsToSlice()
			break
		case "start_event":
			m = elements.ReflectStartEventlMethodsToSlice()
			break

		// flow
		case "association":
			m = flow.ReflectAssociationMethodsToSlice()
			break
		case "data_input_association":
			m = flow.ReflectDataInputAssociationMethodsToSlice()
			break
		case "message_flow":
			m = flow.ReflectMessageFlowMethodsToSlice()
			break
		case "sequence_flow":
			m = flow.ReflectSequenceFlowMethodsToSlice()
			break

		// gateways
		case "complex_gateway":
			m = gateways.ReflectComplexGatewayMethodsToSlice()
			break
		case "event_based_gateway":
			m = gateways.ReflectEventBasedGatewayMethodsToSlice()
			break
		case "exclusive_gateway":
			m = gateways.ReflectExclusiveGatewayMethodsToSlice()
			break
		case "inclsusive_gateway":
			m = gateways.ReflectInclusiveGatewayMethodsToSlice()
			break
		case "parallel_gateway":
			m = gateways.ReflectParallelGatewayMethodsToSlice()
			break

		// loop
		case "loop_cardinality":
			m = loop.ReflectLoopCardinalityMethodsToSlice()
			break
		case "multi_instance_loop_characteristics":
			m = loop.ReflectMultiInstanceLoopCharacteristicsMethodsToSlice()
			break
		case "participant_multiplicity":
			m = loop.ReflectParticipantMultiplicityMethodsToSlice()
			break
		case "standard_loop_characteristics":
			m = loop.ReflectStandardLoopCharacteristicsMethodsToSlice()
			break

		// marker
		case "category":
			m = marker.ReflectCategoryMethodsToSlice()
			break
		case "category_value":
			m = marker.ReflectCategoryValueMethodsToSlice()
			break
		case "group":
			m = marker.ReflectGroupMethodsToSlice()
			break
		case "incoming":
			m = marker.ReflectIncomingMethodsToSlice()
			break
		case "outgoing":
			m = marker.ReflectOutgoingMethodsToSlice()
			break

		// pool
		case "flow_node_ref":
			m = pool.ReflectFlowNodeRefMethodsToSlice()
			break
		case "lane":
			m = pool.ReflectLaneMethodsToSlice()
			break
		case "lane_set":
			m = pool.ReflectLaneSetMethodsToSlice()
			break

		// process
		case "process":
			m = process.ReflectProcessMethodsToSlice()
			break

		// subprocesses
		case "ad_hoc_sub_process":
			m = subprocesses.ReflectAdHocProcessMethodsToSlice()
			break
		case "call_activity":
			m = subprocesses.ReflectCallActivityMethodsToSlice()
			break
		case "sub_process":
			m = subprocesses.ReflectSubProcessMethodsToSlice()
			break
		case "transaction":
			m = subprocesses.ReflectTransactionMethodsToSlice()
			break

		// tasks
		case "business_rule_task":
			m = tasks.ReflectBusinessRuleTaskMethodsToSlice()
			break
		case "manual_task":
			m = tasks.ReflectManualTaskMethodsToSlice()
			break
		case "receive_task":
			m = tasks.ReflectReceiveTaskMethodsToSlice()
			break
		case "script_task":
			m = tasks.ReflectScriptTaskMethodsToSlice()
			break
		case "send_task":
			m = tasks.ReflectSendTaskMethodsToSlice()
			break
		case "service_task":
			m = tasks.ReflectServiceTaskMethodsToSlice()
			break
		case "task":
			m = tasks.ReflectTaskMethodsToSlice()
			break
		case "user_task":
			m = tasks.ReflectUserTaskMethodsToSlice()
			break

		// time
		case "time_cycle":
			m = tm.ReflectTimeCycleMethodsToSlice()
			break
		case "time_date":
			m = tm.ReflectTimeDateMethodsToSlice()
			break
		case "time_duration":
			m = tm.ReflectTimeDurationMethodsToSlice()
			break
		}

		for _, name := range m {
			names += `<button class="list-group-item list-group-item-action">` + name + `</button>`
		}

		methods = strings.Replace(methods, "_", " ", -1)
		methods = strings.Title(strings.ToLower(methods))
		methods = strings.Replace(methods, " ", "", -1)
		html := `<div class="list-group">` + names + `</div>`
		log.Printf("configuration: received methods from %s", methods)

		if _, err := w.Write([]byte(html)); err != nil {
			panic(err)
		}
	}
}

func bpmnJsHandler(w http.ResponseWriter, r *http.Request) {
	bpmnjs.Render(w, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index.Render(w, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	contact.Render(w, nil)
}
