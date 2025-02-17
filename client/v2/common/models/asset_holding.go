package models

// AssetHolding describes an asset held by an account.
// Definition:
// data/basics/userBalance.go : AssetHolding
type AssetHolding struct {
	// Amount (a) number of units held.
	Amount uint64 `json:"amount,omitempty"`

	// AssetId asset ID of the holding.
	AssetId uint64 `json:"asset-id,omitempty"`

	// Creator address that created this asset. This is the address where the
	// parameters for this asset can be found, and also the address where unwanted
	// asset units can be sent in the worst case.
	Creator string `json:"creator,omitempty"`

	// Deleted whether or not the asset holding is currently deleted from its account.
	Deleted bool `json:"deleted,omitempty"`

	// IsFrozen (f) whether or not the holding is frozen.
	IsFrozen bool `json:"is-frozen,omitempty"`

	// OptedInAtRound round during which the account opted into this asset holding.
	OptedInAtRound uint64 `json:"opted-in-at-round,omitempty"`

	// OptedOutAtRound round during which the account opted out of this asset holding.
	OptedOutAtRound uint64 `json:"opted-out-at-round,omitempty"`
}
