package p1

import "fmt"

func Foo() {
	fmt.Println("p1.Foo()!!! - fork 1")
	fmt.Println("p1.Foo()!!! - fork 2")
	if true {
		return
	}
	fmt.Println("p1.Foo()!!! - fork 3")
	fmt.Println("p1.Foo()!!! - fork 4")
}
