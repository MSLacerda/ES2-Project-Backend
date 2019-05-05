package main

import (
	"github.com/MSLacerda/ES2-Project-Backend/handle"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	public := mux.NewRouter()

	public.HandleFunc("/api/casos", handle.ListarCasos).Methods("GET")
	public.HandleFunc("/api/casos/{id:[0-9]+}", handle.BuscarCaso).Methods("GET")
	public.HandleFunc("/api/casos/{id:[0-9]+}", handle.ConferirCaso).Methods("PUT")

	public.HandleFunc("/api/estorias", handle.ListarEstorias).Methods("GET")
	public.HandleFunc("/api/estorias/{id:[0-9]+}", handle.BuscarEstoria).Methods("GET")
	public.HandleFunc("/api/estorias/{id:[0-9]+}", handle.ConferirEstoria).Methods("GET")

	return public
}