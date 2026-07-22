package main

import (
	"fmt"
)

func main() {
	var c = make(map[string]string)
	c["year"] = "A"
	c["brand"] = "B"
	c["model"] = "C"
	delete(c, "model")
	fmt.Println(c)

	var a = map[string]string{"A": "1", "B": "2", "C": "3"}
	var b = a

	fmt.Println(a)
	fmt.Println(b)
	b["year"] = "12312"

	fmt.Println(a)
	fmt.Println(b)
}
