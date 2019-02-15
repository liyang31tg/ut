package ut

import (
	"encoding/json"
	"strconv"

	"bytes"
	"encoding/binary"
)

func ToInt64(s interface{}) int64 {
	switch v := s.(type) {
	case int64:
		return v
	case string:
		i, e := strconv.ParseInt(v, 10, 64)
		if e != nil {
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

func ToString(i interface{}) string {
	switch v := i.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	case int64:
		return ToString(int(v))
	}
	return ""
}

func ToBool(s interface{}) bool {
	switch v := s.(type) {
	case int:
		if v == 0 {
			return false
		} else {
			return true
		}
	case string:
		b, er := strconv.ParseBool(v)
		if er != nil {
			return false
		} else {
			return b
		}
	}
	return false
}

func ToFloat64(s interface{}) float64 {
	switch v := s.(type) {
	case int:
		return float64(v)
	case string:
		r, e := strconv.ParseFloat(v, 64)
		if e != nil {
			return 0
		}
		return r
	case float32:
		return float64(v)
	case float64:
		return v
	}
	return 0
}

func ToFloat32(s interface{}) float32 {
	switch v := s.(type) {
	case int:
		return float32(v)
	case string:
		r, e := strconv.ParseFloat(v, 32)
		if e != nil {
			return 0
		}
		return ToFloat32(r)
	case float32:
		return v
	case float64:
		return float32(v)
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
	binary.Read(buffer, binary.BigEndian, &i)
	return i
}

/*
根据一个字段，将切片变成map
*/
func MapArr2Map(mapArr []map[string]string, key string) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range mapArr {
		result[m[key]] = m
	}
	return result
}

func Byte2Map(content []byte) map[string]interface{} {
	var m = map[string]interface{}{}
	err := json.Unmarshal(content, &m)
	if err != nil {
		return map[string]interface{}{}
	} else {
		return m
	}

}
