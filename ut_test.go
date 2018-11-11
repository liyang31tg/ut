package ut

import (
	"fmt"
	"testing"
)

func TestRandStr(t *testing.T) {
	b := Int64ToByte(2222222223344424444)
	fmt.Println(b)
	i := ByteToInt64(b)
	fmt.Println(i)

}

func TestMapArr2Map(t *testing.T) {
	mapA := []map[string]string{map[string]string{"name": "liyang", "age": "28"}, map[string]string{"name": "limengxia", "age": "29"}}
	fmt.Println(MapArr2Map(mapA, "name"))
}
