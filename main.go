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
	fmt.Println(tt.GetString("sample/a/hello"))
	fmt.Println(tt.GetString("sample/b/hello.txt"))
	fmt.Println(tt.GetString("sample/b/c/quack"))
}
