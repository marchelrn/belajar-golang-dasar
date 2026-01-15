package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestSting(t *testing.T) {
	fmt.Println(version)
}

//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("Newlogo.png", logo, fs.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

//go:embed ../files/file_a.txt
//go:embed files/file_b.txt
//go:embed files/file_c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	files.ReadFile("fles/file_a.txt")
	a, _ := files.ReadFile("files/file_a.txt")
	fmt.Println(string(a))

	files.ReadFile("fles/file_b.txt")
	b, _ := files.ReadFile("files/file_b.txt")
	fmt.Println(string(b))

	files.ReadFile("fles/file_c.txt")
	c, _ := files.ReadFile("files/file_c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var txtFiles embed.FS

func TestPathMatcher(t *testing.T) {
	dir, _ := txtFiles.ReadDir("files")
	for _, file := range dir {
		if !file.IsDir() {
			fmt.Println(file.Name())
			content, _ := txtFiles.ReadFile("files/" + file.Name())
			fmt.Println(string(content))
		}
	}
}
