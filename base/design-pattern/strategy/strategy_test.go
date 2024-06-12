package strategy

import "testing"

func TestCalculation(t *testing.T) {
	calculationAdd := Calculation(new(CalculationAdd), 2, 3)
	t.Logf("calculationAdd: %d", calculationAdd)
	calculationSubtract := Calculation(new(CalculationSubtract), 2, 3)
	t.Logf("calculationSubtract: %d", calculationSubtract)
	calculationMultiply := Calculation(new(CalculationMultiply), 2, 3)
	t.Logf("calculationMultiply: %d", calculationMultiply)
}
