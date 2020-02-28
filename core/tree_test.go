package core

import (
	"reflect"
	"testing"
)

type newTreeTest struct {
	input    map[Symbol]int
	expected map[Symbol]Code
}

var newTreeTestCases = []newTreeTest{
	{
		input: map[Symbol]int{
			" ": 26974,
			"B": 1275,
			"D": 2823,
			"F": 1757,
			"H": 3279,
			"J": 152,
			"L": 3114,
			"N": 5626,
			"P": 2195,
			"R": 5173,
			"T": 8375,
			"V": 928,
			"X": 369,
			"Z": 60,
			"A": 6538,
			"C": 3115,
			"E": 9917,
			"G": 1326,
			"I": 6430,
			"K": 317,
			"M": 1799,
			"O": 6261,
			"Q": 113,
			"S": 5784,
			"U": 2360,
			"W": 987,
			"Y": 1104,
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

type fromCompressTableTest struct {
	input    map[Symbol]Code
	expected *HuffmanTree
}

var fromCompressTableTestCases = []fromCompressTableTest{
	{input: map[Symbol]Code{}, expected: nil},
	{
		input: map[Symbol]Code{"A": "01", "B": "010", "C": "10", "": "011"},
		expected: &HuffmanTree{
			zeroBranch: &HuffmanTree{
				oneBranch: &HuffmanTree{symbol: "A", zeroBranch: &HuffmanTree{symbol: "B"}, oneBranch: &HuffmanTree{eof: true}},
			},
			oneBranch: &HuffmanTree{zeroBranch: &HuffmanTree{symbol: "C"}},
		},
	},
}

func TestFromCompressTable(t *testing.T) {
	for _, test := range fromCompressTableTestCases {
		actual := FromCompressTable(test.input)
		if !actual.equal(test.expected) {
			t.Errorf("FromCompressTable test [%v], expected [%v], actual [%v]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkFromCompressTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range fromCompressTableTestCases {
			FromCompressTable(test.input)
		}
	}
}
