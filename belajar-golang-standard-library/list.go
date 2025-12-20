package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack(1)
	data.PushBack(2)
	data.PushBack(3)

	head := data.Front()
	number := head.Next()
	fmt.Println(head.Value)
	fmt.Println(number.Value)

	for n := data.Front(); n != nil; n = n.Next() {
		if n.Value == 3 {
			break
		}
		fmt.Println(n.Value)
	}
}
