package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed logo.png
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)
	err := ioutil.WriteFile("Newlogo.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	dir, _ := path.ReadDir("files")
	for _, file := range dir {
		if !file.IsDir() {
			fmt.Println(file.Name())
			content, _ := path.ReadFile("files/" + file.Name())
			fmt.Println(string(content))
		}
	}
}
