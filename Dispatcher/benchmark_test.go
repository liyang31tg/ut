package Dispatcher

import (
	"testing"
)

type Person struct {
	Name string
}

func (this *Person) run(a int) string {
	return "ll"
}

func (this Person) Run2(a int) string {
	return this.Name
}

func Benchmark_HandleRoute(b *testing.B) {

	d := NewDispatcher(Person{Name: "liyang"})

	for i := 0; i < b.N; i++ {
		ret, err := d.HandleByRoute("person.run2", i)
		if err != nil {
			b.Error(err)
		}
		if ret[0].String() != "liyang" {

			b.Error("value is error")
		}

	}

}
