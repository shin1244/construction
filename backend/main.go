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
	{Name: "Company B", Year: 2022},
}

func GetCompanyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(companyList)
}

func PostCompanyHandler(w http.ResponseWriter, r *http.Request) {
	var newCompany Company
	err := json.NewDecoder(r.Body).Decode(&newCompany)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	companyList = append(companyList, newCompany)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newCompany)
}

func router(router *mux.Router) {

	router.HandleFunc("/companies", GetCompanyHandler).Methods("GET")
	router.HandleFunc("/companies", PostCompanyHandler).Methods("POST")

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
