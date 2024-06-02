// Convinience functions for calculating block rewards for various bitcoin-like networks.
package blockreward

// Returns the block subsidy at the given height for the given network.
// The subsidy is returned in satoshis and follows the same halving schedule
// as the upstream reference implementation.
func SubsidyAtHeight(network Network, height int64) int64 {
	switch network {
	case BitcoinMainnet:
		// https://github.com/bitcoin/bitcoin/blob/457e1846d2bf6ef9d54b9ba1a330ba8bbff13091/src/kernel/chainparams.cpp#L77
		// https://github.com/bitcoin/bitcoin/blob/457e1846d2bf6ef9d54b9ba1a330ba8bbff13091/src/validation.cpp#L1813
		return SubsidyAtHeightWithConfig(height, 210000, 50*100000000, 64)
	case LitecoinMainnet:
		// https://github.com/litecoin-project/litecoin/blob/cd1660afaf5b31a80e797668b12b5b3933844842/src/chainparams.cpp#L69
		// https://github.com/litecoin-project/litecoin/blob/cd1660afaf5b31a80e797668b12b5b3933844842/src/validation.cpp#L1268
		return SubsidyAtHeightWithConfig(height, 840000, 50*100000000, 64)
	}
	return 0
}

// Returns the block subsidy at the given height for a network with the given
// subsidy interval, initial subsidy, and halving limit.
//
// Use SubsidyAtHeight for well-defined networks, and this function for networks
// whose parameters are not defined in this package.
func SubsidyAtHeightWithConfig(height int64, subsidyInterval int64, initialSubsidy int64, halvingLimit int64) int64 {
	halvings := height / subsidyInterval
	if halvings >= halvingLimit {
		return 0
	}

	return initialSubsidy >> uint(halvings)
}
