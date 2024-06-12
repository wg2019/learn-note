package facade

// Deduplication 去重方法
type Deduplication interface {
	Put(interface{}) bool
	Exist(interface{}) bool
	Delete(interface{}) bool
}
