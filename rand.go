package ut

import (
	"math/big"
	"crypto/rand"
	"fmt"
)

//rand num is in [0,max)
func Rand(max int) int {
	if max <= 0 {
		panic("rand value not less equal zero")
	}
	v := ToInt64(max)
	n, e := rand.Int(rand.Reader, big.NewInt(v))
	if e != nil {
		fmt.Println(e)
		return 0
	}
	return ToInt(n.Int64())
}

//rand num is in [minNum,maxNum]
func RandomInterval(minNum int, maxNum int) int {
	return Rand(maxNum-minNum+1) + minNum
}

func RandomStr(length int) string{
	b := make([]byte,length)
	for i:=0;i<length;i++{
		b[i]=byte(Rand(127)) //ascii
	}
	return string(b)
}