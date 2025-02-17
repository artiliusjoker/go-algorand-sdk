package models

// DryrunState stores the TEAL eval step data
type DryrunState struct {
	// Error evaluation error if any
	Error string `json:"error,omitempty"`

	// Line line number
	Line uint64 `json:"line,omitempty"`

	// Pc program counter
	Pc uint64 `json:"pc,omitempty"`

	// Scratch
	Scratch []TealValue `json:"scratch,omitempty"`

	// Stack
	Stack []TealValue `json:"stack,omitempty"`
}
