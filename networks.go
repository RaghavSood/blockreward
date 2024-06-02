package blockreward

// Network represents a blockchain network.
// This is a blockreward-specific type and does not reference any external
// network identifiers.
type Network int

const (
	BitcoinMainnet Network = iota
	LitecoinMainnet
)
