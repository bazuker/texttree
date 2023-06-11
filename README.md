[![Go Report Card](https://goreportcard.com/badge/github.com/bazuker/texttree)](https://goreportcard.com/report/github.com/bazuker/texttree)

# texttree
TextTree is a file buffer that stores files content in memory and allows access to it by path. It is useful for working with localization trees.

## Install
```Bash
go get -u github.com/kisulken/texttree
```

## Example
```Go
import "github.com/bazuker/texttree/v1"

...

tt, err := texttree.NewTextTree("sample", texttree.DefaultMaxFileSize) // buffers all the files here
if err != nil {
	panic(err)
}
fmt.Println(tt.Entities())
fmt.Println(tt.GetString("a/hello"))
fmt.Println(tt.Get("b/hello.txt").Content)
fmt.Println(tt.GetStringSub("b", "c/quack"))
```
