<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="2.0"
	xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xmi="http://schema.omg.org/spec/XMI"
	xmlns:bpmn="http://www.omg.org/spec/BPMN/20100524/MODEL" xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI"
	xmlns:di="http://www.omg.org/spec/DD/20100524/DI" xmlns:dc="http://www.omg.org/spec/DD/20100524/DC"
	xmlns:bpmnxmi="http://www.omg.org/spec/BPMN/20100524/MODEL-XMI"
	xmlns:bpmndixmi="http://www.omg.org/spec/BPMN/20100524/DI-XMI"
	xmlns:dixmi="http://www.omg.org/spec/DD/20100524/DI-XMI" xmlns:dcxmi="http://www.omg.org/spec/DD/20100524/DC-XMI">
	<xsl:output method="xml" version="1.0" encoding="UTF-8" indent="yes" />

	<xsl:template name="concat">
		<xsl:param name="nodeset" />
		<xsl:if test="$nodeset">
			<xsl:variable name="rest">
				<xsl:call-template name="concat">
					<xsl:with-param name="nodeset"
						select="$nodeset[not (position() = 1)]" />
				</xsl:call-template>
			</xsl:variable>
			<xsl:value-of select="concat($nodeset[1], ' ', $rest)" />
		</xsl:if>
	</xsl:template>

	<xsl:template match="/">
		<xmi:XMI xmi:version="2.0">
			<xsl:for-each select="bpmn:definitions">
				<bpmnxmi:Definitions>
					<xsl:call-template name="DefinitionsTemplate" />
				</bpmnxmi:Definitions>
			</xsl:for-each>
		</xmi:XMI>
	</xsl:template>

	<xsl:template name="EObjectTemplate">
		<!-- TODO: Add your copy logic for extension attributes-->
	</xsl:template>

	<xsl:template name="InterfaceTemplate">
		<xsl:call-template name="RootElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@implementationRef">
			<xsl:attribute name="implementationRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@implementationRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:operation">
			<operations xmi:type="bpmnxmi:Operation">
				<xsl:call-template name="OperationTemplate" />
			</operations>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="RootElementTemplate">
		<xsl:call-template name="BaseElementTemplate" />
	</xsl:template>

	<xsl:template name="BaseElementTemplate">
		<xsl:if test="@id">
			<xsl:attribute name="id"> <xsl:value-of select="@id" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:extensionElements">
			<xsl:attribute name="extensionDefinitions"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:extensionElements" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:extensionValues">
			<extensionValues xmi:type="bpmnxmi:ExtensionAttributeValue">
				<xsl:call-template name="ExtensionAttributeValueTemplate" />
			</extensionValues>
		</xsl:for-each>
		<xsl:for-each select="bpmn:documentation">
			<documentation xmi:type="bpmnxmi:Documentation">
				<xsl:call-template name="DocumentationTemplate" />
			</documentation>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="ExtensionDefinitionTemplate">
		<xsl:if test="bpmn:name">
			<xsl:attribute name="name"> <xsl:value-of select="bpmn:name" /> </xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:extensionAttributeDefinitions">
			<extensionAttributeDefinitions
				xmi:type="bpmnxmi:ExtensionAttributeDefinition">
				<xsl:call-template name="ExtensionAttributeDefinitionTemplate" />
			</extensionAttributeDefinitions>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="ExtensionAttributeDefinitionTemplate">
		<xsl:if test="bpmn:name">
			<xsl:attribute name="name"> <xsl:value-of select="bpmn:name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:type">
			<xsl:attribute name="type"> <xsl:value-of select="bpmn:type" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:isReference">
			<xsl:attribute name="isReference"> <xsl:value-of select="bpmn:isReference" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:extensionDefinition">
			<xsl:attribute name="extensionDefinition"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:extensionDefinition" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ExtensionAttributeValueTemplate">
		<xsl:if test="bpmn:valueRef">
			<xsl:attribute name="valueRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:valueRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:extensionAttributeDefinition">
			<xsl:attribute name="extensionAttributeDefinition"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset"
				select="bpmn:extensionAttributeDefinition" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:value">
			<value xmi:type="bpmnxmi:EObject">
				<xsl:call-template name="EObjectTemplate" />
			</value>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="DocumentationTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="bpmn:text">
			<xsl:attribute name="text"> <xsl:value-of select="bpmn:text" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@textFormat">
			<xsl:attribute name="textFormat"> <xsl:value-of select="@textFormat" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="OperationTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:inMessageRef">
			<xsl:attribute name="inMessageRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:inMessageRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:outMessageRef">
			<xsl:attribute name="outMessageRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:outMessageRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:errorRef">
			<xsl:attribute name="errorRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:errorRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@implementationRef">
			<xsl:attribute name="implementationRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@implementationRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="MessageTemplate">
		<xsl:call-template name="RootElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@itemRef">
			<xsl:attribute name="itemRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@itemRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ItemDefinitionTemplate">
		<xsl:call-template name="RootElementTemplate" />
		<xsl:if test="@itemKind">
			<xsl:attribute name="itemKind"> <xsl:value-of select="@itemKind" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@isCollection">
			<xsl:attribute name="isCollection"> <xsl:value-of select="@isCollection" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@structureRef">
			<xsl:attribute name="structureRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@structureRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:import">
			<xsl:attribute name="import"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:import" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ImportTemplate">
		<xsl:if test="@importType">
			<xsl:attribute name="importType"> <xsl:value-of select="@importType" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@location">
			<xsl:attribute name="location"> <xsl:value-of select="@location" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@namespace">
			<xsl:attribute name="namespace"> <xsl:value-of select="@namespace" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ErrorTemplate">
		<xsl:call-template name="RootElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@errorCode">
			<xsl:attribute name="errorCode"> <xsl:value-of select="@errorCode" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@structureRef">
			<xsl:attribute name="structureRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@structureRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="EndPointTemplate">
		<xsl:call-template name="RootElementTemplate" />
	</xsl:template>

	<xsl:template name="AuditingTemplate">
		<xsl:call-template name="BaseElementTemplate" />
	</xsl:template>

	<xsl:template name="GlobalTaskTemplate">
		<xsl:call-template name="CallableElementTemplate" />
		<xsl:for-each select="bpmn:performer">
			<resources xmi:type="bpmnxmi:Performer">
				<xsl:call-template name="PerformerTemplate" />
			</resources>
		</xsl:for-each>
		<xsl:for-each select="bpmn:humanPerformer">
			<resources xmi:type="bpmnxmi:HumanPerformer">
				<xsl:call-template name="HumanPerformerTemplate" />
			</resources>
		</xsl:for-each>
		<xsl:for-each select="bpmn:potentialOwner">
			<resources xmi:type="bpmnxmi:PotentialOwner">
				<xsl:call-template name="PotentialOwnerTemplate" />
			</resources>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="CallableElementTemplate">
		<xsl:call-template name="RootElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:supportedInterfaceRef">
			<xsl:attribute name="supportedInterfaceRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset"
				select="bpmn:supportedInterfaceRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:ioSpecification">
			<ioSpecification xmi:type="bpmnxmi:InputOutputSpecification">
				<xsl:call-template name="InputOutputSpecificationTemplate" />
			</ioSpecification>
		</xsl:for-each>
		<xsl:for-each select="bpmn:ioBinding">
			<ioBinding xmi:type="bpmnxmi:InputOutputBinding">
				<xsl:call-template name="InputOutputBindingTemplate" />
			</ioBinding>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="InputOutputSpecificationTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:for-each select="bpmn:inputSet">
			<inputSets xmi:type="bpmnxmi:InputSet">
				<xsl:call-template name="InputSetTemplate" />
			</inputSets>
		</xsl:for-each>
		<xsl:for-each select="bpmn:outputSet">
			<outputSets xmi:type="bpmnxmi:OutputSet">
				<xsl:call-template name="OutputSetTemplate" />
			</outputSets>
		</xsl:for-each>
		<xsl:for-each select="bpmn:dataInput">
			<dataInputs xmi:type="bpmnxmi:DataInput">
				<xsl:call-template name="DataInputTemplate" />
			</dataInputs>
		</xsl:for-each>
		<xsl:for-each select="bpmn:dataOutput">
			<dataOutputs xmi:type="bpmnxmi:DataOutput">
				<xsl:call-template name="DataOutputTemplate" />
			</dataOutputs>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="InputSetTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:dataInputRefs">
			<xsl:attribute name="dataInputRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:dataInputRefs" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:optionalInputRefs">
			<xsl:attribute name="optionalInputRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:optionalInputRefs" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:whileExecutingInputRefs">
			<xsl:attribute name="whileExecutingInputRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset"
				select="bpmn:whileExecutingInputRefs" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:outputSetRefs">
			<xsl:attribute name="outputSetRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:outputSetRefs" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="DataInputTemplate">
		<xsl:call-template name="ItemAwareElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@isCollection">
			<xsl:attribute name="isCollection"> <xsl:value-of select="@isCollection" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:inputSetWithOptional">
			<xsl:attribute name="inputSetWithOptional"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:inputSetWithOptional" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:inputSetWithWhileExecuting">
			<xsl:attribute name="inputSetWithWhileExecuting"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset"
				select="bpmn:inputSetWithWhileExecuting" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:inputSetRefs">
			<xsl:attribute name="inputSetRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:inputSetRefs" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ItemAwareElementTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@itemSubjectRef">
			<xsl:attribute name="itemSubjectRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@itemSubjectRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:dataState">
			<dataState xmi:type="bpmnxmi:DataState">
				<xsl:call-template name="DataStateTemplate" />
			</dataState>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="DataStateTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="OutputSetTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:dataOutputRefs">
			<xsl:attribute name="dataOutputRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:dataOutputRefs" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:optionalOutputRefs">
			<xsl:attribute name="optionalOutputRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:optionalOutputRefs" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:whileExecutingOutputRefs">
			<xsl:attribute name="whileExecutingOutputRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset"
				select="bpmn:whileExecutingOutputRefs" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:inputSetRefs">
			<xsl:attribute name="inputSetRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:inputSetRefs" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="DataOutputTemplate">
		<xsl:call-template name="ItemAwareElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@isCollection">
			<xsl:attribute name="isCollection"> <xsl:value-of select="@isCollection" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:outputSetWithOptional">
			<xsl:attribute name="outputSetWithOptional"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset"
				select="bpmn:outputSetWithOptional" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:outputSetWithWhileExecuting">
			<xsl:attribute name="outputSetWithWhileExecuting"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset"
				select="bpmn:outputSetWithWhileExecuting" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:outputSetRefs">
			<xsl:attribute name="outputSetRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:outputSetRefs" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="InputOutputBindingTemplate">
		<xsl:if test="@inputDataRef">
			<xsl:attribute name="inputDataRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@inputDataRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@outputDataRef">
			<xsl:attribute name="outputDataRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@outputDataRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@operationRef">
			<xsl:attribute name="operationRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@operationRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ResourceRoleTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:resourceRef">
			<xsl:attribute name="resourceRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:resourceRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:resourceParameterBinding">
			<resourceParameterBindings xmi:type="bpmnxmi:ResourceParameterBinding">
				<xsl:call-template name="ResourceParameterBindingTemplate" />
			</resourceParameterBindings>
		</xsl:for-each>
		<xsl:for-each select="bpmn:resourceAssignmentExpression">
			<resourceAssignmentExpression xmi:type="bpmnxmi:ResourceAssignmentExpression">
				<xsl:call-template name="ResourceAssignmentExpressionTemplate" />
			</resourceAssignmentExpression>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="ResourceTemplate">
		<xsl:call-template name="RootElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:resourceParameter">
			<resourceParameters xmi:type="bpmnxmi:ResourceParameter">
				<xsl:call-template name="ResourceParameterTemplate" />
			</resourceParameters>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="ResourceParameterTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@isRequired">
			<xsl:attribute name="isRequired"> <xsl:value-of select="@isRequired" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@type">
			<xsl:attribute name="type"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@type" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ResourceParameterBindingTemplate">
		<xsl:if test="@parameterRef">
			<xsl:attribute name="parameterRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@parameterRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:formalExpression">
			<expression xmi:type="bpmnxmi:FormalExpression">
				<xsl:call-template name="FormalExpressionTemplate" />
			</expression>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="ExpressionTemplate">
		<xsl:call-template name="BaseElementTemplate" />
	</xsl:template>

	<xsl:template name="ResourceAssignmentExpressionTemplate">
		<xsl:for-each select="bpmn:formalExpression">
			<expression xmi:type="bpmnxmi:FormalExpression">
				<xsl:call-template name="FormalExpressionTemplate" />
			</expression>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="MonitoringTemplate">
		<xsl:call-template name="BaseElementTemplate" />
	</xsl:template>

	<xsl:template name="PerformerTemplate">
		<xsl:call-template name="ResourceRoleTemplate" />
	</xsl:template>

	<!-- v2 -->
	<xsl:template name="ProcessTemplate">
		<!-- Call the CallableElementTemplate first -->
		<xsl:call-template name="CallableElementTemplate" />
		
		<!-- Add all attributes before any child elements -->
		<xsl:if test="@processType">
			<xsl:attribute name="processType">
				<xsl:value-of select="@processType" />
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="@isClosed">
			<xsl:attribute name="isClosed">
				<xsl:value-of select="@isClosed" />
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="@isExecutable">
			<xsl:attribute name="isExecutable">
				<xsl:value-of select="@isExecutable" />
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:supports">
			<xsl:attribute name="supports">
				<xsl:call-template name="concat">
					<xsl:with-param name="nodeset" select="bpmn:supports" />
				</xsl:call-template>
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="@definitionalCollaborationRef">
			<xsl:attribute name="definitionalCollaborationRef">
				<xsl:call-template name="concat">
					<xsl:with-param name="nodeset" select="@definitionalCollaborationRef" />
				</xsl:call-template>
			</xsl:attribute>
		</xsl:if>

		<!-- Now call the FlowElementsContainerTemplate -->
		<xsl:call-template name="FlowElementsContainerTemplate" />

		<!-- Add child elements -->
		<xsl:for-each select="bpmn:auditing">
			<auditing xmi:type="bpmnxmi:Auditing">
				<xsl:call-template name="AuditingTemplate" />
			</auditing>
		</xsl:for-each>
		<xsl:for-each select="bpmn:monitoring">
			<monitoring xmi:type="bpmnxmi:Monitoring">
				<xsl:call-template name="MonitoringTemplate" />
			</monitoring>
		</xsl:for-each>
		<xsl:for-each select="bpmn:property">
			<properties xmi:type="bpmnxmi:Property">
				<xsl:call-template name="PropertyTemplate" />
			</properties>
		</xsl:for-each>
		<xsl:for-each select="bpmn:performer">
			<resources xmi:type="bpmnxmi:Performer">
				<xsl:call-template name="PerformerTemplate" />
			</resources>
		</xsl:for-each>
		<xsl:for-each select="bpmn:humanPerformer">
			<resources xmi:type="bpmnxmi:HumanPerformer">
				<xsl:call-template name="HumanPerformerTemplate" />
			</resources>
		</xsl:for-each>
		<xsl:for-each select="bpmn:potentialOwner">
			<resources xmi:type="bpmnxmi:PotentialOwner">
				<xsl:call-template name="PotentialOwnerTemplate" />
			</resources>
		</xsl:for-each>
		<xsl:for-each select="bpmn:association">
			<artifacts xmi:type="bpmnxmi:Association">
				<xsl:call-template name="AssociationTemplate" />
			</artifacts>
		</xsl:for-each>
		<xsl:for-each select="bpmn:group">
			<artifacts xmi:type="bpmnxmi:Group">
				<xsl:call-template name="GroupTemplate" />
			</artifacts>
		</xsl:for-each>
		<xsl:for-each select="bpmn:textAnnotation">
			<artifacts xmi:type="bpmnxmi:TextAnnotation">
				<xsl:call-template name="TextAnnotationTemplate" />
			</artifacts>
		</xsl:for-each>
		<xsl:for-each select="bpmn:correlationSubscription">
			<correlationSubscriptions xmi:type="bpmnxmi:CorrelationSubscription">
				<xsl:call-template name="CorrelationSubscriptionTemplate" />
			</correlationSubscriptions>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="FlowElementsContainerTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:for-each select="bpmn:adHocSubProcess">
			<flowElements xmi:type="bpmnxmi:AdHocSubProcess">
				<xsl:call-template name="AdHocSubProcessTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:boundaryEvent">
			<flowElements xmi:type="bpmnxmi:BoundaryEvent">
				<xsl:call-template name="BoundaryEventTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:businessRuleTask">
			<flowElements xmi:type="bpmnxmi:BusinessRuleTask">
				<xsl:call-template name="BusinessRuleTaskTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:callActivity">
			<flowElements xmi:type="bpmnxmi:CallActivity">
				<xsl:call-template name="CallActivityTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:callChoreography">
			<flowElements xmi:type="bpmnxmi:CallChoreography">
				<xsl:call-template name="CallChoreographyTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:choreographyTask">
			<flowElements xmi:type="bpmnxmi:ChoreographyTask">
				<xsl:call-template name="ChoreographyTaskTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:complexGateway">
			<flowElements xmi:type="bpmnxmi:ComplexGateway">
				<xsl:call-template name="ComplexGatewayTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:dataObject">
			<flowElements xmi:type="bpmnxmi:DataObject">
				<xsl:call-template name="DataObjectTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:dataObjectReference">
			<flowElements xmi:type="bpmnxmi:DataObjectReference">
				<xsl:call-template name="DataObjectReferenceTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:dataStoreReference">
			<flowElements xmi:type="bpmnxmi:DataStoreReference">
				<xsl:call-template name="DataStoreReferenceTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:endEvent">
			<flowElements xmi:type="bpmnxmi:EndEvent">
				<xsl:call-template name="EndEventTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:event">
			<flowElements xmi:type="bpmnxmi:Event">
				<xsl:call-template name="EventTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:eventBasedGateway">
			<flowElements xmi:type="bpmnxmi:EventBasedGateway">
				<xsl:call-template name="EventBasedGatewayTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:exclusiveGateway">
			<flowElements xmi:type="bpmnxmi:ExclusiveGateway">
				<xsl:call-template name="ExclusiveGatewayTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:implicitThrowEvent">
			<flowElements xmi:type="bpmnxmi:ImplicitThrowEvent">
				<xsl:call-template name="ImplicitThrowEventTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:inclusiveGateway">
			<flowElements xmi:type="bpmnxmi:InclusiveGateway">
				<xsl:call-template name="InclusiveGatewayTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:intermediateCatchEvent">
			<flowElements xmi:type="bpmnxmi:IntermediateCatchEvent">
				<xsl:call-template name="IntermediateCatchEventTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:intermediateThrowEvent">
			<flowElements xmi:type="bpmnxmi:IntermediateThrowEvent">
				<xsl:call-template name="IntermediateThrowEventTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:manualTask">
			<flowElements xmi:type="bpmnxmi:ManualTask">
				<xsl:call-template name="ManualTaskTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:parallelGateway">
			<flowElements xmi:type="bpmnxmi:ParallelGateway">
				<xsl:call-template name="ParallelGatewayTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:receiveTask">
			<flowElements xmi:type="bpmnxmi:ReceiveTask">
				<xsl:call-template name="ReceiveTaskTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:scriptTask">
			<flowElements xmi:type="bpmnxmi:ScriptTask">
				<xsl:call-template name="ScriptTaskTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:sendTask">
			<flowElements xmi:type="bpmnxmi:SendTask">
				<xsl:call-template name="SendTaskTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:sequenceFlow">
			<flowElements xmi:type="bpmnxmi:SequenceFlow">
				<xsl:call-template name="SequenceFlowTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:serviceTask">
			<flowElements xmi:type="bpmnxmi:ServiceTask">
				<xsl:call-template name="ServiceTaskTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:startEvent">
			<flowElements xmi:type="bpmnxmi:StartEvent">
				<xsl:call-template name="StartEventTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:subChoreography">
			<flowElements xmi:type="bpmnxmi:SubChoreography">
				<xsl:call-template name="SubChoreographyTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:subProcess">
			<flowElements xmi:type="bpmnxmi:SubProcess">
				<xsl:call-template name="SubProcessTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:task">
			<flowElements xmi:type="bpmnxmi:Task">
				<xsl:call-template name="TaskTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:transaction">
			<flowElements xmi:type="bpmnxmi:Transaction">
				<xsl:call-template name="TransactionTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:userTask">
			<flowElements xmi:type="bpmnxmi:UserTask">
				<xsl:call-template name="UserTaskTemplate" />
			</flowElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:laneSet">
			<laneSets xmi:type="bpmnxmi:LaneSet">
				<xsl:call-template name="LaneSetTemplate" />
			</laneSets>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="FlowElementTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:categoryValueRef">
			<xsl:attribute name="categoryValueRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:categoryValueRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:auditing">
			<auditing xmi:type="bpmnxmi:Auditing">
				<xsl:call-template name="AuditingTemplate" />
			</auditing>
		</xsl:for-each>
		<xsl:for-each select="bpmn:monitoring">
			<monitoring xmi:type="bpmnxmi:Monitoring">
				<xsl:call-template name="MonitoringTemplate" />
			</monitoring>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="CategoryValueTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@value">
			<xsl:attribute name="value"> <xsl:value-of select="@value" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:categorizedFlowElements">
			<xsl:attribute name="categorizedFlowElements"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset"
				select="bpmn:categorizedFlowElements" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="LaneSetTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:lane">
			<lanes xmi:type="bpmnxmi:Lane">
				<xsl:call-template name="LaneTemplate" />
			</lanes>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="LaneTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@partitionElementRef">
			<xsl:attribute name="partitionElementRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@partitionElementRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:flowNodeRef">
			<xsl:attribute name="flowNodeRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:flowNodeRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:childLaneSet">
			<childLaneSet xmi:type="bpmnxmi:LaneSet">
				<xsl:call-template name="LaneSetTemplate" />
			</childLaneSet>
		</xsl:for-each>
		<xsl:for-each select="bpmn:partitionElement">
			<partitionElement xmi:type="bpmnxmi:BaseElement">
				<xsl:call-template name="BaseElementTemplate" />
			</partitionElement>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="FlowNodeTemplate">
		<xsl:call-template name="FlowElementTemplate" />
		<xsl:if test="bpmn:outgoing">
			<xsl:attribute name="outgoing"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:outgoing" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:incoming">
			<xsl:attribute name="incoming"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:incoming" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:lanes">
			<xsl:attribute name="lanes"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:lanes" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="SequenceFlowTemplate">
		<xsl:call-template name="FlowElementTemplate" />
		<xsl:if test="@isImmediate">
			<xsl:attribute name="isImmediate"> <xsl:value-of select="@isImmediate" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@targetRef">
			<xsl:attribute name="targetRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@targetRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@sourceRef">
			<xsl:attribute name="sourceRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@sourceRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:conditionExpression">
			<conditionExpression xmi:type="bpmnxmi:Expression">
				<xsl:call-template name="ExpressionTemplate" />
			</conditionExpression>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="PropertyTemplate">
		<xsl:call-template name="ItemAwareElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="CollaborationTemplate">
		<xsl:call-template name="RootElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@isClosed">
			<xsl:attribute name="isClosed"> <xsl:value-of select="@isClosed" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:choreographyRef">
			<xsl:attribute name="choreographyRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:choreographyRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:association">
			<artifacts xmi:type="bpmnxmi:Association">
				<xsl:call-template name="AssociationTemplate" />
			</artifacts>
		</xsl:for-each>
		<xsl:for-each select="bpmn:group">
			<artifacts xmi:type="bpmnxmi:Group">
				<xsl:call-template name="GroupTemplate" />
			</artifacts>
		</xsl:for-each>
		<xsl:for-each select="bpmn:textAnnotation">
			<artifacts xmi:type="bpmnxmi:TextAnnotation">
				<xsl:call-template name="TextAnnotationTemplate" />
			</artifacts>
		</xsl:for-each>
		<xsl:for-each select="bpmn:participantAssociation">
			<participantAssociations xmi:type="bpmnxmi:ParticipantAssociation">
				<xsl:call-template name="ParticipantAssociationTemplate" />
			</participantAssociations>
		</xsl:for-each>
		<xsl:for-each select="bpmn:messageFlowAssociation">
			<messageFlowAssociations xmi:type="bpmnxmi:MessageFlowAssociation">
				<xsl:call-template name="MessageFlowAssociationTemplate" />
			</messageFlowAssociations>
		</xsl:for-each>
		<xsl:for-each select="bpmn:conversationAssociation">
			<conversationAssociations xmi:type="bpmnxmi:ConversationAssociation">
				<xsl:call-template name="ConversationAssociationTemplate" />
			</conversationAssociations>
		</xsl:for-each>
		<xsl:for-each select="bpmn:participant">
			<participants xmi:type="bpmnxmi:Participant">
				<xsl:call-template name="ParticipantTemplate" />
			</participants>
		</xsl:for-each>
		<xsl:for-each select="bpmn:messageFlow">
			<messageFlows xmi:type="bpmnxmi:MessageFlow">
				<xsl:call-template name="MessageFlowTemplate" />
			</messageFlows>
		</xsl:for-each>
		<xsl:for-each select="bpmn:correlationKey">
			<correlationKeys xmi:type="bpmnxmi:CorrelationKey">
				<xsl:call-template name="CorrelationKeyTemplate" />
			</correlationKeys>
		</xsl:for-each>
		<xsl:for-each select="bpmn:callConversation">
			<conversations xmi:type="bpmnxmi:CallConversation">
				<xsl:call-template name="CallConversationTemplate" />
			</conversations>
		</xsl:for-each>
		<xsl:for-each select="bpmn:conversation">
			<conversations xmi:type="bpmnxmi:Conversation">
				<xsl:call-template name="ConversationTemplate" />
			</conversations>
		</xsl:for-each>
		<xsl:for-each select="bpmn:subConversation">
			<conversations xmi:type="bpmnxmi:SubConversation">
				<xsl:call-template name="SubConversationTemplate" />
			</conversations>
		</xsl:for-each>
		<xsl:for-each select="bpmn:conversationLink">
			<conversationLinks xmi:type="bpmnxmi:ConversationLink">
				<xsl:call-template name="ConversationLinkTemplate" />
			</conversationLinks>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="ChoreographyTemplate">
		<xsl:call-template name="CollaborationTemplate" />
		<xsl:call-template name="FlowElementsContainerTemplate" />
	</xsl:template>

	<xsl:template name="ArtifactTemplate">
		<xsl:call-template name="BaseElementTemplate" />
	</xsl:template>

	<xsl:template name="ParticipantAssociationTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="bpmn:innerParticipantRef">
			<xsl:attribute name="innerParticipantRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:innerParticipantRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:outerParticipantRef">
			<xsl:attribute name="outerParticipantRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:outerParticipantRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ParticipantTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:call-template name="InteractionNodeTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:interfaceRef">
			<xsl:attribute name="interfaceRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:interfaceRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:endPointRef">
			<xsl:attribute name="endPointRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:endPointRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@processRef">
			<xsl:attribute name="processRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@processRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:participantMultiplicity">
			<participantMultiplicity xmi:type="bpmnxmi:ParticipantMultiplicity">
				<xsl:call-template name="ParticipantMultiplicityTemplate" />
			</participantMultiplicity>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="InteractionNodeTemplate">
		<xsl:if test="bpmn:incomingConversationLinks">
			<xsl:attribute name="incomingConversationLinks"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset"
				select="bpmn:incomingConversationLinks" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:outgoingConversationLinks">
			<xsl:attribute name="outgoingConversationLinks"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset"
				select="bpmn:outgoingConversationLinks" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ConversationLinkTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@sourceRef">
			<xsl:attribute name="sourceRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@sourceRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@targetRef">
			<xsl:attribute name="targetRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@targetRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ParticipantMultiplicityTemplate">
		<xsl:if test="@minimum">
			<xsl:attribute name="minimum"> <xsl:value-of select="@minimum" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@maximum">
			<xsl:attribute name="maximum"> <xsl:value-of select="@maximum" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="MessageFlowAssociationTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@innerMessageFlowRef">
			<xsl:attribute name="innerMessageFlowRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@innerMessageFlowRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@outerMessageFlowRef">
			<xsl:attribute name="outerMessageFlowRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@outerMessageFlowRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="MessageFlowTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@sourceRef">
			<xsl:attribute name="sourceRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@sourceRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@targetRef">
			<xsl:attribute name="targetRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@targetRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@messageRef">
			<xsl:attribute name="messageRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@messageRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ConversationAssociationTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@innerConversationNodeRef">
			<xsl:attribute name="innerConversationNodeRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@innerConversationNodeRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@outerConversationNodeRef">
			<xsl:attribute name="outerConversationNodeRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@outerConversationNodeRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ConversationNodeTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:call-template name="InteractionNodeTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:participantRef">
			<xsl:attribute name="participantRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:participantRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:messageFlowRef">
			<xsl:attribute name="messageFlowRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:messageFlowRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:correlationKey">
			<correlationKeys xmi:type="bpmnxmi:CorrelationKey">
				<xsl:call-template name="CorrelationKeyTemplate" />
			</correlationKeys>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="CorrelationKeyTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:correlationPropertyRef">
			<xsl:attribute name="correlationPropertyRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset"
				select="bpmn:correlationPropertyRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="CorrelationPropertyTemplate">
		<xsl:call-template name="RootElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@type">
			<xsl:attribute name="type"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@type" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:correlationPropertyRetrievalExpression">
			<correlationPropertyRetrievalExpression
				xmi:type="bpmnxmi:CorrelationPropertyRetrievalExpression">
				<xsl:call-template name="CorrelationPropertyRetrievalExpressionTemplate" />
			</correlationPropertyRetrievalExpression>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="CorrelationPropertyRetrievalExpressionTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@messageRef">
			<xsl:attribute name="messageRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@messageRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:messagePath">
			<messagePath xmi:type="bpmnxmi:FormalExpression">
				<xsl:call-template name="FormalExpressionTemplate" />
			</messagePath>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="FormalExpressionTemplate">
		<xsl:call-template name="ExpressionTemplate" />
		<xsl:if test="@language">
			<xsl:attribute name="language"> <xsl:value-of select="@language" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:body">
			<xsl:attribute name="body"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:body" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@evaluatesToTypeRef">
			<xsl:attribute name="evaluatesToTypeRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@evaluatesToTypeRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="CorrelationSubscriptionTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@correlationKeyRef">
			<xsl:attribute name="correlationKeyRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@correlationKeyRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:correlationPropertyBinding">
			<correlationPropertyBinding xmi:type="bpmnxmi:CorrelationPropertyBinding">
				<xsl:call-template name="CorrelationPropertyBindingTemplate" />
			</correlationPropertyBinding>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="CorrelationPropertyBindingTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@correlationPropertyRef">
			<xsl:attribute name="correlationPropertyRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@correlationPropertyRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:dataPath">
			<dataPath xmi:type="bpmnxmi:FormalExpression">
				<xsl:call-template name="FormalExpressionTemplate" />
			</dataPath>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="GlobalManualTaskTemplate">
		<xsl:call-template name="GlobalTaskTemplate" />
	</xsl:template>

	<xsl:template name="ManualTaskTemplate">
		<xsl:call-template name="TaskTemplate" />
	</xsl:template>

	<xsl:template name="TaskTemplate">
		<xsl:call-template name="ActivityTemplate" />
		<xsl:call-template name="InteractionNodeTemplate" />
	</xsl:template>

	<xsl:template name="ActivityTemplate">
		<xsl:call-template name="FlowNodeTemplate" />
		<xsl:if test="@isForCompensation">
			<xsl:attribute name="isForCompensation"> <xsl:value-of select="@isForCompensation" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@startQuantity">
			<xsl:attribute name="startQuantity"> <xsl:value-of select="@startQuantity" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@completionQuantity">
			<xsl:attribute name="completionQuantity"> <xsl:value-of
				select="@completionQuantity" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@default">
			<xsl:attribute name="default"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@default" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:boundaryEventRefs">
			<xsl:attribute name="boundaryEventRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:boundaryEventRefs" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:multiInstanceLoopCharacteristics">
			<loopCharacteristics xmi:type="bpmnxmi:MultiInstanceLoopCharacteristics">
				<xsl:call-template name="MultiInstanceLoopCharacteristicsTemplate" />
			</loopCharacteristics>
		</xsl:for-each>
		<xsl:for-each select="bpmn:standardLoopCharacteristics">
			<loopCharacteristics xmi:type="bpmnxmi:StandardLoopCharacteristics">
				<xsl:call-template name="StandardLoopCharacteristicsTemplate" />
			</loopCharacteristics>
		</xsl:for-each>
		<xsl:for-each select="bpmn:performer">
			<resources xmi:type="bpmnxmi:Performer">
				<xsl:call-template name="PerformerTemplate" />
			</resources>
		</xsl:for-each>
		<xsl:for-each select="bpmn:humanPerformer">
			<resources xmi:type="bpmnxmi:HumanPerformer">
				<xsl:call-template name="HumanPerformerTemplate" />
			</resources>
		</xsl:for-each>
		<xsl:for-each select="bpmn:potentialOwner">
			<resources xmi:type="bpmnxmi:PotentialOwner">
				<xsl:call-template name="PotentialOwnerTemplate" />
			</resources>
		</xsl:for-each>
		<xsl:for-each select="bpmn:property">
			<properties xmi:type="bpmnxmi:Property">
				<xsl:call-template name="PropertyTemplate" />
			</properties>
		</xsl:for-each>
		<xsl:for-each select="bpmn:ioSpecification">
			<ioSpecification xmi:type="bpmnxmi:InputOutputSpecification">
				<xsl:call-template name="InputOutputSpecificationTemplate" />
			</ioSpecification>
		</xsl:for-each>
		<xsl:for-each select="bpmn:dataInputAssociation">
			<dataInputAssociations xmi:type="bpmnxmi:DataInputAssociation">
				<xsl:call-template name="DataInputAssociationTemplate" />
			</dataInputAssociations>
		</xsl:for-each>
		<xsl:for-each select="bpmn:dataOutputAssociation">
			<dataOutputAssociations xmi:type="bpmnxmi:DataOutputAssociation">
				<xsl:call-template name="DataOutputAssociationTemplate" />
			</dataOutputAssociations>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="LoopCharacteristicsTemplate">
		<xsl:call-template name="BaseElementTemplate" />
	</xsl:template>

	<!-- v2 -->
	<xsl:template name="BoundaryEventTemplate">

		<!-- Add all attributes before any child elements -->
		<xsl:if test="@cancelActivity">
			<xsl:attribute name="cancelActivity">
				<xsl:value-of select="@cancelActivity" />
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="@attachedToRef">
			<xsl:attribute name="attachedToRef">
				<xsl:call-template name="concat">
					<xsl:with-param name="nodeset" select="@attachedToRef" />
				</xsl:call-template>
			</xsl:attribute>
		</xsl:if>

		<!-- Now call the CatchEventTemplate -->
		<xsl:call-template name="CatchEventTemplate" />
	</xsl:template>

	<xsl:template name="CatchEventTemplate">
		<xsl:call-template name="EventTemplate" />
		<xsl:if test="@parallelMultiple">
			<xsl:attribute name="parallelMultiple"> <xsl:value-of select="@parallelMultiple" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:eventDefinitionRef">
			<xsl:attribute name="eventDefinitionRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:eventDefinitionRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:outputSet">
			<outputSet xmi:type="bpmnxmi:OutputSet">
				<xsl:call-template name="OutputSetTemplate" />
			</outputSet>
		</xsl:for-each>
		<xsl:for-each select="bpmn:dataOutputAssociation">
			<dataOutputAssociation xmi:type="bpmnxmi:DataOutputAssociation">
				<xsl:call-template name="DataOutputAssociationTemplate" />
			</dataOutputAssociation>
		</xsl:for-each>
		<xsl:for-each select="bpmn:dataOutput">
			<dataOutputs xmi:type="bpmnxmi:DataOutput">
				<xsl:call-template name="DataOutputTemplate" />
			</dataOutputs>
		</xsl:for-each>
		<xsl:for-each select="bpmn:cancelEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:CancelEventDefinition">
				<xsl:call-template name="CancelEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:compensateEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:CompensateEventDefinition">
				<xsl:call-template name="CompensateEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:conditionalEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:ConditionalEventDefinition">
				<xsl:call-template name="ConditionalEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:errorEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:ErrorEventDefinition">
				<xsl:call-template name="ErrorEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:escalationEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:EscalationEventDefinition">
				<xsl:call-template name="EscalationEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:linkEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:LinkEventDefinition">
				<xsl:call-template name="LinkEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:messageEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:MessageEventDefinition">
				<xsl:call-template name="MessageEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:signalEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:SignalEventDefinition">
				<xsl:call-template name="SignalEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:terminateEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:TerminateEventDefinition">
				<xsl:call-template name="TerminateEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:timerEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:TimerEventDefinition">
				<xsl:call-template name="TimerEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="EventTemplate">
		<xsl:call-template name="FlowNodeTemplate" />
		<xsl:call-template name="InteractionNodeTemplate" />
		<xsl:for-each select="bpmn:property">
			<properties xmi:type="bpmnxmi:Property">
				<xsl:call-template name="PropertyTemplate" />
			</properties>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="EventDefinitionTemplate">
		<xsl:call-template name="RootElementTemplate" />
	</xsl:template>

	<xsl:template name="DataOutputAssociationTemplate">
		<xsl:call-template name="DataAssociationTemplate" />
	</xsl:template>

	<xsl:template name="DataAssociationTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="bpmn:targetRef">
			<xsl:attribute name="targetRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:targetRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:sourceRef">
			<xsl:attribute name="sourceRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:sourceRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:transformation">
			<transformation xmi:type="bpmnxmi:FormalExpression">
				<xsl:call-template name="FormalExpressionTemplate" />
			</transformation>
		</xsl:for-each>
		<xsl:for-each select="bpmn:assignment">
			<assignment xmi:type="bpmnxmi:Assignment">
				<xsl:call-template name="AssignmentTemplate" />
			</assignment>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="AssignmentTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:for-each select="bpmn:from">
			<from xmi:type="bpmnxmi:Expression">
				<xsl:call-template name="ExpressionTemplate" />
			</from>
		</xsl:for-each>
		<xsl:for-each select="bpmn:to">
			<to xmi:type="bpmnxmi:Expression">
				<xsl:call-template name="ExpressionTemplate" />
			</to>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="DataInputAssociationTemplate">
		<xsl:call-template name="DataAssociationTemplate" />
	</xsl:template>

	<xsl:template name="UserTaskTemplate">
		<xsl:call-template name="TaskTemplate" />
		<xsl:if test="@implementation">
			<xsl:attribute name="implementation"> <xsl:value-of select="@implementation" /> </xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:rendering">
			<renderings xmi:type="bpmnxmi:Rendering">
				<xsl:call-template name="RenderingTemplate" />
			</renderings>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="RenderingTemplate">
		<xsl:call-template name="BaseElementTemplate" />
	</xsl:template>

	<xsl:template name="HumanPerformerTemplate">
		<xsl:call-template name="PerformerTemplate" />
	</xsl:template>

	<xsl:template name="PotentialOwnerTemplate">
		<xsl:call-template name="HumanPerformerTemplate" />
	</xsl:template>

	<xsl:template name="GlobalUserTaskTemplate">
		<xsl:call-template name="GlobalTaskTemplate" />
		<xsl:if test="@implementation">
			<xsl:attribute name="implementation"> <xsl:value-of select="@implementation" /> </xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:rendering">
			<renderings xmi:type="bpmnxmi:Rendering">
				<xsl:call-template name="RenderingTemplate" />
			</renderings>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="GatewayTemplate">
		<xsl:call-template name="FlowNodeTemplate" />
		<xsl:if test="@gatewayDirection">
			<xsl:attribute name="gatewayDirection"> <xsl:value-of select="@gatewayDirection" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="EventBasedGatewayTemplate">
		<xsl:call-template name="GatewayTemplate" />
		<xsl:if test="@instantiate">
			<xsl:attribute name="instantiate"> <xsl:value-of select="@instantiate" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@eventGatewayType">
			<xsl:attribute name="eventGatewayType"> <xsl:value-of select="@eventGatewayType" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ComplexGatewayTemplate">
		<xsl:call-template name="GatewayTemplate" />
		<xsl:if test="@default">
			<xsl:attribute name="default"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@default" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:activationCondition">
			<activationCondition xmi:type="bpmnxmi:Expression">
				<xsl:call-template name="ExpressionTemplate" />
			</activationCondition>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="ExclusiveGatewayTemplate">
		<xsl:call-template name="GatewayTemplate" />
		<xsl:if test="@default">
			<xsl:attribute name="default"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@default" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="InclusiveGatewayTemplate">
		<xsl:call-template name="GatewayTemplate" />
		<xsl:if test="@default">
			<xsl:attribute name="default"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@default" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ParallelGatewayTemplate">
		<xsl:call-template name="GatewayTemplate" />
	</xsl:template>

	<xsl:template name="RelationshipTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@type">
			<xsl:attribute name="type"> <xsl:value-of select="@type" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@direction">
			<xsl:attribute name="direction"> <xsl:value-of select="@direction" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:source">
			<xsl:attribute name="sources"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:source" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:target">
			<xsl:attribute name="targets"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:target" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ExtensionTemplate">
		<xsl:if test="@mustUnderstand">
			<xsl:attribute name="mustUnderstand"> <xsl:value-of select="@mustUnderstand" /> </xsl:attribute>
		</xsl:if>
		<xsl:for-each select="@definition">
			<definition xmi:type="bpmnxmi:ExtensionDefinition">
				<xsl:call-template name="ExtensionDefinitionTemplate" />
			</definition>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="IntermediateCatchEventTemplate">
		<xsl:call-template name="CatchEventTemplate" />
	</xsl:template>

	<xsl:template name="IntermediateThrowEventTemplate">
		<xsl:call-template name="ThrowEventTemplate" />
	</xsl:template>

	<xsl:template name="ThrowEventTemplate">
		<xsl:call-template name="EventTemplate" />
		<xsl:if test="bpmn:eventDefinitionRef">
			<xsl:attribute name="eventDefinitionRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:eventDefinitionRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:inputSet">
			<inputSet xmi:type="bpmnxmi:InputSet">
				<xsl:call-template name="InputSetTemplate" />
			</inputSet>
		</xsl:for-each>
		<xsl:for-each select="bpmn:dataInputAssociation">
			<dataInputAssociation xmi:type="bpmnxmi:DataInputAssociation">
				<xsl:call-template name="DataInputAssociationTemplate" />
			</dataInputAssociation>
		</xsl:for-each>
		<xsl:for-each select="bpmn:dataInput">
			<dataInputs xmi:type="bpmnxmi:DataInput">
				<xsl:call-template name="DataInputTemplate" />
			</dataInputs>
		</xsl:for-each>
		<xsl:for-each select="bpmn:cancelEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:CancelEventDefinition">
				<xsl:call-template name="CancelEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:compensateEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:CompensateEventDefinition">
				<xsl:call-template name="CompensateEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:conditionalEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:ConditionalEventDefinition">
				<xsl:call-template name="ConditionalEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:errorEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:ErrorEventDefinition">
				<xsl:call-template name="ErrorEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:escalationEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:EscalationEventDefinition">
				<xsl:call-template name="EscalationEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:linkEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:LinkEventDefinition">
				<xsl:call-template name="LinkEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:messageEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:MessageEventDefinition">
				<xsl:call-template name="MessageEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:signalEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:SignalEventDefinition">
				<xsl:call-template name="SignalEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:terminateEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:TerminateEventDefinition">
				<xsl:call-template name="TerminateEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:timerEventDefinition">
			<eventDefinitions xmi:type="bpmnxmi:TimerEventDefinition">
				<xsl:call-template name="TimerEventDefinitionTemplate" />
			</eventDefinitions>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="EndEventTemplate">
		<xsl:call-template name="ThrowEventTemplate" />
	</xsl:template>

	<xsl:template name="StartEventTemplate">
		<xsl:call-template name="CatchEventTemplate" />
		<xsl:if test="@isInterrupting">
			<xsl:attribute name="isInterrupting"> <xsl:value-of select="@isInterrupting" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="CancelEventDefinitionTemplate">
		<xsl:call-template name="EventDefinitionTemplate" />
	</xsl:template>

	<xsl:template name="ErrorEventDefinitionTemplate">
		<xsl:call-template name="EventDefinitionTemplate" />
		<xsl:if test="@errorRef">
			<xsl:attribute name="errorRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@errorRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="TerminateEventDefinitionTemplate">
		<xsl:call-template name="EventDefinitionTemplate" />
	</xsl:template>

	<xsl:template name="EscalationEventDefinitionTemplate">
		<xsl:call-template name="EventDefinitionTemplate" />
		<xsl:if test="@escalationRef">
			<xsl:attribute name="escalationRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@escalationRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="EscalationTemplate">
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@escalationCode">
			<xsl:attribute name="escalationCode"> <xsl:value-of select="@escalationCode" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@structureRef">
			<xsl:attribute name="structureRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@structureRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="CompensateEventDefinitionTemplate">
		<xsl:call-template name="EventDefinitionTemplate" />
		<xsl:if test="@waitForCompletion">
			<xsl:attribute name="waitForCompletion"> <xsl:value-of select="@waitForCompletion" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@activityRef">
			<xsl:attribute name="activityRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@activityRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="TimerEventDefinitionTemplate">
		<xsl:call-template name="EventDefinitionTemplate" />
		<xsl:for-each select="bpmn:timeDate">
			<timeDate xmi:type="bpmnxmi:Expression">
				<xsl:call-template name="ExpressionTemplate" />
			</timeDate>
		</xsl:for-each>
		<xsl:for-each select="bpmn:timeCycle">
			<timeCycle xmi:type="bpmnxmi:Expression">
				<xsl:call-template name="ExpressionTemplate" />
			</timeCycle>
		</xsl:for-each>
		<xsl:for-each select="bpmn:timeDuration">
			<timeDuration xmi:type="bpmnxmi:Expression">
				<xsl:call-template name="ExpressionTemplate" />
			</timeDuration>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="LinkEventDefinitionTemplate">
		<xsl:call-template name="EventDefinitionTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:target">
			<xsl:attribute name="target"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:target" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:source">
			<xsl:attribute name="source"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:source" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="MessageEventDefinitionTemplate">
		<xsl:call-template name="EventDefinitionTemplate" />
		<xsl:if test="@messageRef">
			<xsl:attribute name="messageRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@messageRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:operationRef">
			<xsl:attribute name="operationRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:operationRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ConditionalEventDefinitionTemplate">
		<xsl:call-template name="EventDefinitionTemplate" />
		<xsl:for-each select="bpmn:condition">
			<condition xmi:type="bpmnxmi:Expression">
				<xsl:call-template name="ExpressionTemplate" />
			</condition>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="SignalEventDefinitionTemplate">
		<xsl:call-template name="EventDefinitionTemplate" />
		<xsl:if test="@signalRef">
			<xsl:attribute name="signalRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@signalRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="SignalTemplate">
		<xsl:call-template name="RootElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@structureRef">
			<xsl:attribute name="structureRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@structureRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ImplicitThrowEventTemplate">
		<xsl:call-template name="ThrowEventTemplate" />
	</xsl:template>

	<xsl:template name="DataObjectTemplate">
		<xsl:call-template name="FlowElementTemplate" />
		<xsl:call-template name="ItemAwareElementTemplate" />
		<xsl:if test="@isCollection">
			<xsl:attribute name="isCollection"> <xsl:value-of select="@isCollection" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="DataStoreTemplate">
		<xsl:call-template name="ItemAwareElementTemplate" />
		<xsl:call-template name="RootElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@capacity">
			<xsl:attribute name="capacity"> <xsl:value-of select="@capacity" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@isUnlimited">
			<xsl:attribute name="isUnlimited"> <xsl:value-of select="@isUnlimited" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="DataStoreReferenceTemplate">
		<xsl:call-template name="FlowElementTemplate" />
		<xsl:call-template name="ItemAwareElementTemplate" />
		<xsl:if test="@dataStoreRef">
			<xsl:attribute name="dataStoreRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@dataStoreRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="DataObjectReferenceTemplate">
		<xsl:call-template name="FlowElementTemplate" />
		<xsl:call-template name="ItemAwareElementTemplate" />
		<xsl:if test="@dataObjectRef">
			<xsl:attribute name="dataObjectRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@dataObjectRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="CallConversationTemplate">
		<xsl:call-template name="ConversationNodeTemplate" />
		<xsl:if test="@calledCollaborationRef">
			<xsl:attribute name="calledCollaborationRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@calledCollaborationRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:participantAssociation">
			<participantAssociations xmi:type="bpmnxmi:ParticipantAssociation">
				<xsl:call-template name="ParticipantAssociationTemplate" />
			</participantAssociations>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="ConversationTemplate">
		<xsl:call-template name="ConversationNodeTemplate" />
	</xsl:template>

	<xsl:template name="SubConversationTemplate">
		<xsl:call-template name="ConversationNodeTemplate" />
		<xsl:for-each select="bpmn:callConversation">
			<conversationNodes xmi:type="bpmnxmi:CallConversation">
				<xsl:call-template name="CallConversationTemplate" />
			</conversationNodes>
		</xsl:for-each>
		<xsl:for-each select="bpmn:conversation">
			<conversationNodes xmi:type="bpmnxmi:Conversation">
				<xsl:call-template name="ConversationTemplate" />
			</conversationNodes>
		</xsl:for-each>
		<xsl:for-each select="bpmn:subConversation">
			<conversationNodes xmi:type="bpmnxmi:SubConversation">
				<xsl:call-template name="SubConversationTemplate" />
			</conversationNodes>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="GlobalConversationTemplate">
		<xsl:call-template name="CollaborationTemplate" />
	</xsl:template>

	<xsl:template name="PartnerEntityTemplate">
		<xsl:call-template name="RootElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:participantRef">
			<xsl:attribute name="participantRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:participantRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="PartnerRoleTemplate">
		<xsl:call-template name="RootElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:participantRef">
			<xsl:attribute name="participantRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:participantRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ChoreographyActivityTemplate">
		<xsl:call-template name="FlowNodeTemplate" />
		<xsl:if test="@loopType">
			<xsl:attribute name="loopType"> <xsl:value-of select="@loopType" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:participantRef">
			<xsl:attribute name="participantRefs"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:participantRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@initiatingParticipantRef">
			<xsl:attribute name="initiatingParticipantRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@initiatingParticipantRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:correlationKey">
			<correlationKeys xmi:type="bpmnxmi:CorrelationKey">
				<xsl:call-template name="CorrelationKeyTemplate" />
			</correlationKeys>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="CallChoreographyTemplate">
		<xsl:call-template name="ChoreographyActivityTemplate" />
		<xsl:if test="@calledChoreographyRef">
			<xsl:attribute name="calledChoreographyRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@calledChoreographyRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:participantAssociation">
			<participantAssociations xmi:type="bpmnxmi:ParticipantAssociation">
				<xsl:call-template name="ParticipantAssociationTemplate" />
			</participantAssociations>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="SubChoreographyTemplate">
		<xsl:call-template name="ChoreographyActivityTemplate" />
		<xsl:call-template name="FlowElementsContainerTemplate" />
		<xsl:for-each select="bpmn:association">
			<artifacts xmi:type="bpmnxmi:Association">
				<xsl:call-template name="AssociationTemplate" />
			</artifacts>
		</xsl:for-each>
		<xsl:for-each select="bpmn:group">
			<artifacts xmi:type="bpmnxmi:Group">
				<xsl:call-template name="GroupTemplate" />
			</artifacts>
		</xsl:for-each>
		<xsl:for-each select="bpmn:textAnnotation">
			<artifacts xmi:type="bpmnxmi:TextAnnotation">
				<xsl:call-template name="TextAnnotationTemplate" />
			</artifacts>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="ChoreographyTaskTemplate">
		<xsl:call-template name="ChoreographyActivityTemplate" />
		<xsl:if test="bpmn:messageFlowRef">
			<xsl:attribute name="messageFlowRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:messageFlowRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="GlobalChoreographyTaskTemplate">
		<xsl:call-template name="ChoreographyTemplate" />
		<xsl:if test="@initiatingParticipantRef">
			<xsl:attribute name="initiatingParticipantRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@initiatingParticipantRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="TextAnnotationTemplate">
		<xsl:call-template name="ArtifactTemplate" />
		<xsl:if test="bpmn:text">
			<xsl:attribute name="text"> <xsl:value-of select="bpmn:text" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@textFormat">
			<xsl:attribute name="textFormat"> <xsl:value-of select="@textFormat" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="GroupTemplate">
		<xsl:call-template name="ArtifactTemplate" />
		<xsl:if test="@categoryValueRef">
			<xsl:attribute name="categoryValueRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@categoryValueRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="AssociationTemplate">
		<xsl:call-template name="ArtifactTemplate" />
		<xsl:if test="@associationDirection">
			<xsl:attribute name="associationDirection"> <xsl:value-of
				select="@associationDirection" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@sourceRef">
			<xsl:attribute name="sourceRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@sourceRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@targetRef">
			<xsl:attribute name="targetRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@targetRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="CategoryTemplate">
		<xsl:call-template name="RootElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:categoryValue">
			<categoryValue xmi:type="bpmnxmi:CategoryValue">
				<xsl:call-template name="CategoryValueTemplate" />
			</categoryValue>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="ServiceTaskTemplate">
		<xsl:call-template name="TaskTemplate" />
		<xsl:if test="@implementation">
			<xsl:attribute name="implementation"> <xsl:value-of select="@implementation" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@operationRef">
			<xsl:attribute name="operationRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@operationRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="SubProcessTemplate">
		<xsl:call-template name="ActivityTemplate" />
		<xsl:call-template name="FlowElementsContainerTemplate" />
		<xsl:if test="@triggeredByEvent">
			<xsl:attribute name="triggeredByEvent"> <xsl:value-of select="@triggeredByEvent" /> </xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:association">
			<artifacts xmi:type="bpmnxmi:Association">
				<xsl:call-template name="AssociationTemplate" />
			</artifacts>
		</xsl:for-each>
		<xsl:for-each select="bpmn:group">
			<artifacts xmi:type="bpmnxmi:Group">
				<xsl:call-template name="GroupTemplate" />
			</artifacts>
		</xsl:for-each>
		<xsl:for-each select="bpmn:textAnnotation">
			<artifacts xmi:type="bpmnxmi:TextAnnotation">
				<xsl:call-template name="TextAnnotationTemplate" />
			</artifacts>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="MultiInstanceLoopCharacteristicsTemplate">
		<xsl:call-template name="LoopCharacteristicsTemplate" />
		<xsl:if test="@isSequential">
			<xsl:attribute name="isSequential"> <xsl:value-of select="@isSequential" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@behavior">
			<xsl:attribute name="behavior"> <xsl:value-of select="@behavior" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:loopDataInputRef">
			<xsl:attribute name="loopDataInputRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:loopDataInputRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:loopDataOutputRef">
			<xsl:attribute name="loopDataOutputRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:loopDataOutputRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@oneBehaviorEventRef">
			<xsl:attribute name="oneBehaviorEventRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@oneBehaviorEventRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@noneBehaviorEventRef">
			<xsl:attribute name="noneBehaviorEventRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@noneBehaviorEventRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:loopCardinality">
			<loopCardinality xmi:type="bpmnxmi:Expression">
				<xsl:call-template name="ExpressionTemplate" />
			</loopCardinality>
		</xsl:for-each>
		<xsl:for-each select="bpmn:inputDataItem">
			<inputDataItem xmi:type="bpmnxmi:DataInput">
				<xsl:call-template name="DataInputTemplate" />
			</inputDataItem>
		</xsl:for-each>
		<xsl:for-each select="bpmn:outputDataItem">
			<outputDataItem xmi:type="bpmnxmi:DataOutput">
				<xsl:call-template name="DataOutputTemplate" />
			</outputDataItem>
		</xsl:for-each>
		<xsl:for-each select="bpmn:completionCondition">
			<completionCondition xmi:type="bpmnxmi:Expression">
				<xsl:call-template name="ExpressionTemplate" />
			</completionCondition>
		</xsl:for-each>
		<xsl:for-each select="bpmn:complexBehaviorDefinition">
			<complexBehaviorDefinition xmi:type="bpmnxmi:ComplexBehaviorDefinition">
				<xsl:call-template name="ComplexBehaviorDefinitionTemplate" />
			</complexBehaviorDefinition>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="ComplexBehaviorDefinitionTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:for-each select="bpmn:condition">
			<condition xmi:type="bpmnxmi:FormalExpression">
				<xsl:call-template name="FormalExpressionTemplate" />
			</condition>
		</xsl:for-each>
		<xsl:for-each select="bpmn:event">
			<event xmi:type="bpmnxmi:ImplicitThrowEvent">
				<xsl:call-template name="ImplicitThrowEventTemplate" />
			</event>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="StandardLoopCharacteristicsTemplate">
		<xsl:call-template name="LoopCharacteristicsTemplate" />
		<xsl:if test="@testBefore">
			<xsl:attribute name="testBefore"> <xsl:value-of select="@testBefore" /> </xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:loopCondition">
			<loopCondition xmi:type="bpmnxmi:Expression">
				<xsl:call-template name="ExpressionTemplate" />
			</loopCondition>
		</xsl:for-each>
		<xsl:for-each select="@loopMaximum">
			<loopMaximum xmi:type="bpmnxmi:Expression">
				<xsl:call-template name="ExpressionTemplate" />
			</loopMaximum>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="CallActivityTemplate">
		<xsl:call-template name="ActivityTemplate" />
		<xsl:if test="@calledElement">
			<xsl:attribute name="calledElementRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@calledElement" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="SendTaskTemplate">
		<xsl:call-template name="TaskTemplate" />
		<xsl:if test="@implementation">
			<xsl:attribute name="implementation"> <xsl:value-of select="@implementation" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@operationRef">
			<xsl:attribute name="operationRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@operationRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@messageRef">
			<xsl:attribute name="messageRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@messageRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ReceiveTaskTemplate">
		<xsl:call-template name="TaskTemplate" />
		<xsl:if test="@implementation">
			<xsl:attribute name="implementation"> <xsl:value-of select="@implementation" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@instantiate">
			<xsl:attribute name="instantiate"> <xsl:value-of select="@instantiate" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@operationRef">
			<xsl:attribute name="operationRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@operationRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="@messageRef">
			<xsl:attribute name="messageRef"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@messageRef" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="ScriptTaskTemplate">
		<xsl:call-template name="TaskTemplate" />
		<xsl:if test="@scriptFormat">
			<xsl:attribute name="scriptFormat"> <xsl:value-of select="@scriptFormat" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:script">
			<xsl:attribute name="script"> <xsl:value-of select="bpmn:script" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="BusinessRuleTaskTemplate">
		<xsl:call-template name="TaskTemplate" />
		<xsl:if test="@implementation">
			<xsl:attribute name="implementation"> <xsl:value-of select="@implementation" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="AdHocSubProcessTemplate">
		<xsl:call-template name="SubProcessTemplate" />
		<xsl:if test="@ordering">
			<xsl:attribute name="ordering"> <xsl:value-of select="@ordering" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@cancelRemainingInstances">
			<xsl:attribute name="cancelRemainingInstances"> <xsl:value-of
				select="@cancelRemainingInstances" /> </xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:completionCondition">
			<completionCondition xmi:type="bpmnxmi:Expression">
				<xsl:call-template name="ExpressionTemplate" />
			</completionCondition>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="TransactionTemplate">
		<xsl:call-template name="SubProcessTemplate" />
		<xsl:if test="bpmn:protocol">
			<xsl:attribute name="protocol"> <xsl:value-of select="bpmn:protocol" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@method">
			<xsl:attribute name="method"> <xsl:value-of select="@method" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="GlobalScriptTaskTemplate">
		<xsl:call-template name="GlobalTaskTemplate" />
		<xsl:if test="@scriptLanguage">
			<xsl:attribute name="scriptLanguage"> <xsl:value-of select="@scriptLanguage" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:script">
			<xsl:attribute name="script"> <xsl:value-of select="bpmn:script" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="GlobalBusinessRuleTaskTemplate">
		<xsl:call-template name="GlobalTaskTemplate" />
		<xsl:if test="@implementation">
			<xsl:attribute name="implementation"> <xsl:value-of select="@implementation" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="DefinitionsTemplate">
		<xsl:call-template name="BaseElementTemplate" />
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@targetNamespace">
			<xsl:attribute name="targetNamespace"> <xsl:value-of select="@targetNamespace" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@expressionLanguage">
			<xsl:attribute name="expressionLanguage"> <xsl:value-of
				select="@expressionLanguage" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@typeLanguage">
			<xsl:attribute name="typeLanguage"> <xsl:value-of select="@typeLanguage" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@exporter">
			<xsl:attribute name="exporter"> <xsl:value-of select="@exporter" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@exporterVersion">
			<xsl:attribute name="exporterVersion"> <xsl:value-of select="@exporterVersion" /> </xsl:attribute>
		</xsl:if>
		<xsl:for-each select="bpmn:import">
			<imports xmi:type="bpmnxmi:Import">
				<xsl:call-template name="ImportTemplate" />
			</imports>
		</xsl:for-each>
		<xsl:for-each select="bpmn:extension">
			<extensions xmi:type="bpmnxmi:Extension">
				<xsl:call-template name="ExtensionTemplate" />
			</extensions>
		</xsl:for-each>
		<xsl:for-each select="bpmn:relationship">
			<relationships xmi:type="bpmnxmi:Relationship">
				<xsl:call-template name="RelationshipTemplate" />
			</relationships>
		</xsl:for-each>
		<xsl:for-each select="bpmn:eventDefinition">
			<rootElements xmi:type="bpmnxmi:EventDefinition">
				<xsl:call-template name="EventDefinitionTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:category">
			<rootElements xmi:type="bpmnxmi:Category">
				<xsl:call-template name="CategoryTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:collaboration">
			<rootElements xmi:type="bpmnxmi:Collaboration">
				<xsl:call-template name="CollaborationTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:correlationProperty">
			<rootElements xmi:type="bpmnxmi:CorrelationProperty">
				<xsl:call-template name="CorrelationPropertyTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:dataStore">
			<rootElements xmi:type="bpmnxmi:DataStore">
				<xsl:call-template name="DataStoreTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:endPoint">
			<rootElements xmi:type="bpmnxmi:EndPoint">
				<xsl:call-template name="EndPointTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:error">
			<rootElements xmi:type="bpmnxmi:Error">
				<xsl:call-template name="ErrorTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:escalation">
			<rootElements xmi:type="bpmnxmi:Escalation">
				<xsl:call-template name="EscalationTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:globalBusinessRuleTask">
			<rootElements xmi:type="bpmnxmi:GlobalBusinessRuleTask">
				<xsl:call-template name="GlobalBusinessRuleTaskTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:globalManualTask">
			<rootElements xmi:type="bpmnxmi:GlobalManualTask">
				<xsl:call-template name="GlobalManualTaskTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:globalScriptTask">
			<rootElements xmi:type="bpmnxmi:GlobalScriptTask">
				<xsl:call-template name="GlobalScriptTaskTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:globalTask">
			<rootElements xmi:type="bpmnxmi:GlobalTask">
				<xsl:call-template name="GlobalTaskTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:globalUserTask">
			<rootElements xmi:type="bpmnxmi:GlobalUserTask">
				<xsl:call-template name="GlobalUserTaskTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:interface">
			<rootElements xmi:type="bpmnxmi:Interface">
				<xsl:call-template name="InterfaceTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:itemDefinition">
			<rootElements xmi:type="bpmnxmi:ItemDefinition">
				<xsl:call-template name="ItemDefinitionTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:message">
			<rootElements xmi:type="bpmnxmi:Message">
				<xsl:call-template name="MessageTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:partnerEntity">
			<rootElements xmi:type="bpmnxmi:PartnerEntity">
				<xsl:call-template name="PartnerEntityTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:partnerRole">
			<rootElements xmi:type="bpmnxmi:PartnerRole">
				<xsl:call-template name="PartnerRoleTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:process">
			<rootElements xmi:type="bpmnxmi:Process">
				<xsl:call-template name="ProcessTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:resource">
			<rootElements xmi:type="bpmnxmi:Resource">
				<xsl:call-template name="ResourceTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:signal">
			<rootElements xmi:type="bpmnxmi:Signal">
				<xsl:call-template name="SignalTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:cancelEventDefinition">
			<rootElements xmi:type="bpmnxmi:CancelEventDefinition">
				<xsl:call-template name="CancelEventDefinitionTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:compensateEventDefinition">
			<rootElements xmi:type="bpmnxmi:CompensateEventDefinition">
				<xsl:call-template name="CompensateEventDefinitionTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:conditionalEventDefinition">
			<rootElements xmi:type="bpmnxmi:ConditionalEventDefinition">
				<xsl:call-template name="ConditionalEventDefinitionTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:errorEventDefinition">
			<rootElements xmi:type="bpmnxmi:ErrorEventDefinition">
				<xsl:call-template name="ErrorEventDefinitionTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:escalationEventDefinition">
			<rootElements xmi:type="bpmnxmi:EscalationEventDefinition">
				<xsl:call-template name="EscalationEventDefinitionTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:linkEventDefinition">
			<rootElements xmi:type="bpmnxmi:LinkEventDefinition">
				<xsl:call-template name="LinkEventDefinitionTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:messageEventDefinition">
			<rootElements xmi:type="bpmnxmi:MessageEventDefinition">
				<xsl:call-template name="MessageEventDefinitionTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:signalEventDefinition">
			<rootElements xmi:type="bpmnxmi:SignalEventDefinition">
				<xsl:call-template name="SignalEventDefinitionTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:terminateEventDefinition">
			<rootElements xmi:type="bpmnxmi:TerminateEventDefinition">
				<xsl:call-template name="TerminateEventDefinitionTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:timerEventDefinition">
			<rootElements xmi:type="bpmnxmi:TimerEventDefinition">
				<xsl:call-template name="TimerEventDefinitionTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:choreography">
			<rootElements xmi:type="bpmnxmi:Choreography">
				<xsl:call-template name="ChoreographyTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:globalConversation">
			<rootElements xmi:type="bpmnxmi:GlobalConversation">
				<xsl:call-template name="GlobalConversationTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmn:globalChoreographyTask">
			<rootElements xmi:type="bpmnxmi:GlobalChoreographyTask">
				<xsl:call-template name="GlobalChoreographyTaskTemplate" />
			</rootElements>
		</xsl:for-each>
		<xsl:for-each select="bpmndi:BPMNDiagram">
			<diagrams xmi:type="bpmndixmi:BPMNDiagram">
				<xsl:call-template name="BPMNDiagramTemplate" />
			</diagrams>
		</xsl:for-each>
	</xsl:template>

	<!-- v2 -->
	<xsl:template name="BPMNPlaneTemplate">
		<!-- Add all attributes before any child elements -->
		<xsl:if test="@bpmnElement">
			<xsl:attribute name="bpmnElement">
				<xsl:call-template name="concat">
					<xsl:with-param name="nodeset" select="@bpmnElement" />
				</xsl:call-template>
			</xsl:attribute>
		</xsl:if>

		<!-- Now call the PlaneTemplate -->
		<xsl:call-template name="PlaneTemplate" />
	</xsl:template>

	<xsl:template name="BPMNShapeTemplate">
		<!-- Add all attributes before any child elements -->
		<xsl:if test="@isHorizontal">
			<xsl:attribute name="isHorizontal">
				<xsl:value-of select="@isHorizontal" />
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="@isExpanded">
			<xsl:attribute name="isExpanded">
				<xsl:value-of select="@isExpanded" />
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="@isMarkerVisible">
			<xsl:attribute name="isMarkerVisible">
				<xsl:value-of select="@isMarkerVisible" />
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="@isMessageVisible">
			<xsl:attribute name="isMessageVisible">
				<xsl:value-of select="@isMessageVisible" />
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="@participantBandKind">
			<xsl:attribute name="participantBandKind">
				<xsl:value-of select="@participantBandKind" />
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="@bpmnElement">
			<xsl:attribute name="bpmnElement">
				<xsl:call-template name="concat">
					<xsl:with-param name="nodeset" select="@bpmnElement" />
				</xsl:call-template>
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="@choreographyActivityShape">
			<xsl:attribute name="choreographyActivityShape">
				<xsl:call-template name="concat">
					<xsl:with-param name="nodeset" select="@choreographyActivityShape" />
				</xsl:call-template>
			</xsl:attribute>
		</xsl:if>

		<!-- Now call the LabeledShapeTemplate -->
		<xsl:call-template name="LabeledShapeTemplate" />

		<!-- Add child elements -->
		<xsl:for-each select="bpmndi:BPMNLabel">
			<label xmi:type="bpmndixmi:BPMNLabel">
				<xsl:call-template name="BPMNLabelTemplate" />
			</label>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="BPMNLabelTemplate">
		<xsl:call-template name="LabelTemplate" />
		<xsl:if test="@labelStyle">
			<xsl:attribute name="labelStyle"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="@labelStyle" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="BPMNLabelStyleTemplate">
		<xsl:call-template name="StyleTemplate" />
		<xsl:for-each select="dc:Font">
			<font xmi:type="dcxmi:Font">
				<xsl:call-template name="FontTemplate" />
			</font>
		</xsl:for-each>
	</xsl:template>

	<!-- v2 -->
	<xsl:template name="BPMNEdgeTemplate">
		<!-- Add all attributes before any child elements -->
		<xsl:if test="@messageVisibleKind">
			<xsl:attribute name="messageVisibleKind">
				<xsl:value-of select="@messageVisibleKind" />
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="@bpmnElement">
			<xsl:attribute name="bpmnElement">
				<xsl:call-template name="concat">
					<xsl:with-param name="nodeset" select="@bpmnElement" />
				</xsl:call-template>
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="@sourceElement">
			<xsl:attribute name="sourceElement">
				<xsl:call-template name="concat">
					<xsl:with-param name="nodeset" select="@sourceElement" />
				</xsl:call-template>
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="@targetElement">
			<xsl:attribute name="targetElement">
				<xsl:call-template name="concat">
					<xsl:with-param name="nodeset" select="@targetElement" />
				</xsl:call-template>
			</xsl:attribute>
		</xsl:if>

		<!-- Now call the LabeledEdgeTemplate -->
		<xsl:call-template name="LabeledEdgeTemplate" />

		<!-- Add child elements -->
		<xsl:for-each select="bpmndi:BPMNLabel">
			<label xmi:type="bpmndixmi:BPMNLabel">
				<xsl:call-template name="BPMNLabelTemplate" />
			</label>
		</xsl:for-each>
	</xsl:template>

	<!-- v2 -->
	<xsl:template name="BPMNDiagramTemplate">
		<!-- Add all attributes before any child elements -->
		<xsl:if test="@name">
			<xsl:attribute name="name">
				<xsl:value-of select="@name" />
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="@documentation">
			<xsl:attribute name="documentation">
				<xsl:value-of select="@documentation" />
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="@resolution">
			<xsl:attribute name="resolution">
				<xsl:value-of select="@resolution" />
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:ownedStyle">
			<xsl:attribute name="ownedStyle">
				<xsl:call-template name="concat">
					<xsl:with-param name="nodeset" select="bpmn:ownedStyle" />
				</xsl:call-template>
			</xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:rootElement">
			<xsl:attribute name="rootElement">
				<xsl:call-template name="concat">
					<xsl:with-param name="nodeset" select="bpmn:rootElement" />
				</xsl:call-template>
			</xsl:attribute>
		</xsl:if>

		<!-- Now add child elements -->
		<xsl:for-each select="bpmndi:BPMNPlane">
			<plane xmi:type="bpmndixmi:BPMNPlane">
				<xsl:call-template name="BPMNPlaneTemplate" />
			</plane>
		</xsl:for-each>
		<xsl:for-each select="bpmndi:BPMNLabelStyle">
			<labelStyle xmi:type="bpmndixmi:BPMNLabelStyle">
				<xsl:call-template name="BPMNLabelStyleTemplate" />
			</labelStyle>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="DiagramElementTemplate">
		<xsl:if test="bpmn:owningDiagram">
			<xsl:attribute name="owningDiagram"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:owningDiagram" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:owningElement">
			<xsl:attribute name="owningElement"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:owningElement" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:ownedElement">
			<xsl:attribute name="ownedElement"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:ownedElement" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:modelElement">
			<xsl:attribute name="modelElement"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:modelElement" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:style">
			<xsl:attribute name="style"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:style" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="DiagramTemplate">
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@documentation">
			<xsl:attribute name="documentation"> <xsl:value-of select="@documentation" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@resolution">
			<xsl:attribute name="resolution"> <xsl:value-of select="@resolution" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:ownedStyle">
			<xsl:attribute name="ownedStyle"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:ownedStyle" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:rootElement">
			<xsl:attribute name="rootElement"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:rootElement" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="StyleTemplate">
	</xsl:template>

	<xsl:template name="NodeTemplate">
		<xsl:call-template name="DiagramElementTemplate" />
	</xsl:template>

	<xsl:template name="EdgeTemplate">
		<xsl:call-template name="DiagramElementTemplate" />
		<xsl:if test="bpmn:source">
			<xsl:attribute name="source"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:source" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:if test="bpmn:target">
			<xsl:attribute name="target"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:target" /></xsl:call-template></xsl:attribute>
		</xsl:if>
		<xsl:for-each select="di:waypoint">
			<waypoint xmi:type="dixmi:Point">
				<xsl:call-template name="PointTemplate" />
			</waypoint>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="ShapeTemplate">
		<xsl:call-template name="NodeTemplate" />
		<xsl:for-each select="dc:Bounds">
			<bounds xmi:type="dcxmi:Bounds">
				<xsl:call-template name="BoundsTemplate" />
			</bounds>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="LabeledEdgeTemplate">
		<xsl:call-template name="EdgeTemplate" />
		<xsl:if test="bpmn:ownedLabel">
			<xsl:attribute name="ownedLabel"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:ownedLabel" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="LabelTemplate">
		<xsl:call-template name="NodeTemplate" />
		<xsl:for-each select="dc:Bounds">
			<bounds xmi:type="dcxmi:Bounds">
				<xsl:call-template name="BoundsTemplate" />
			</bounds>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="LabeledShapeTemplate">
		<xsl:call-template name="ShapeTemplate" />
		<xsl:if test="bpmn:ownedLabel">
			<xsl:attribute name="ownedLabel"><xsl:call-template
				name="concat"><xsl:with-param name="nodeset" select="bpmn:ownedLabel" /></xsl:call-template></xsl:attribute>
		</xsl:if>
	</xsl:template>


	<!-- v2 -->
	<xsl:template name="PlaneTemplate">
		<!-- Add all attributes before any child elements -->
		<xsl:if test="@bpmnElement">
			<xsl:attribute name="bpmnElement">
				<xsl:call-template name="concat">
					<xsl:with-param name="nodeset" select="@bpmnElement" />
				</xsl:call-template>
			</xsl:attribute>
		</xsl:if>

		<!-- Now add child elements -->
		<xsl:for-each select="bpmndi:BPMNEdge">
			<planeElement xmi:type="bpmndixmi:BPMNEdge">
				<xsl:call-template name="BPMNEdgeTemplate" />
			</planeElement>
		</xsl:for-each>
		<xsl:for-each select="bpmndi:BPMNShape">
			<planeElement xmi:type="bpmndixmi:BPMNShape">
				<xsl:call-template name="BPMNShapeTemplate" />
			</planeElement>
		</xsl:for-each>
	</xsl:template>

	<xsl:template name="FontTemplate">
		<xsl:if test="@name">
			<xsl:attribute name="name"> <xsl:value-of select="@name" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@size">
			<xsl:attribute name="size"> <xsl:value-of select="@size" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@isBold">
			<xsl:attribute name="isBold"> <xsl:value-of select="@isBold" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@isItalic">
			<xsl:attribute name="isItalic"> <xsl:value-of select="@isItalic" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@isUnderline">
			<xsl:attribute name="isUnderline"> <xsl:value-of select="@isUnderline" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@isStrikeThrough">
			<xsl:attribute name="isStrikeThrough"> <xsl:value-of select="@isStrikeThrough" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="PointTemplate">
		<xsl:if test="@x">
			<xsl:attribute name="x"> <xsl:value-of select="@x" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@y">
			<xsl:attribute name="y"> <xsl:value-of select="@y" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>

	<xsl:template name="BoundsTemplate">
		<xsl:if test="@x">
			<xsl:attribute name="x"> <xsl:value-of select="@x" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@y">
			<xsl:attribute name="y"> <xsl:value-of select="@y" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@width">
			<xsl:attribute name="width"> <xsl:value-of select="@width" /> </xsl:attribute>
		</xsl:if>
		<xsl:if test="@height">
			<xsl:attribute name="height"> <xsl:value-of select="@height" /> </xsl:attribute>
		</xsl:if>
	</xsl:template>


</xsl:stylesheet>