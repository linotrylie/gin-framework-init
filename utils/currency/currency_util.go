package currency

import (
	"github.com/shopspring/decimal"
)

// Fen2Yuan 分转元,
func Fen2Yuan(price uint64) string {
	d := decimal.New(1, 2) // 分除以100得到元
	result := decimal.NewFromInt(int64(price)).DivRound(d, 2).String()
	return result
}

// Yuan2Fen 元转分
func Yuan2Fen(price float64) int64 {
	d := decimal.New(1, 2)
	d1 := decimal.New(1, 0)
	dff := decimal.NewFromFloat(price).Mul(d).DivRound(d1, 0).IntPart()
	return dff
}
