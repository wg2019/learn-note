package composite

import "testing"

func TestNewProvince(t *testing.T) {
	province := NewProvince("黑龙江")
	heb := NewCity("哈尔滨")
	heb.AddComposite(NewRegion("呼兰区"))
	heb.AddComposite(NewRegion("南岗区"))
	heb.AddComposite(NewRegion("香坊区"))
	province.AddComposite(heb)

	mdj := NewCity("牡丹江")
	mdj.AddComposite(NewRegion("爱民区"))
	mdj.AddComposite(NewRegion("东安区"))
	mdj.AddComposite(NewRegion("阳明区"))
	province.AddComposite(mdj)

	province.Features()
}
