package Dispatcher

import (
	"fmt"
	"testing"
)

func TestHandleRoute(t *testing.T) {
	p := &Person{
		Name: "pp",
	}
	d := NewDispatcher()
	d.Regist("ppmoudle", p)
	ret, err := d.HandleByRoute("ppmoudle.run", "22")
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("ret:", ret)
	}
	t.Error("d")

}

type Person struct {
	Name string
}

func (this *Person) Run(a int) string {
	fmt.Println("a is ", a)
	return "ll"
}
