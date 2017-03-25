package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	aInput := scanner.Text()
	aList := strings.Split(aInput, " ")
	var sum int64
	var a int64
	for i := 0; i < len(aList); i++ {
		a, _ = strconv.ParseInt(aList[i], 0, 64)
		sum += a
	}
	fmt.Println(sum)
}
