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