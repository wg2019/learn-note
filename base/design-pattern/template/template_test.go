package template

import "testing"

func TestOrderMeal(t *testing.T) {
	OrderMeal(new(ChineseFood))
	OrderMeal(new(WesternFood))
}
