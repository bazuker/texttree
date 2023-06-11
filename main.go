package main

import (
	"fmt"
	"texttree/v1"
)

func main() {
	tt, err := texttree.NewTextTree("sample/", texttree.DefaultMaxFileSize)
	if err != nil {
		panic(err)
	}
	fmt.Println(tt.Entities())
	fmt.Println(tt.SubExists("b"))
	fmt.Println(tt.GetIfExists("b/hello.txt"))
	fmt.Println(tt.GetStringSubIfExists("b/c", "quack"))
}
