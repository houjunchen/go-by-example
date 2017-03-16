package main

import "fmt"

type persion struct {
	name string
	age  int
}

func main() {
	fmt.Println(persion{"Bob", 20})
	fmt.Println(persion{name: "Alice", age: 30})
	fmt.Println(persion{name: "Fred"})
	fmt.Println(&persion{name: "Ann", age: 40})

	s := persion{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)
}
