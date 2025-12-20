package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "root", "database username")
	password := flag.String("password", "root", "database password")
	host := flag.String("host", "localhost", "database host")
	port := flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println("username:", *username)
	fmt.Println("password:", *password)
	fmt.Println("host:", *host)
	fmt.Println("port:", *port)
}
