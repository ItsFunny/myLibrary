package utils

import (
	"fmt"
	"testing"
)

func TestCombineOrConditionSQL(t *testing.T) {
	// bools := make([]bool, 3)
	// names := make([]string, 3)
	// oragDepth, oragName, oragCode := "depth", "name", "code"
	// bools[0], bools[1], bools[2] = true, false, true
	// names[0], names[1], names[2] = "DEPTH", "NAME", "CODE"
	// vagueSql := CombineOrConditionSQL("OR", bools, names, &oragDepth, &oragName, &oragCode)
	// fmt.Println(vagueSql)

	str := "qwe"
	test(str)
	// change(&str)
	// fmt.Println(str)
}
func test(a interface{}) {
	s := a.(string)
	s += "%"
	fmt.Println(s)
}
func change(s interface{}) {
	i := s.(*string)
	*i += "kkkk"
}
