package utils

import (
	"math/big"
)

func ConvertFrom18(amount *big.Int) float64 {
	floatAmount := new(big.Float).SetInt(amount)
	result := floatAmount.Quo(floatAmount, big.NewFloat(1e18))
	float64_Res, _ := result.Float64()
	return float64_Res
}
