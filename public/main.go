package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	"github.com/deemount/gobpmn/models/bpmn/time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	html := `<head>
                    <script src='//ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js'></script>
              </head>
                  <html>
				  <body>
                  <h1>Methods</h1>
                  <p>
				  <select>
				    <option value="text">artifacts.Text</option>
					<option value="text_annotation">artifacts.TextAnnotation</option>
				    <option value="documentation">attributes.Documentation</option>
					<option value="extension_elements">attributes.ExtensionElements</option>
					<option value="property">attributes.Property</option>
					<option value="bounds">canvas.Bounds</option>
					<option value="diagram">canvas.Diagram</option>
					<option value="edge">canvas.Edge</option>
					<option value="label">canvas.Label</option>
					<option value="plane">canvas.Plane</option>
					<option value="plane">canvas.Shape</option>
					<option value="plane">canvas.Waypoint</option>
					<option value="camunda_connector">camunda.CamundaConnector</option>
					<option value="camunda_connector_id">camunda.CamundaConnectorID</option>
					<option value="camunda_constraint">camunda.CamundaConstraint</option>
					<option value="camunda_entry">camunda.CamundaEntry</option>
					<option value="camunda_execution_listener">camunda.CamundaExecutionListener</option>
					<option value="camunda_expression">camunda.CamundaExpression</option>
					<option value="camunda_failed_job_retry_cycle">camunda.CamundaFailedJobretryCycle</option>
					<option value="camunda_field">camunda.CamundaField</option>
					<option value="camunda_form_data">camunda.CamundaFormData</option>
					<option value="camunda_form_field">camunda.CamundaFormField</option>
					<option value="camunda_in">camunda.CamundaIn</option>
					<option value="camunda_input_output">camunda.CamundaInputOutput</option>
					<option value="camunda_input_parameter">camunda.CamundaInpuParameter</option>
					<option value="camunda_list">camunda.CamundaList</option>
					<option value="camunda_map">camunda.CamundaMap</option>
					<option value="camunda_out">camunda.CamundaOut</option>
					<option value="camunda_output_parameter">camunda.CamundaOutputParameter</option>
					<option value="camunda_properties">camunda.CamundaProperties</option>
					<option value="camunda_property">camunda.CamundaProperty</option>
					<option value="camunda_script">camunda.CamundaScript</option>
					<option value="camunda_string">camunda.CamundaString</option>
					<option value="camunda_task_listener">camunda.CamundaTaskListener</option>
					<option value="camunda_validation">camunda.CamundaValidation</option>
					<option value="camunda_value">camunda.CamundaValue</option>
					<option value="collaboration">collaboration.Collaboration</option>
					<option value="participant">collaboration.Participant</option>
					<option value="completion_condition">conditional.CompletionCondition</option>
					<option value="condition">conditional.Condition</option>
					<option value="condition_expression">conditional.ConditionExpression</option>
					<option value="definitions">core.Definitions</option>
					<option value="data_object">data.DataObject</option>
					<option value="data_object_reference">data.DataObjectReference</option>
					<option value="data_store_reference">data.DataStoreReference</option>
					<option value="cancel_event_definition">definitions.CancelEventDefinition</option>
					<option value="compensate_event_definition">definitions.CompensateEventDefinition</option>
					<option value="conditional_event_definition">definitions.ConditionalEventDefinition</option>
					<option value="error_event_definition">definitions.ErrorEventDefinition</option>
					<option value="escalation_event_definition">definitions.EscalationEventDefinition</option>
					<option value="link_event_definition">definitions.LinkEventDefinition</option>
					<option value="message_event_definition">definitions.MessageEventDefinition</option>
					<option value="signal_event_definition">definitions.SignalEventDefinition</option>
					<option value="terminate_event_definition">definitions.TerminateEventDefinition</option>
					<option value="timer_event_definition">definitions.TimerEventDefinition</option>
					<option value="boundary_event">elements.BoundaryEvent</option>
					<option value="end_event">elements.EndEvent</option>
					<option value="intermediate_catch_event">elements.IntermediateCatchEvent</option>
					<option value="intermediate_throw_event">elements.IntermediateThrowEvent</option>
					<option value="message">elements.Message</option>
					<option value="signal">elements.Signal</option>
					<option value="start_event">elements.StartEvent</option>
					<option value="association">flow.Association</option>
					<option value="data_input_association">flow.DataInputAssociation</option>
					<option value="message_flow">flow.MessageFlow</option>
					<option value="sequence_flow">flow.SequenceFlow</option>
					<option value="complex_gateway">gateways.ComplexGateway</option>
					<option value="event_based_gateway">gateways.EventBasedGateway</option>
					<option value="exclusive_gateway">gateways.ExclusiveGateway</option>
					<option value="inclusive_gateway">gateways.InclusiveGateway</option>
					<option value="parallel_gateway">gateways.ParallelGateway</option>
					<option value="loop_cardinality">loop.LoopCardinality</option>
					<option value="multi_instance_loop_characteristics">loop.MultiInstanceLoopCharacteristics</option>
					<option value="participant_multiplicity">loop.ParticipantMultiplicity</option>
					<option value="standard_loop_characteristics">loop.StandardLoopCharacteristics</option>
					<option value="category">marker.Category</option>
					<option value="category_value">marker.CategoryValue</option>
					<option value="group">marker.Group</option>
					<option value="incoming">marker.Incoming</option>
					<option value="outgoing">marker.Outgoing</option>
					<option value="flow_node_ref">pool.FlowNodeRef</option>
					<option value="lane">pool.Lane</option>
					<option value="lane_set">pool.LaneSet</option>
					<option value="process">process.Process</option>
					<option value="adhoc_subprocess">subprocesses.AdHocSubprocess</option>
					<option value="call_activity">subprocesses.CallActivity</option>
					<option value="subprocess">subprocesses.SubProcess</option>
					<option value="transaction">subprocesses.Transaction</option>
					<option value="business_rule_task">tasks.BusinessRuleTask</option>
					<option value="manual_task">tasks.ManualTask</option>
					<option value="receive_task">tasks.ReceiveTask</option>
					<option value="script_task">tasks.ScriptTask</option>
					<option value="send_task">tasks.SendTask</option>
					<option value="service_task">tasks.ServiceTask</option>
					<option value="task">tasks.Task</option>
					<option value="user_task">tasks.UserTask</option>
					<option value="time_cycle">time.TimeCycle</option>
					<option value="time_date">time.TimeDate</option>
					<option value="time_duration">time.TimeDuration</option>
				  </select>
				  <input id='methods_btn' type='button' value='methods'>
				  </p>
				  <div id='result'></div><br><br>
                  </body></html>

                   <script>
                   $(document).ready(function () {
                         $('#methods_btn').click(function () {
                             $.ajax({
                               url: 'methods',
                               type: 'post',
                               dataType: 'html',
                               data : { methods: $("select option:selected").val()},
                               success : function(data) {
                                 $('#result').html(data);
                               },
                             });
                          });
                	});
                    </script>`

	w.Write([]byte(fmt.Sprintf(html)))

}

func receiveMethods(w http.ResponseWriter, r *http.Request) {

	var names string
	var m map[int]string

	if r.Method == "POST" {
		methods := r.FormValue("methods")

		switch methods {

		// artifacts
		case "text":
			m = artifacts.ReflectTextMethodsToMap()
			break
		case "text_annotation":
			m = artifacts.ReflectTextAnnotationMethodsToMap()
			break

		// attributes
		case "documentation":
			m = attributes.ReflectDocumentationMethodsToMap()
			break
		case "extension_elements":
			m = attributes.ReflectExtensionElementsMethodsToMap()
			break
		case "property":
			m = attributes.ReflectPropertyMethodsToMap()
			break

		// camunda
		case "camunda_connector":
			m = camunda.ReflectCamundaConnectorMethodsToMap()
			break
		case "camunda_connector_id":
			m = camunda.ReflectCamundaConnectorIDMethodsToMap()
			break
		case "camunda_constraint":
			m = camunda.ReflectCamundaConstraintMethodsToMap()
			break
		case "camunda_execution_listener":
			m = camunda.ReflectCamundaExecutionListenerMethodsToMap()
			break
		case "camunda_expression":
			m = camunda.ReflectCamundaExpressionMethodsToMap()
			break
		case "camunda_failed_job_retry_cycle":
			m = camunda.ReflectCamundaFailedJobRetryCycleMethodsToMap()
			break
		case "camunda_field":
			m = camunda.ReflectCamundaFieldMethodsToMap()
			break
		case "camunda_form_data":
			m = camunda.ReflectCamundaFormDataMethodsToMap()
			break
		case "camunda_form_field":
			m = camunda.ReflectCamundaFormFieldMethodsToMap()
			break
		case "camunda_in":
			m = camunda.ReflectCamundaInMethodsToMap()
			break
		case "camunda_input_output":
			m = camunda.ReflectCamundaInputOutputMethodsToMap()
			break
		case "camunda_list":
			m = camunda.ReflectCamundaListMethodsToMap()
			break
		case "camunda_map":
			m = camunda.ReflectCamundaMapMethodsToMap()
			break
		case "camunda_out":
			m = camunda.ReflectCamundaOutMethodsToMap()
			break
		case "camunda_output_parameter":
			m = camunda.ReflectCamundaOutputParameterMethodsToMap()
			break
		case "camunda_properties":
			m = camunda.ReflectCamundaPropertiesMethodsToMap()
			break
		case "camunda_property":
			m = camunda.ReflectCamundaPropertyMethodsToMap()
			break
		case "camunda_script":
			m = camunda.ReflectCamundaScriptMethodsToMap()
			break
		case "camunda_string":
			m = camunda.ReflectCamundaStringMethodsToMap()
			break
		case "camunda_task_listener":
			m = camunda.ReflectCamundaTaskListenerMethodsToMap()
			break
		case "camunda_validation":
			m = camunda.ReflectCamundaValidationMethodsToMap()
			break
		case "camunda_value":
			m = camunda.ReflectCamundaValueMethodsToMap()
			break

		// canvas
		case "bounds":
			m = canvas.ReflectBoundsMethodsToMap()
			break
		case "diagram":
			m = canvas.ReflectDiagramMethodsToMap()
			break
		case "edge":
			m = canvas.ReflectEdgeMethodsToMap()
			break
		case "label":
			m = canvas.ReflectLabelMethodsToMap()
			break
		case "plane":
			m = canvas.ReflectPlaneMethodsToMap()
			break
		case "shape":
			m = canvas.ReflectShapeMethodsToMap()
			break
		case "waypoint":
			m = canvas.ReflectWaypointMethodsToMap()
			break

		// collaboration
		case "collaboration":
			m = collaboration.ReflectCollaborationMethodsToMap()
			break
		case "participant":
			m = collaboration.ReflectParticipantMethodsToMap()
			break

		// conditional
		case "completion_condition":
			m = conditional.ReflectCompletionConditionMethodsToMap()
			break
		case "condition":
			m = conditional.ReflectConditionMethodsToMap()
			break
		case "condition_expression":
			m = conditional.ReflectConditionExpressionMethodsToMap()
			break

		// core
		case "definitions":
			m = core.ReflectDefinitionsMethodsToMap()
			break

		// data
		case "data_object":
			m = data.ReflectDataObjectMethodsToMap()
			break
		case "data_object_reference":
			m = data.ReflectDataObjectReferenceMethodsToMap()
			break
		case "data_store_reference":
			m = data.ReflectDataStoreReferenceMethodsToMap()
			break

		// event definitions
		case "cancel_event_definition":
			m = definitions.ReflectCancelEventDefinitionMethodsToMap()
			break
		case "compensate_event_defintion":
			m = definitions.ReflectCompensateEventDefinitionMethodsToMap()
			break
		case "conditional_event_definition":
			m = definitions.ReflectConditionalEventDefinitionMethodsToMap()
			break
		case "escalation_event_definition":
			m = definitions.ReflectEscalationEventDefinitionMethodsToMap()
			break
		case "link_event_definition":
			m = definitions.ReflectLinkEventDefinitionMethodsToMap()
			break
		case "message_event_definition":
			m = definitions.ReflectMessageEventDefinitionMethodsToMap()
			break
		case "signal_event_definition":
			m = definitions.ReflectSignalEventDefinitionMethodsToMap()
			break
		case "terminate_event_definition":
			m = definitions.ReflectTerminateEventDefinitionMethodsToMap()
			break
		case "timer_event_definition":
			m = definitions.ReflectTimerEventDefinitionMethodsToMap()
			break

		// elements
		case "boundary_event":
			m = elements.ReflectBoundaryEventMethodsToMap()
			break
		case "end_event":
			m = elements.ReflectEndEventMethodsToMap()
			break
		case "intermediate_catch_event":
			m = elements.ReflectIntermediateCatchEventMethodsToMap()
			break
		case "intermediate_throw_event":
			m = elements.ReflectIntermediateThrowEventMethodsToMap()
			break
		case "message":
			m = elements.ReflectMessageMethodsToMap()
			break
		case "signal":
			m = elements.ReflectSignalMethodsToMap()
			break
		case "start_event":
			m = elements.ReflectStartEventlMethodsToMap()
			break

		// flow
		case "association":
			m = flow.ReflectAssociationMethodsToMap()
			break
		case "data_input_association":
			m = flow.ReflectDataInputAssociationMethodsToMap()
			break
		case "message_flow":
			m = flow.ReflectMessageFlowMethodsToMap()
			break
		case "sequence_flow":
			m = flow.ReflectSequenceFlowMethodsToMap()
			break

		// gateways
		case "complex_gateway":
			m = gateways.ReflectComplexGatewayMethodsToMap()
			break
		case "event_based_gateway":
			m = gateways.ReflectEventBasedGatewayMethodsToMap()
			break
		case "exclusive_gateway":
			m = gateways.ReflectExclusiveGatewayMethodsToMap()
			break
		case "inclsusive_gateway":
			m = gateways.ReflectInclusiveGatewayMethodsToMap()
			break
		case "parallel_gateway":
			m = gateways.ReflectParallelGatewayMethodsToMap()
			break

		// loop
		case "loop_cardinality":
			m = loop.ReflectLoopCardinalityMethodsToMap()
			break
		case "multi_instance_loop_characteristics":
			m = loop.ReflectMultiInstanceLoopCharacteristicsMethodsToMap()
			break
		case "participant_multiplicity":
			m = loop.ReflectParticipantMultiplicityMethodsToMap()
			break
		case "standard_loop_characteristics":
			m = loop.ReflectStandardLoopCharacteristicsMethodsToMap()
			break

		// marker
		case "category":
			m = marker.ReflectCategoryMethodsToMap()
			break
		case "category_value":
			m = marker.ReflectCategoryValueMethodsToMap()
			break
		case "group":
			m = marker.ReflectGroupMethodsToMap()
			break
		case "incoming":
			m = marker.ReflectIncomingMethodsToMap()
			break
		case "outgoing":
			m = marker.ReflectOutgoingMethodsToMap()
			break

		// pool
		case "flow_node_ref":
			m = pool.ReflectFlowNodeRefMethodsToMap()
			break
		case "lane":
			m = pool.ReflectLaneMethodsToMap()
			break
		case "lane_set":
			m = pool.ReflectLaneSetMethodsToMap()
			break

		// process
		case "process":
			m = process.ReflectProcessMethodsToMap()
			break

		// subprocesses
		case "adhoc_subprocess":
			m = subprocesses.ReflectAdHocProcessMethodsToMap()
			break
		case "call_activity":
			m = subprocesses.ReflectCallActivityMethodsToMap()
			break
		case "subprocess":
			m = subprocesses.ReflectSubProcessMethodsToMap()
			break
		case "transaction":
			m = subprocesses.ReflectTransactionMethodsToMap()
			break

		// tasks
		case "business_rule_task":
			m = tasks.ReflectBusinessRuleTaskMethodsToMap()
			break
		case "manual_task":
			m = tasks.ReflectManualTaskMethodsToMap()
			break
		case "receive_task":
			m = tasks.ReflectReceiveTaskMethodsToMap()
			break
		case "script_task":
			m = tasks.ReflectScriptTaskMethodsToMap()
			break
		case "send_task":
			m = tasks.ReflectSendTaskMethodsToMap()
			break
		case "service_task":
			m = tasks.ReflectServiceTaskMethodsToMap()
			break
		case "task":
			m = tasks.ReflectTaskMethodsToMap()
			break
		case "user_task":
			m = tasks.ReflectUserTaskMethodsToMap()
			break

		// time
		case "time_cycle":
			m = time.ReflectTimeCycleMethodsToMap()
			break
		case "time_date":
			m = time.ReflectTimeDateMethodsToMap()
			break
		case "time_duration":
			m = time.ReflectTimeDurationMethodsToMap()
			break
		}

		for _, name := range m {
			names += name + "<br>"
		}

		methods = strings.Replace(methods, "_", " ", -1)
		methods = strings.Title(strings.ToLower(methods))
		methods = strings.Replace(methods, " ", "", -1)
		log.Printf("Receive methods %s", methods)

		w.Write([]byte("<p>" + names + "<p>"))
	}
}

func main() {
	// http.Handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/methods", receiveMethods)

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	http.ListenAndServe(":8080", mux)
}
