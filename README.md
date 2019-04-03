[![Go Report Card](https://goreportcard.com/badge/github.com/kisulken/texttree)](https://goreportcard.com/report/github.com/kisulken/texttree)

# texttree
TextTree is a file buffer that stores files content in memory and allows access to it by path. It is useful for working with localization trees.

## Install
```Bash
go get -u github.com/kisulken/texttree
```

## Example
```Go
tt, err := texttree.NewTextTree("sample", texttree.DefaultMaxFileSize)
if err != nil {
	panic(err)
}
fmt.Println(tt.Entities())
fmt.Println(tt.GetString("a/hello"))
fmt.Println(tt.GetString("b/hello.txt"))
fmt.Println(tt.GetString("b/c/quack"))
```
