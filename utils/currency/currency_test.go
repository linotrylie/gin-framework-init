package currency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYuan2Fen(t *testing.T) {
	var tests = []struct {
		input  float64
		output int64
	}{
		{100.12477777777, 10012},
		{100.86666666, 10087},
		{100, 10000},
		{1, 100},
		{0.1, 10},
		{0.01, 1},
	}

	for _, test := range tests {
		fen := Yuan2Fen(test.input)
		assert.Equal(t, fen, test.output)
	}

}

func TestFen2Yuan(t *testing.T) {
	var tests = []struct {
		input  uint64
		output string
	}{
		{10012477777777, "100124777777.77"},
		{10086666666, "100866666.66"},
		{100, "1"},
		{10, "0.1"},
		{1000, "10"},
		{10000, "100"},
	}

	for _, test := range tests {
		fen := Fen2Yuan(test.input)
		assert.Equal(t, fen, test.output)
	}
}
