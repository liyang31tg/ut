package Dispatcher

import (
	"errors"
	"fmt"
	"testing"
)

func TestHandleRoute(t *testing.T) {
	p := &Person{
		Name: "pp",
	}
	d := NewDispatcher()
	d.Regist(Person{Name: "liyang"})
	err := d.Regist(p)

	if err != nil {

	}
	t.Error(errors.New("asdf"))

	ret, err := d.HandleByRoute("person.run2", 22)
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

func (this *Person) run(a int) string {
	fmt.Println("a is ", a)
	return "ll"
}

func (this Person) run2(a int) string {
	fmt.Println("a is ", a)
	return this.Name
}
