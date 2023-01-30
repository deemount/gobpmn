package subprocesses

import "github.com/deemount/gobpmn/models/bpmn/attributes"

type DOCUMENTATION_PTR *attributes.Documentation
type EXTENSION_ELEMENTS_PTR *attributes.ExtensionElements

type ADHOC_SUBPROCESS_SLC []AdHocSubProcess
type CALL_ACTIVITY_SLC []CallActivity
type SUBPROCESS_SLC []SubProcess
type TRANSACTION_SLC []Transaction

type TADHOC_SUBPROCESS_SLC []AdHocSubProcess
type TCALL_ACTIVITY_SLC []CallActivity
type TSUBPROCESS_SLC []SubProcess
type TTRANSACTION_SLC []Transaction

type ADHOC_SUBPROCESS_PTR *AdHocSubProcess
type CALL_ACTIVITY_PTR *CallActivity
type SUBPROCESS_PTR *SubProcess
type TRANSACTION_PTR *Transaction
