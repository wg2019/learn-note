package observer

import "fmt"

// newYork 纽约
type newYork struct {
	content   string
	orderList []Order
}

func (n *newYork) GetContent() (content string) {
	return n.content
}

// GiveOrder 下达命令
func (n *newYork) GiveOrder(order Order) {
	n.orderList = append(n.orderList, order)
}

// Update 更新命令信息
func (n *newYork) Update(content string) {
	n.content = content
	for _, order := range n.orderList {
		order.Update()
	}
}

// GetNewYork 获取纽约市
func GetNewYork() Governor {
	n := new(newYork)
	n.orderList = make([]Order, 0, 3)
	n.content = ""
	return n
}

// policeStation 警察局
type policeStation struct {
	content  string
	governor Governor
}

func (p *policeStation) GetContent() (content string) {
	return p.content
}

// Leadership 领导
func (p *policeStation) Leadership(governor Governor) {
	p.governor = governor
	governor.GiveOrder(p)
}

// Update 更新命令
func (p *policeStation) Update() {
	if p.governor != nil {
		p.content = fmt.Sprintf("警察局：%s", p.governor.GetContent())
	}
}

// GetPoliceStation 获取警察局
func GetPoliceStation(governor Governor) Order {
	p := new(policeStation)
	p.Leadership(governor)
	p.Update()
	return p
}

// educationBureau 教育局
type educationBureau struct {
	content  string
	governor Governor
}

func (e *educationBureau) GetContent() (content string) {
	return e.content
}

// Leadership 领导
func (e *educationBureau) Leadership(governor Governor) {
	e.governor = governor
	governor.GiveOrder(e)
}

// Update 更新命令
func (e *educationBureau) Update() {
	if e.governor != nil {
		e.content = fmt.Sprintf("教育局：%s", e.governor.GetContent())
	}
}

func GetEducationBureau(governor Governor) Order {
	e := new(educationBureau)
	e.Leadership(governor)
	e.Update()
	return e
}
