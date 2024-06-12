package strategy

// Strategy 策略模式
type Strategy interface {
	Calculation(number1, number2 int) int
}
