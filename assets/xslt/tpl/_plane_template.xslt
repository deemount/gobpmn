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