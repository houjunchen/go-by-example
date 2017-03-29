// https://github.com/golang/go/wiki/SliceTricks
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{7, 8, 9}

	// AppendVector
	a = append(a, b...)
	fmt.Println("a: ", a)

	// Copy
	c := make([]int, len(a))
	copy(c, a)
	c = append([]int(nil), a...)
	fmt.Println("a: ", a)
	fmt.Println("c: ", c)

	// Cut
	d := append(a[:3], a[5:]...)
	fmt.Println("a: ", a)
	fmt.Println("d: ", d)

	// Delete
	e := append(a[:3], a[4:]...)
	fmt.Println("a: ", a)
	fmt.Println("e: ", e)

	// Expand
	f := append(a[:3], append(make([]int, 5), a[3:]...)...)
	fmt.Println("a: ", a)
	fmt.Println("f: ", f)

	// Extend
	g := append(a, make([]int, 6)...)
	fmt.Println("a: ", a)
	fmt.Println("g: ", g)

	// Insert
	h := append(a[:5], append([]int{6}, a[5:]...)...)
	fmt.Println("h: ", h)

	// InsertVector
	i := append(a[:3], append(b, a[3:]...)...)
	fmt.Println("i: ", i)

	// Pop
	x, a := a[len(a)-1], a[:len(a)-1]
	fmt.Println("a: ", a)
	fmt.Println("x: ", x)

	// Push
	a = append(a, 11)
	fmt.Println("a: ", a)

	// Shift
	y, a := a[0], a[1:]
	fmt.Println("y: ", y)
	fmt.Println("a: ", a)

	// Unshift
	a = append([]int{1}, a...)
	fmt.Println("a: ", a)
}
