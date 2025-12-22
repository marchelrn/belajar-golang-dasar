package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.Join("hello", "world", ".go"))
	fmt.Println(filepath.IsLocal("hello/world.go"))
	fmt.Println(path.IsAbs("hello/world.go"))
}
