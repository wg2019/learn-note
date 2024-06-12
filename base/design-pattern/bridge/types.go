package bridge

// Plasticine 橡皮泥
type Plasticine interface {
	SetColor(color Color)
	SetShape(shape Shape)
	Draw() string
}

// Shape 形状
type Shape interface {
	Making() string
}

// Color 颜色
type Color interface {
	BePaint() string
}
