package models

// DryrunTxnResult dryrunTxnResult contains any LogicSig or ApplicationCall program
// debug information and state updates from a dryrun.
type DryrunTxnResult struct {
	// AppCallMessages
	AppCallMessages []string `json:"app-call-messages,omitempty"`

	// AppCallTrace
	AppCallTrace []DryrunState `json:"app-call-trace,omitempty"`

	// Disassembly disassembled program line by line.
	Disassembly []string `json:"disassembly,omitempty"`

	// GlobalDelta application state delta.
	GlobalDelta []EvalDeltaKeyValue `json:"global-delta,omitempty"`

	// LocalDeltas
	LocalDeltas []AccountStateDelta `json:"local-deltas,omitempty"`

	// LogicSigMessages
	LogicSigMessages []string `json:"logic-sig-messages,omitempty"`

	// LogicSigTrace
	LogicSigTrace []DryrunState `json:"logic-sig-trace,omitempty"`
}
