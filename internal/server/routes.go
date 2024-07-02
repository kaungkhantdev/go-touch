package server

import (
	"fmt"
	"net/http"
)

func TestFun(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func TestRoute() {
	http.HandleFunc("/test", TestFun);
}