package abstract

// Product 定义产品.
type Product interface {
	// 名称
	Name() string
	// 颜色
	Color() string
	// 重量
	Weight() int
}

// Factory 定义工厂.
type Factory interface {
	// 获取香蕉数据
	GetBanana() Product
	// 获取苹果数据
	GetApple() Product
}
