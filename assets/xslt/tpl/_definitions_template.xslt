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