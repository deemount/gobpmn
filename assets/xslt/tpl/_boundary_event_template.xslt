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