package ut

import (
	"strconv"

	"bytes"
	"encoding/binary"
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

func Int64ToByte(num int64) []byte {
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, num)
	return buffer.Bytes()
}

func ByteToInt64(data []byte) int64 {
	buffer := bytes.NewBuffer(data)
	var i int64
	binary.Read(buffer,binary.BigEndian,&i)
	return i
}


