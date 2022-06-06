package main

import (
	"fmt"
)

func main() {

	//boolean
	var n bool = false
	fmt.Printf("%v, %T\n", n, n)

	m := 1 == 1
	fmt.Printf("%v, %T\n", m, m)

	//default value is always 0. in bool it is "false"
	var k bool
	fmt.Printf("%v, %T\n", k, k)

}
