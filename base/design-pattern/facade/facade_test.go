package facade

import (
	"strconv"
	"testing"
)

func TestNewDeduplication(t *testing.T) {
	funcCheck := NewDeduplication()
	funcPut := funcCheck.Put(strconv.Atoi)
	t.Logf("funcPut: %v", funcPut)
	funcExist := funcCheck.Exist(strconv.Atoi)
	t.Logf("funcExist: %v", funcExist)
	funcDelete := funcCheck.Delete(strconv.Atoi)
	t.Logf("funcDelete: %v", funcDelete)

	stringCheck := NewDeduplication()
	stringPut := stringCheck.Put("qq")
	t.Logf("stringPut: %v", stringPut)
	stringExist := stringCheck.Exist("qq")
	t.Logf("stringExist: %v", stringExist)
	stringDelete := stringCheck.Delete("qq")
	t.Logf("stringDelete: %v", stringDelete)
	stringDelete2 := stringCheck.Delete("qq")
	t.Logf("stringDelete2: %v", stringDelete2)

}
