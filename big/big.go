// https://golang.org/pkg/math/big/#example_Int_Scan
package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	// The Scan function is rarely used directly;
	// the fmt package recognizes it as an implementation of fmt.Scanner.
	i := new(big.Int)
	_, err := fmt.Sscan("0x5b5ff4eeba3614f594f40b2a", i)
	if err != nil {
		log.Println("error scanning value:", err)
	} else {
		fmt.Println(i)
	}
	is := fmt.Sprintf("%s", i)
	fmt.Println(is)
}
