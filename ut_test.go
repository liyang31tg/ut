package ut

import (
	"testing"
	"fmt"
	"strings"
)

func TestRandStr(t *testing.T)  {
	b:=Int64ToByte(2222222223344424444)
	fmt.Println(b)
	i:=ByteToInt64(b)
	fmt.Println(i)

}

func TestMd5FileReader(t *testing.T) {
	s := strings.NewReader("I amqwrqwrerwrwerwerwerwerwer a")
		fmt.Println(Md5FileReader(s))
	}
