package main

import (
	"fmt"
	"strconv"
)

//var block
var (
	d = 34
	e = "antonio"
)

func main() {

	//useful for initialisation before you want to use variable
	var i int
	i = 42
	fmt.Println(i)

	//useful to control types on the fly
	var j int = 29
	fmt.Println(j)

	//short and sweet
	var k = 38
	fmt.Println(k)

	//can shadow variables in different scopes
	d := "my"
	fmt.Println(d, e)

	//converting to string with strconv
	f := strconv.Itoa(k)
	fmt.Println(f)
}
