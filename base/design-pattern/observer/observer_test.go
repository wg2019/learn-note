package observer

import "testing"

func TestNewYork_Update(t *testing.T) {
	york := GetNewYork()
	policeStation := GetPoliceStation(york)
	educationBureau := GetEducationBureau(york)
	york.Update("第一道命令")

	t.Log(policeStation.GetContent())
	t.Log(educationBureau.GetContent())
	york.Update("第二道命令")

	t.Log(policeStation.GetContent())
	t.Log(educationBureau.GetContent())
	york.Update("第三道命令")

	t.Log(policeStation.GetContent())
	t.Log(educationBureau.GetContent())
}
