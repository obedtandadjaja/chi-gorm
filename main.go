package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := InitRouter()

	fmt.Println("Now serving server at :8000")
	http.ListenAndServe(":8000", r)
}
