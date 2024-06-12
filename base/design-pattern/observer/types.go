package observer

// Governor 省长
type Governor interface {
	GiveOrder(order Order)
	Update(content string)
	GetContent() (content string)
}

// Order 任务
type Order interface {
	Leadership(governor Governor)
	Update()
	GetContent() (content string)
}
