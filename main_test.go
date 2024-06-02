package blockreward

import "testing"

func TestSubsidyAtHeight(t *testing.T) {
	tests := []struct {
		network Network
		height  int64
		want    int64
	}{
		// Bitcoin Mainnet tests
		{BitcoinMainnet, 0, 5000000000},
		{BitcoinMainnet, 209999, 5000000000},
		{BitcoinMainnet, 210000, 2500000000},
		{BitcoinMainnet, 210001, 2500000000},
		{BitcoinMainnet, 419999, 2500000000},
		{BitcoinMainnet, 420000, 1250000000},
		{BitcoinMainnet, 630000, 625000000},
		{BitcoinMainnet, 840000, 312500000},
		{BitcoinMainnet, 1050000, 156250000},
		{BitcoinMainnet, 1260000, 78125000},
		{BitcoinMainnet, 1470000, 39062500},
		{BitcoinMainnet, 1680000, 19531250},
		{BitcoinMainnet, 1890000, 9765625},
		{BitcoinMainnet, 2100000, 4882812},
		{BitcoinMainnet, 2310000, 2441406},
		{BitcoinMainnet, 2520000, 1220703},
		{BitcoinMainnet, 2730000, 610351},
		{BitcoinMainnet, 2940000, 305175},
		{BitcoinMainnet, 3150000, 152587},
		{BitcoinMainnet, 3360000, 76293},
		{BitcoinMainnet, 3570000, 38146},
		{BitcoinMainnet, 3780000, 19073},
		{BitcoinMainnet, 3990000, 9536},
		{BitcoinMainnet, 4200000, 4768},
		{BitcoinMainnet, 4410000, 2384},
		{BitcoinMainnet, 4620000, 1192},
		{BitcoinMainnet, 4830000, 596},
		{BitcoinMainnet, 5040000, 298},
		{BitcoinMainnet, 5250000, 149},
		{BitcoinMainnet, 5460000, 74},
		{BitcoinMainnet, 5670000, 37},
		{BitcoinMainnet, 5880000, 18},
		{BitcoinMainnet, 6090000, 9},
		{BitcoinMainnet, 6300000, 4},
		{BitcoinMainnet, 6510000, 2},
		{BitcoinMainnet, 6720000, 1},
		{BitcoinMainnet, 6930000, 0},
		{BitcoinMainnet, 10000000, 0},

		// Litecoin Mainnet tests
		{LitecoinMainnet, 0, 5000000000},
		{LitecoinMainnet, 839999, 5000000000},
		{LitecoinMainnet, 840000, 2500000000},
		{LitecoinMainnet, 840001, 2500000000},
		{LitecoinMainnet, 1679999, 2500000000},
		{LitecoinMainnet, 1680000, 1250000000},
		{LitecoinMainnet, 2520000, 625000000},
		{LitecoinMainnet, 3360000, 312500000},
		{LitecoinMainnet, 4200000, 156250000},
		{LitecoinMainnet, 5040000, 78125000},
		{LitecoinMainnet, 5880000, 39062500},
		{LitecoinMainnet, 6720000, 19531250},
		{LitecoinMainnet, 7560000, 9765625},
		{LitecoinMainnet, 8400000, 4882812},
		{LitecoinMainnet, 9240000, 2441406},
		{LitecoinMainnet, 10080000, 1220703},
		{LitecoinMainnet, 10920000, 610351},
		{LitecoinMainnet, 11760000, 305175},
		{LitecoinMainnet, 12600000, 152587},
		{LitecoinMainnet, 13440000, 76293},
		{LitecoinMainnet, 14280000, 38146},
		{LitecoinMainnet, 15120000, 19073},
		{LitecoinMainnet, 15960000, 9536},
		{LitecoinMainnet, 16800000, 4768},
		{LitecoinMainnet, 17640000, 2384},
		{LitecoinMainnet, 18480000, 1192},
		{LitecoinMainnet, 19320000, 596},
		{LitecoinMainnet, 20160000, 298},
		{LitecoinMainnet, 21000000, 149},
		{LitecoinMainnet, 21840000, 74},
		{LitecoinMainnet, 22680000, 37},
		{LitecoinMainnet, 23520000, 18},
		{LitecoinMainnet, 24360000, 9},
		{LitecoinMainnet, 25200000, 4},
		{LitecoinMainnet, 26040000, 2},
		{LitecoinMainnet, 26880000, 1},
		{LitecoinMainnet, 27720000, 0},
	}
	for _, test := range tests {
		if got := SubsidyAtHeight(test.network, test.height); got != test.want {
			t.Errorf("SubsidyAtHeight(%v, %v) = %v, want %v", test.network, test.height, got, test.want)
		}
	}
}

func TestSubsidySchedule(t *testing.T) {
	tests := []struct {
		network Network
		want    []Subsidy
	}{
		// Bitcoin Mainnet tests
		{
			BitcoinMainnet,
			[]Subsidy{
				{0, 5000000000},
				{210000, 2500000000},
				{420000, 1250000000},
				{630000, 625000000},
				{840000, 312500000},
				{1050000, 156250000},
				{1260000, 78125000},
				{1470000, 39062500},
				{1680000, 19531250},
				{1890000, 9765625},
				{2100000, 4882812},
				{2310000, 2441406},
				{2520000, 1220703},
				{2730000, 610351},
				{2940000, 305175},
				{3150000, 152587},
				{3360000, 76293},
				{3570000, 38146},
				{3780000, 19073},
				{3990000, 9536},
				{4200000, 4768},
				{4410000, 2384},
				{4620000, 1192},
				{4830000, 596},
				{5040000, 298},
				{5250000, 149},
				{5460000, 74},
				{5670000, 37},
				{5880000, 18},
				{6090000, 9},
				{6300000, 4},
				{6510000, 2},
				{6720000, 1},
				{6930000, 0},
			},
		},
		// Litecoin Mainnet tests
		{
			LitecoinMainnet,
			[]Subsidy{
				{0, 5000000000},
				{840000, 2500000000},
				{1680000, 1250000000},
				{2520000, 625000000},
				{3360000, 312500000},
				{4200000, 156250000},
				{5040000, 78125000},
				{5880000, 39062500},
				{6720000, 19531250},
				{7560000, 9765625},
				{8400000, 4882812},
				{9240000, 2441406},
				{10080000, 1220703},
				{10920000, 610351},
				{11760000, 305175},
				{12600000, 152587},
				{13440000, 76293},
				{14280000, 38146},
				{15120000, 19073},
				{15960000, 9536},
				{16800000, 4768},
				{17640000, 2384},
				{18480000, 1192},
				{19320000, 596},
				{20160000, 298},
				{21000000, 149},
				{21840000, 74},
				{22680000, 37},
				{23520000, 18},
				{24360000, 9},
				{25200000, 4},
				{26040000, 2},
				{26880000, 1},
				{27720000, 0},
			},
		},
	}

	for _, test := range tests {
		subsidies := SubsidySchedule(test.network)
		if len(subsidies) != len(test.want) {
			t.Errorf("SubsidySchedule(%v) = %v, want %v", test.network, subsidies, test.want)
			continue
		}

		for i, subsidy := range subsidies {
			if subsidy != test.want[i] {
				t.Errorf("SubsidySchedule(%v) = %v, want %v", test.network, subsidies, test.want)
				break
			}
		}
	}
}
