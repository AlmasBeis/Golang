package main

import (
	"fmt"
	"io"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request")
	io.WriteString(w, "My first string\n")
}
func main() {
	http.HandleFunc("/", GetRoot)
	fmt.Println("server is listening..")
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		fmt.Print(err)
		return
	}
}
