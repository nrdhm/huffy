package core

import (
	"encoding/base64"
	"testing"
)

type compressTest struct {
	clear      string
	compressed string
}

var compressTestCases = []compressTest{
	{
		clear:      "abc",
		compressed: "Qf+BAwEBCEh1ZmZUcmVlAf+CAAEFAQRaZXJvAf+CAAEDT25lAf+CAAEDU3ltAQwAAQNDbnQBBAABA0VPRgECAAAAMv+CAQEBAwFiAQIAAQEBAQABBQEAAAEAAAEAAAICAAEDAWEBAgACBAABAwFjAQIAAgYAIgk=",
	},
	{
		clear:      "",
		compressed: "",
	},
}

func TestCompress(t *testing.T) {
	for _, test := range compressTestCases {
		compressed, err := Compress(test.clear)
		actual := base64.StdEncoding.EncodeToString([]byte(compressed))
		if err != nil {
			t.Error(err)
		}
		if actual != test.compressed {
			t.Errorf("Compress test [%v], expected [%v], actual [%v]", test.clear, test.compressed, actual)
		}
	}
}

func BenchmarkCompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range compressTestCases {
			Compress(test.clear)
		}
	}
}

func TestDecompress(t *testing.T) {
	for _, test := range compressTestCases {
		compressed, err := base64.StdEncoding.DecodeString(test.compressed)
		if err != nil {
			t.Error(err)
		}
		actual, err := Decompress(string(compressed))
		if err != nil {
			t.Error(err)
		}
		if actual != test.clear {
			t.Errorf("Decompress test [%v], expected [%v], actual [%v]", test.compressed, test.clear, actual)
		}
	}
}

func BenchmarkDecompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range compressTestCases {
			Decompress(test.clear)
		}
	}
}
