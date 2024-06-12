// Package builder 实现
package builder

// HuaweiPhone 华为手机
type HuaweiPhone struct{}

// Name 名称
func (h *HuaweiPhone) Name() string {
	return "华为手机"
}

// Color 颜色
func (h *HuaweiPhone) Color() string {
	return "金色"
}

// Size 尺寸
func (h *HuaweiPhone) Size() int {
	return 5
}

// ApplePhone 苹果手机
type ApplePhone struct {
}

// Name 名称
func (h *ApplePhone) Name() string {
	return "苹果手机"
}

// Color 颜色
func (h *ApplePhone) Color() string {
	return "银白色"
}

// Size 尺寸
func (h *ApplePhone) Size() int {
	return 4
}

// foxConn 富士康
type foxConn struct {
	productList []Product
}

// GetCount 获取总数
func (f *foxConn) GetCount() int {
	return len(f.productList)
}

// GetProductList 获取产品列表
func (f *foxConn) GetProductList() []Product {
	return f.productList
}

// AddProduct 添加产品
func (f *foxConn) AddProduct(product Product) {
	f.productList = append(f.productList, product)
}

// GetFoxConn 获取工厂.
func GetFoxConn() Factory {
	f := new(foxConn)
	f.productList = make([]Product, 0, 3)
	return f
}

// MingDirector 指导者（拥有人小明）
func MingDirector() (factory Factory) {
	conn := GetFoxConn()
	conn.AddProduct(new(HuaweiPhone))
	conn.AddProduct(new(ApplePhone))
	conn.AddProduct(new(HuaweiPhone))
	return conn
}

// KingDirector 指导者（拥有人小王）
func KingDirector() (factory Factory) {
	conn := GetFoxConn()
	conn.AddProduct(new(ApplePhone))
	conn.AddProduct(new(HuaweiPhone))
	return conn
}
