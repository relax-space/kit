package base

import (
	"testing"

	"github.com/relax-space/go-kit/test"
)

func Test_toFixed_default(t *testing.T) {
	priceSetting := DefaultPriceSetting

	var num1 float64 = 11.231
	var num2 float64 = 11.236
	test.Equals(t, 11.23, ToFixed(num1, priceSetting))
	test.Equals(t, 11.24, ToFixed(num2, priceSetting))

}
func Test_toFixed_RoundDigit(t *testing.T) {
	priceSetting := DefaultPriceSetting
	priceSetting.RoundDigit = 3

	var num1 float64 = 11.2315
	var num2 float64 = 11.2364
	test.Equals(t, 11.232, ToFixed(num1, priceSetting))
	test.Equals(t, 11.236, ToFixed(num2, priceSetting))
}

func Test_toFixed_RoundStrategy(t *testing.T) {
	priceSetting := DefaultPriceSetting
	priceSetting.RoundStrategy = "ceil"

	var num1 float64 = 11.2315
	var num2 float64 = 11.2364
	test.Equals(t, 11.24, ToFixed(num1, priceSetting))
	test.Equals(t, 11.24, ToFixed(num2, priceSetting))
	priceSetting.RoundStrategy = "floor"
	test.Equals(t, 11.23, ToFixed(num1, priceSetting))
	test.Equals(t, 11.23, ToFixed(num2, priceSetting))
}

func Test_round(t *testing.T) {
	num := 11.236
	num2 := 11.6
	test.Equals(t, 11, Round(num))
	test.Equals(t, 12, Round(num2))
}
