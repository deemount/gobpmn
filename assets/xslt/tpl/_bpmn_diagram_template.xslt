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