// https://blog.golang.org/laws-of-reflection
package main

import "fmt"
import "reflect"

type T struct {
	A int
	B string
}

func main() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(x))
	fmt.Println(p.Interface())
	v := p.Elem()
	v.SetFloat(5.5)
	fmt.Println(v)
	fmt.Println(x)

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(typeOfT)
	fmt.Println(s.Interface())

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("s is now", s)
	fmt.Println("t is now", t)

	// float64
	// 3.4
	// 0xc42000e270
	// 5.5
	// 5.5

	// {23 skidoo}
	// {23 skidoo}
	// main.T
	// {23 skidoo}

	// 0: A int = 23
	// 1: B string = skidoo
	// s is now {77 Sunset Strip}
	// t is now {77 Sunset Strip}
}
