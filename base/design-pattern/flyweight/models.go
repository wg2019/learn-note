package flyweight

// GoPiece 围棋子
type GoPiece struct {
	typeName string
	position [2]int
}

// GetType 获取类型
func (g *GoPiece) GetType() string {
	return g.typeName
}

// SetType 设置类型
func (g *GoPiece) SetType(typeName string) {
	g.typeName = typeName
}

// GetPosition 获取位置
func (g *GoPiece) GetPosition() [2]int {
	return g.position
}

// SetPosition 设置位置
func (g *GoPiece) SetPosition(position [2]int) {
	g.position = position
}

// ChessBox 棋盒
type ChessBox map[string]Chessman

var chessBox = make(ChessBox)

// GetGoPiece 获取棋子
func GetGoPiece(typeName string) Chessman {
	chessman, ok := chessBox[typeName]
	if ok {
		return chessman
	}
	chessman = new(GoPiece)
	chessman.SetType(typeName)
	chessBox[typeName] = chessman
	return chessman
}
