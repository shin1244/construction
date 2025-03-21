package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Company struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

var companyList = []Company{
	{Name: "Company A", Year: 2020},
}

func GetCompanyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(companyList)
	fmt.Fprint(w, string(data))
}

func router(router *mux.Router) {

	router.HandleFunc("/companies", GetCompanyHandler).Methods("GET")

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
