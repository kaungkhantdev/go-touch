package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func NewServer() {
	TestRoute();

	port := os.Getenv("PORT")
	fmt.Println("GO is running on: http://localhost"+port)
	log.Fatal(http.ListenAndServe(port, nil))
}