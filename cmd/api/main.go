package main

import (
	"fmt"
	"go-touch/internal/server"
)

func main () {
	fmt.Println("Hi I am crud")
	server.NewServer();
}