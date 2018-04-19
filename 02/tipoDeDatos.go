package main

import "fmt"

 func iniciar(){
	var a int
	var b int64

	a = 12
	b = 5

	//casting

	c := a + int(b)
	fmt.Println(c)
}