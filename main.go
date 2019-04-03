package main

import (
	"fmt"
	"texttree/v1"
)

func main() {
	tt, err := texttree.NewTextTree("sample", texttree.DefaultMaxFileSize)
	if err != nil {
		panic(err)
	}
	fmt.Println(tt.Entities())
	fmt.Println(tt.GetString("a/hello"))
	fmt.Println(tt.Get("b/hello.txt").Content)
	fmt.Println(tt.GetStringSub("b", "c/quack"))
}
