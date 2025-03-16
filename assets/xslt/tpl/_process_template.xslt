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