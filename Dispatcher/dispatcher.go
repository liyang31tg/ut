package Dispatcher

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

/*
先是模块，然后在根据模块获得方法(导出的方法)，再缓存方法，避免每次调用的查找方法
*/
type dispatcher struct {
	handlers map[string]map[string]reflect.Value
}

func NewDispatcher(values ...interface{}) *dispatcher {
	d := &dispatcher{
		handlers: map[string]map[string]reflect.Value{},
	}
	for _, v := range values {
		d.Regist(v)
	}
	return d
}

func (d *dispatcher) Regist(value interface{}) error {
	return d.RegistByName("", value)
}
func (d *dispatcher) RegistByName(module string, value interface{}) error {
	v := reflect.ValueOf(value)
	vk := v.Kind()
	if vk == reflect.Struct || (vk == reflect.Ptr && v.Elem().Kind() == reflect.Struct) {
		if module == "" {
			if vk == reflect.Ptr {
				module = strings.ToLower(v.Elem().Type().Name())
			} else {
				module = strings.ToLower(v.Type().Name())
			}
		}
		d.setMethod(module, v)
		return nil
	} else {
		return errors.New("模块类型必须是结构体，或者结构体的指针")
	}
}

func (d *dispatcher) setMethod(module string, v reflect.Value) {
	vt := v.Type()
	mc := vt.NumMethod() //只能获得导出的方法
	var c map[string]reflect.Value
	if moduleV, ok := d.handlers[module]; ok {
		c = moduleV
	} else {
		c = map[string]reflect.Value{}
		d.handlers[module] = c
	}
	for i := 0; i < mc; i++ {
		mn := strings.ToLower(vt.Method(i).Name)
		mt := v.Method(i)
		c[mn] = mt
	}
}

func (this *dispatcher) HandleByRoute(route string, arges ...interface{}) (res []reflect.Value, err error) {
	routes := strings.Split(route, ".")
	if len(routes) != 2 {
		err = errors.New("routes is not XX.YYY")
		return
	}
	return this.handle(routes[0], routes[1], arges...)
}

func (this *dispatcher) handle(module, function string, args ...interface{}) (res []reflect.Value, err error) {
	cb := this.getFunc(module, function)
	if !cb.IsValid() {
		err = fmt.Errorf("%v.%v is not valid", module, function)
		return
	}
	len := len(args)
	cbargs := make([]reflect.Value, len)
	for i := 0; i < len; i++ {
		cbargs[i] = reflect.ValueOf(args[i])
	}
	res = cb.Call(cbargs)
	return
}

func (this *dispatcher) getFunc(module, function string) reflect.Value {
	if m, ok := this.handlers[module]; ok {
		if f, ok := m[function]; ok {
			return f
		}
	}
	return reflect.Value{}
}
