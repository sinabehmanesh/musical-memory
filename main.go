package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello lads")

}

func main() {
	fmt.Printf("web server starting....")

	http.HandleFunc("/hello", hello)

	fmt.Printf("web server started.")
	http.ListenAndServe(":8080", nil)
}
