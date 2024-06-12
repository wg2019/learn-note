package bridge

import "fmt"

// toy 玩具
type toy struct {
	color Color
	shape Shape
}

// NewToy 新玩具
func NewToy() Plasticine {
	t := new(toy)
	return t
}

// SetColor 设置颜色
func (t *toy) SetColor(color Color) {
	t.color = color
}

// SetShape 设置形状
func (t *toy) SetShape(shape Shape) {
	t.shape = shape
}

// Draw 画
func (t *toy) Draw() string {
	return fmt.Sprintf("%s的%s", t.color.BePaint(), t.shape.Making())
}

// Circle 圆形
type Circle struct {
}

// Making 生产
func (c *Circle) Making() string {
	return "圆形"
}

// Square 正方形
type Square struct {
}

// Making 生产
func (s *Square) Making() string {
	return "正方形"
}

// White 白色
type White struct {
}

// BePaint 涂料
func (w *White) BePaint() string {
	return "白色"
}

// Black 黑色
type Black struct {
}

// BePaint 涂料
func (b *Black) BePaint() string {
	return "黑色"
}
