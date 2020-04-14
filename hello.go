package main

import (
	"fmt"
)

func main() {
	x := make(map[string]string)

	x["key"] = "hehe"
	fmt.Println(x["jj"])

	name, ok := x["Al"]
	fmt.Println(name, ok)
}
