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