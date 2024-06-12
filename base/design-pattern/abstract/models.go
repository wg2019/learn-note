package abstract

// banana 香蕉.
type banana struct {
	name string
}

// NewBanana 创建香蕉数据
func NewBanana(name string) *banana {
	b := new(banana)
	b.name = name
	return b
}

// Name 名称
func (b *banana) Name() string {
	return b.name
}

// Color 颜色.
func (b *banana) Color() string {
	return "yellow"
}

// Weight 重量(单位克).
func (b *banana) Weight() int {
	return 200
}

// apple 苹果
type apple struct {
	name string
}

// NewApple 创建苹果数据
func NewApple(name string) *apple {
	a := new(apple)
	a.name = name
	return a
}
func (a *apple) Name() string {
	return a.name
}

// Color 颜色.
func (a *apple) Color() string {
	return "red"
}

// Weight 重量(单位克).
func (a *apple) Weight() int {
	return 500
}

// ShanDong 山东工厂
type ShanDong struct{}

// GetBanana 山东香蕉
func (s *ShanDong) GetBanana() Product {
	return NewBanana("山东香蕉")
}

// GetApple 山东苹果
func (s *ShanDong) GetApple() Product {
	return NewApple("山东苹果")
}

// ShanDong 云南工厂
type YunNan struct{}

// GetBanana 云南香蕉
func (y *YunNan) GetBanana() Product {
	return NewBanana("云南香蕉")
}

// GetApple 云南苹果
func (y *YunNan) GetApple() Product {
	return NewApple("云南苹果")
}
