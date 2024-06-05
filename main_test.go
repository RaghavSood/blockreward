package blockreward

import (
	"testing"
)

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
				{0, 5000000000, 5000000000},
				{210000, 2500000000, 1050002500000000},
				{420000, 1250000000, 1575001250000000},
				{630000, 625000000, 1837500625000000},
				{840000, 312500000, 1968750312500000},
				{1050000, 156250000, 2034375156250000},
				{1260000, 78125000, 2067187578125000},
				{1470000, 39062500, 2083593789062500},
				{1680000, 19531250, 2091796894531250},
				{1890000, 9765625, 2095898447265625},
				{2100000, 4882812, 2097949223632812},
				{2310000, 2441406, 2098974611711406},
				{2520000, 1220703, 2099487305750703},
				{2730000, 610351, 2099743652770351},
				{2940000, 305175, 2099871826175175},
				{3150000, 152587, 2099935912772587},
				{3360000, 76293, 2099967955966293},
				{3570000, 38146, 2099983977458146},
				{3780000, 19073, 2099991988099073},
				{3990000, 9536, 2099995993419536},
				{4200000, 4768, 2099997995974768},
				{4410000, 2384, 2099998997252384},
				{4620000, 1192, 2099999497891192},
				{4830000, 596, 2099999748210596},
				{5040000, 298, 2099999873370298},
				{5250000, 149, 2099999935950149},
				{5460000, 74, 2099999967240074},
				{5670000, 37, 2099999982780037},
				{5880000, 18, 2099999990550018},
				{6090000, 9, 2099999994330009},
				{6300000, 4, 2099999996220004},
				{6510000, 2, 2099999997060002},
				{6720000, 1, 2099999997480001},
				{6930000, 0, 2099999997690000},
			},
		},
		// Litecoin Mainnet tests
		{
			LitecoinMainnet,
			[]Subsidy{
				{0, 5000000000, 5000000000},
				{840000, 2500000000, 4200002500000000},
				{1680000, 1250000000, 6300001250000000},
				{2520000, 625000000, 7350000625000000},
				{3360000, 312500000, 7875000312500000},
				{4200000, 156250000, 8137500156250000},
				{5040000, 78125000, 8268750078125000},
				{5880000, 39062500, 8334375039062500},
				{6720000, 19531250, 8367187519531250},
				{7560000, 9765625, 8383593759765625},
				{8400000, 4882812, 8391796879882812},
				{9240000, 2441406, 8395898439521406},
				{10080000, 1220703, 8397949219340703},
				{10920000, 610351, 8398974609250351},
				{11760000, 305175, 8399487303785175},
				{12600000, 152587, 8399743650632587},
				{13440000, 76293, 8399871823636293},
				{14280000, 38146, 8399935909718146},
				{15120000, 19073, 8399967952339073},
				{15960000, 9536, 8399983973649536},
				{16800000, 4768, 8399991983884768},
				{17640000, 2384, 8399995989002384},
				{18480000, 1192, 8399997991561192},
				{19320000, 596, 8399998992840596},
				{20160000, 298, 8399999493480298},
				{21000000, 149, 8399999743800149},
				{21840000, 74, 8399999868960074},
				{22680000, 37, 8399999931120037},
				{23520000, 18, 8399999962200018},
				{24360000, 9, 8399999977320009},
				{25200000, 4, 8399999984880004},
				{26040000, 2, 8399999988240002},
				{26880000, 1, 8399999989920001},
				{27720000, 0, 8399999990760000},
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
				t.Errorf("SubsidySchedule(%v) = %v, want %v", test.network, subsidy, test.want[i])
				break
			}
		}
	}
}

func TestSupplyAtHeight(t *testing.T) {
	tests := []struct {
		network Network
		height  int64
		want    int64
	}{
		{BitcoinMainnet, 0, 5000000000},
		{BitcoinMainnet, 1, 10000000000},
		{BitcoinMainnet, 209999, 1050000000000000},
		{BitcoinMainnet, 210000, 1050002500000000},
		{BitcoinMainnet, 419999, 1575000000000000},
		{BitcoinMainnet, 420000, 1575001250000000},
		{BitcoinMainnet, 629999, 1837500000000000},
		{BitcoinMainnet, 630000, 1837500625000000},
		{BitcoinMainnet, 839999, 1968750000000000},
		{BitcoinMainnet, 840000, 1968750312500000},
		{BitcoinMainnet, 1049999, 2034375000000000},
		{BitcoinMainnet, 1050000, 2034375156250000},
		{BitcoinMainnet, 1259999, 2067187500000000},
		{BitcoinMainnet, 1260000, 2067187578125000},
		{BitcoinMainnet, 1469999, 2083593750000000},
		{BitcoinMainnet, 1470000, 2083593789062500},
		{BitcoinMainnet, 1679999, 2091796875000000},
		{BitcoinMainnet, 1680000, 2091796894531250},
		{BitcoinMainnet, 1889999, 2095898437500000},
		{BitcoinMainnet, 1890000, 2095898447265625},
		{BitcoinMainnet, 2099999, 2097949218750000},
		{BitcoinMainnet, 2100000, 2097949223632812},
		{BitcoinMainnet, 2309999, 2098974609270000},
		{BitcoinMainnet, 2310000, 2098974611711406},
		{BitcoinMainnet, 2519999, 2099487304530000},
		{BitcoinMainnet, 2520000, 2099487305750703},
		{BitcoinMainnet, 2729999, 2099743652160000},
		{BitcoinMainnet, 2730000, 2099743652770351},
	}

	for _, test := range tests {
		if got := SupplyAtHeight(test.network, test.height); got != test.want {
			t.Errorf("ExpectedSupplyAtHeight(%v, %v) = %v, want %v", test.network, test.height, got, test.want)
		}
	}
}
