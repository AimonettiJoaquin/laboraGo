package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Definimos estrctura para el mapeo de la info

type Gasto struct {
	ID       string  `json:"id"`
	Concepto string  `json:"concepto"`
	Importe  float64 `json:"importe"`
	Fecha    string  `json:"fecha"`
}

// Simular una db

var gastos []Gasto

func getGastos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gastos)
}

func getGastoById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	for _, item := range gastos {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// Si no se encuentra el gasto
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&Gasto{})
}

func createGasto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var gasto Gasto
	_ = json.NewDecoder(r.Body).Decode(&gasto)
	gastos = append(gastos, gasto)
	json.NewEncoder(w).Encode(gasto)
}

func updateGasto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range gastos {
		if item.ID == params["id"] {
			gastos = append(gastos[:index], gastos[index+1:]...)
			var gasto Gasto
			_ = json.NewDecoder(r.Body).Decode(&gasto)
			gasto.ID = params["id"]
			gastos = append(gastos, gasto)
			json.NewEncoder(w).Encode(gasto)
			return
		}
	}
	// Si no se encuentra el gasto
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(gastos)
}

func deleteGasto(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	for index, item := range gastos {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			gastos = append(gastos[:index], gastos[index+1:]...)
			json.NewEncoder(w).Encode(gastos)
			break
		}
	}
	// Si no se encuentra el gasto
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&Gasto{})
}

// Funcion main

// Inicializar el server

func main() {
	// Inicializar algunso datos
	gastos = append(gastos, Gasto{ID: "1", Concepto: "compra de articulos", Importe: 2500, Fecha: "2018-10-10"})
	gastos = append(gastos, Gasto{ID: "2", Concepto: "compra de alimentos", Importe: 250, Fecha: "2018-11-10"})
	r := mux.NewRouter()

	r.HandleFunc("/gastos", getGastos).Methods("GET")
	r.HandleFunc("/gastos/{id}", getGastoById).Methods("GET")
	r.HandleFunc("/gastos", createGasto).Methods("POST")
	// r.HandleFunc("/gastos/{id}", updateGasto).Methods("PUT")
	// r.HandleFunc("/gastos/{id}", deleteGasto).Methods("DELETE")

	// Iniciamos el server
	log.Println("Server iniciado en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
