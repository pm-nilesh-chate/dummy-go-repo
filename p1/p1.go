package p1

import "fmt"

func Foo() {
	fmt.Println("p1.Foo()!!! 1")
	fmt.Println("p1.Foo()!!! 2")
	fmt.Println("p1.Foo()!!! 3")
	if true {
		return
	}
	fmt.Println("p1.Foo()!!! 4")
}
