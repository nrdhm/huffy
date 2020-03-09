package core

import (
	"reflect"
	"testing"
)

type newTreeTest struct {
	input    []SymbolCount
	expected map[Symbol]Code
}

var newTreeTestCases = []newTreeTest{
	{
		input: []SymbolCount{
			{Symbol: " ", Count: 26974},
			{Symbol: "B", Count: 1275},
			{Symbol: "D", Count: 2823},
			{Symbol: "F", Count: 1757},
			{Symbol: "H", Count: 3279},
			{Symbol: "J", Count: 152},
			{Symbol: "L", Count: 3114},
			{Symbol: "N", Count: 5626},
			{Symbol: "P", Count: 2195},
			{Symbol: "R", Count: 5173},
			{Symbol: "T", Count: 8375},
			{Symbol: "V", Count: 928},
			{Symbol: "X", Count: 369},
			{Symbol: "Z", Count: 60},
			{Symbol: "A", Count: 6538},
			{Symbol: "C", Count: 3115},
			{Symbol: "E", Count: 9917},
			{Symbol: "G", Count: 1326},
			{Symbol: "I", Count: 6430},
			{Symbol: "K", Count: 317},
			{Symbol: "M", Count: 1799},
			{Symbol: "O", Count: 6261},
			{Symbol: "Q", Count: 113},
			{Symbol: "S", Count: 5784},
			{Symbol: "U", Count: 2360},
			{Symbol: "W", Count: 987},
			{Symbol: "Y", Count: 1104},
		},
		expected: map[Symbol]Code{
			" ": "01",
			"T": "0010",
			"A": "1000",
			"I": "1001",
			"E": "0000",
			"O": "1010",
			"S": "1100",
			"R": "1111",
			"N": "1101",
			"L": "10111",
			"C": "10110",
			"D": "11100",
			"H": "00111",
			"G": "111010",
			"B": "111011",
			"P": "000101",
			"F": "001101",
			"U": "000100",
			"M": "001100",
			"W": "0001110",
			"V": "0001111",
			"Y": "0001100",
			"X": "00011011",
			"K": "000110101",
			"J": "0001101001",
			"Q": "00011010000",
			"Z": "000110100010",
			"":  "000110100011",
		},
	},
}

func TestNewTree(t *testing.T) {
	for _, test := range newTreeTestCases {
		tree := NewTree(test.input)
		actual := GetCompressTable(tree)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("NewTree test [%v], expected [%v], actual [%v]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkNewTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range newTreeTestCases {
			tree := NewTree(test.input)
			GetCompressTable(tree)
		}
	}
}
