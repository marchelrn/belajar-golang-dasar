package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	age  int
}

type userSlice []User

func (s userSlice) Len() int {
	fmt.Println("len dipanggil")
	return len(s)
}

func (s userSlice) Less(i, j int) bool {
	fmt.Println("less dipanggil")
	return s[i].age < s[j].age
}

func (s userSlice) Swap(i, j int) {
	fmt.Println("swap dipanggil")
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{Name: "Alice", age: 30},
		{Name: "Bob", age: 25},
		{Name: "Charlie", age: 35},
	}

	sort.Sort(userSlice(users))
	fmt.Println(users)
}
