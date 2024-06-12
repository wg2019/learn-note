package flyweight

// Chessman 棋子
type Chessman interface {
	GetType() string
	SetType(typeName string)
	GetPosition() [2]int
	SetPosition(position [2]int)
}
