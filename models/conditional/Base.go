package conditional

// @type STR_PTR
type STR_PTR *string

// ConditionalScriptFormat ...
type ConditionalScriptFormat interface {
	SetScriptFormat(format string)
	GetScriptFormat() STR_PTR
}

// ConditionalScript ...
type ConditionalScript interface {
	SetScript(script string)
	GetScript() STR_PTR
}
