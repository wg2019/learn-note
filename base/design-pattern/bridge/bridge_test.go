package bridge

import "testing"

func TestNewToy(t *testing.T) {
	toy1 := NewToy()
	toy1.SetColor(new(White))
	toy1.SetShape(new(Circle))
	t.Logf("toy1: %s", toy1.Draw())
	toy2 := NewToy()
	toy2.SetColor(new(Black))
	toy2.SetShape(new(Square))
	t.Logf("toy2: %s", toy2.Draw())
}
