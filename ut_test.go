package ut

import (
	"testing"
	"fmt"
)

func TestRandStr(t *testing.T)  {
	b:=Int64ToByte(2222222223344424444)
	fmt.Println(b)
	i:=ByteToInt64(b)
	fmt.Println(i)

}

func TestMd5FileReader(t *testing.T) {
	}

