package strategy

// CalculationAdd 加
type CalculationAdd struct {
}

// Calculation 计算
func (c *CalculationAdd) Calculation(number1, number2 int) int {
	return number1 + number2
}

// CalculationSubtract 剪
type CalculationSubtract struct {
}

// Calculation 计算
func (c *CalculationSubtract) Calculation(number1, number2 int) int {
	return number1 - number2
}

// CalculationMultiply 除
type CalculationMultiply struct {
}

// Calculation 计算
func (c *CalculationMultiply) Calculation(number1, number2 int) int {
	return number1 * number2
}

// Calculation 计算
func Calculation(strategy Strategy, number1, number2 int) int {
	return strategy.Calculation(number1, number2)
}
