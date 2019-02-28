package Dispatcher

import (
	"errors"
	"reflect"
	"strings"
	"sync"

	"github.com/liyang31tg/ut"
)

type dispatcher struct {
	handlers map[string]reflect.Value
	mtx      sync.Mutex
}

func NewDispatcher() *dispatcher {
	return &dispatcher{
		handlers: map[string]reflect.Value{},
	}
}

func (this *dispatcher) Regist(module string, value interface{}) {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	this.handlers[module] = reflect.ValueOf(value)
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
		err = errors.New("cb is not valid")
		return
	}
	cbargs := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		cbargs[i] = reflect.ValueOf(args[i])
	}
	res = cb.Call(cbargs)
	return
}

func (this *dispatcher) getFunc(module, function string) reflect.Value {
	if m, ok := this.handlers[module]; ok {
		return m.MethodByName(ut.ToCapitalize(function))
	} else {
		return reflect.Value{}
	}
}
