package core

import (
	"reflect"
	"testing"
)

type tokenizeTest struct {
	input    string
	ctx      Context
	expected []Symbol
}

var tokenizeTestCases = []tokenizeTest{
	{
		"",
		Context{MaxSymbolLen: 3},
		[]Symbol{},
	},
	{
		"a",
		Context{MaxSymbolLen: 3},
		[]Symbol{},
	},
	{
		"a",
		Context{MaxSymbolLen: 1},
		[]Symbol{"a"},
	},
	{
		"quick",
		Context{MaxSymbolLen: 3},
		[]Symbol{"qui", "uic", "ick"},
	},
}

func TestTokenize(t *testing.T) {
	for _, test := range tokenizeTestCases {
		actual := symChanToArray(tokenize(test.ctx, test.input))
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("tokenize test [%v], expected [%v], actual [%v]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkTokenize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tokenizeTestCases {
			tokenize(test.ctx, test.input)
		}
	}
}
