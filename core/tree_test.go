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
			{symbol: " ", count: 26974},
			{symbol: "B", count: 1275},
			{symbol: "D", count: 2823},
			{symbol: "F", count: 1757},
			{symbol: "H", count: 3279},
			{symbol: "J", count: 152},
			{symbol: "L", count: 3114},
			{symbol: "N", count: 5626},
			{symbol: "P", count: 2195},
			{symbol: "R", count: 5173},
			{symbol: "T", count: 8375},
			{symbol: "V", count: 928},
			{symbol: "X", count: 369},
			{symbol: "Z", count: 60},
			{symbol: "A", count: 6538},
			{symbol: "C", count: 3115},
			{symbol: "E", count: 9917},
			{symbol: "G", count: 1326},
			{symbol: "I", count: 6430},
			{symbol: "K", count: 317},
			{symbol: "M", count: 1799},
			{symbol: "O", count: 6261},
			{symbol: "Q", count: 113},
			{symbol: "S", count: 5784},
			{symbol: "U", count: 2360},
			{symbol: "W", count: 987},
			{symbol: "Y", count: 1104},
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
