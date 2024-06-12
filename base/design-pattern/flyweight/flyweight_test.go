package flyweight

import (
	"math/rand"
	"testing"
	"time"
)

func TestGetGoPiece(t *testing.T) {
	total := 100
	pieceList := make([]Chessman, 0, total)
	for i := 0; i < total; i++ {
		x := getPosition(i)
		y := getPosition(i + total)
		piece := GetGoPiece(getTypeName(i))
		piece.SetPosition([2]int{x, y})
		pieceList = append(pieceList, piece)
	}
	for _, piece := range pieceList {
		t.Logf("piece type: %s, x: %d, y: %d\n",
			piece.GetType(), piece.GetPosition()[0],
			piece.GetPosition()[1])
	}
}

func getTypeName(number int) string {
	if 0 == number%2 {
		return "白色"
	}
	return "黑色"
}
func getPosition(number int) int {
	rand.Seed(time.Now().UnixNano() + int64(number))
	return rand.Intn(18)
}
