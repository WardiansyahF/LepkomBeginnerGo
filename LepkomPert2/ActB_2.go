package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var f_name string = "Wardiansyah"
	var l_name string = "Fauzi"
	var age int = 19

	fmt.Println("My name is " + f_name + " " + l_name + "." + " I am " + strconv.Itoa(age) + " years old")
}