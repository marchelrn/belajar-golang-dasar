package main

import "fmt"

func main() {

	names := [...]string{
		"Marchel",
		"Bintang",
		"Rizky",
		"Fadilla",
		"Joko",
	}

	slice1 := names[3:5]
	fmt.Println(slice1, "==> Slice 1")

	slice2 := names[:]
	fmt.Println(slice2, "==> Slice 2")

	slice3 := names[:4]
	fmt.Println(slice3, "==> Slice 3")

	slice4 := names[2:]
	fmt.Println(slice4, "==> Slice 4")

	// LENGTH

	fmt.Println(len(slice1))
	fmt.Println(len(slice2))
	fmt.Println(len(names), "==> Panjang Array") // Panjang Array

	// CAPACITY

	fmt.Println(cap(slice1))
	fmt.Println(cap(slice2))
	fmt.Println(cap(names), "==> Kapasitas Array") // Kapasitas Array

	// CREATE SLICE

	crSlice := make([]string, 2, 5)
	fmt.Println(crSlice, "==> Create Slice")
	crSlice[0] = "Marchel"
	crSlice[1] = "Bintang"
	fmt.Println(crSlice)

	//APPEND SLICE

	appndSlice := append(crSlice, "Rizky", "Fadilla", "Joko")
	fmt.Println(appndSlice, "==> Append from new Slice")

	appndSlice2 := append(slice1, "Budi", "Sukma")
	fmt.Println(appndSlice2, "==> Append from Slice 1")

	//COPY SLICE

	copySlice := make([]string, len(appndSlice), cap(appndSlice))
	copy(copySlice, appndSlice)
	fmt.Println(copySlice, "==> Copy from new Slice")

	copySlice2 := make([]string, len(slice1), cap(slice1))
	copy(copySlice2, slice1)
	fmt.Println(copySlice2, "==> Copy from Array")
}
