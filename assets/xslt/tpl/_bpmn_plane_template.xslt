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