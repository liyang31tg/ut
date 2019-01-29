package ut

import "reflect"

func Array2Map(arr interface{}, mapkey func(item interface{}) string) map[string]interface{} {
	m := make(map[string]interface{}, 0)
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("to slice arr not slice")
	}
	l := v.Len()
	for i := 0; i < l; i++ {
		tmpArr := v.Index(i).Interface()
		m[mapkey(tmpArr)] = tmpArr
	}
	return m
}
