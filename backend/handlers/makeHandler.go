package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/unrolled/render"
)

var rd *render.Render

func MakeHandler() http.Handler {
	rd = render.New()
	r := mux.NewRouter()

	r.HandleFunc("/companies", getCompanyHandler).Methods("GET")
	r.HandleFunc("/companies", addCompanyHandler).Methods("POST")

	return cors.Default().Handler(r)
}
