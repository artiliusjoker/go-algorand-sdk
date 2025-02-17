package models

// CompileResponse teal compile Result
type CompileResponse struct {
	// Hash base32 SHA512_256 of program bytes (Address style)
	Hash string `json:"hash,omitempty"`

	// Result base64 encoded program bytes
	Result string `json:"result,omitempty"`
}
