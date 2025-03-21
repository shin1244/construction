package handlers

import (
	"encoding/json"
	"net/http"
	"tidy/backend/models"
)

func GetCompanyHandler(w http.ResponseWriter, r *http.Request) {
	companies := models.GetAllCompanies()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(companies)
}

// func PostCompanyHandler(w http.ResponseWriter, r *http.Request) {
// 	var newCompany Company
// 	err := json.NewDecoder(r.Body).Decode(&newCompany)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	companyList = append(companyList, newCompany)

// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)

// 	json.NewEncoder(w).Encode(newCompany)
// }
