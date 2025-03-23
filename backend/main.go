package main

import (
	"fmt"
	"net/http"
	"tidy/backend/handlers"
)

func main() {
	m := handlers.MakeHandler()

	fmt.Println("Started Server")
	err := http.ListenAndServe(":3000", m)
	if err != nil {
		panic(err)
	}
}
