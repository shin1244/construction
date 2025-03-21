package main

import (
	"fmt"
	"net/http"
	"tidy/backend/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func router(router *mux.Router) {

	router.HandleFunc("/companies", handlers.GetCompanyHandler).Methods("GET")
	// router.HandleFunc("/companies", PostCompanyHandler).Methods("POST")

	fmt.Println("서버 시작")

	handler := cors.Default().Handler(router)
	err := http.ListenAndServe(":3000", handler)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	r := mux.NewRouter()

	router(r)
}
