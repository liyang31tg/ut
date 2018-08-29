package ut

import (
	"testing"
	"fmt"
)

func TestRandStr(t *testing.T)  {
	str := RandomStr(127)
	fmt.Println(str)
}
