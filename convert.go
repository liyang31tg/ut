package ut

import (
	"strconv"
	"math/big"
	"crypto/rand"
	"fmt"
)

func ToInt64(s interface{}) int64 {
	switch v := s.(type) {
	case int64:
		return v
	case string:
		i,e:=strconv.ParseInt(v, 10, 64)
		if e!=nil{
			return 0
		}
		return i
	case int:
		return int64(v)
	}
	return 0
}

func ToInt(i interface{}) int {
	switch v := i.(type) {
	case int:
		return v
	case string:
		intV, err := strconv.Atoi(v)
		if err != nil {
			return 0
		} else {
			return intV
		}
	case int64:

		return int(v)
	}
	return 0
}

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
