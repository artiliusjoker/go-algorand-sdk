package models

// AccountParticipation accountParticipation describes the parameters used by this
// account in consensus protocol.
type AccountParticipation struct {
	// SelectionParticipationKey (sel) Selection public key (if any) currently
	// registered for this round.
	SelectionParticipationKey []byte `json:"selection-participation-key,omitempty"`

	// VoteFirstValid (voteFst) First round for which this participation is valid.
	VoteFirstValid uint64 `json:"vote-first-valid,omitempty"`

	// VoteKeyDilution (voteKD) Number of subkeys in each batch of participation keys.
	VoteKeyDilution uint64 `json:"vote-key-dilution,omitempty"`

	// VoteLastValid (voteLst) Last round for which this participation is valid.
	VoteLastValid uint64 `json:"vote-last-valid,omitempty"`

	// VoteParticipationKey (vote) root participation public key (if any) currently
	// registered for this round.
	VoteParticipationKey []byte `json:"vote-participation-key,omitempty"`
}
