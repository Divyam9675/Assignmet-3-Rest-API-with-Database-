package main

import (
	"api/customer_api"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/customer/findall", customer_api.FindAll).Methods("GET")
	router.HandleFunc("/api/customer/findall/{keyword}", customer_api.Search).Methods("GET")
	router.HandleFunc("/api/customer/Creates", customer_api.Create).Methods("POST")
	router.HandleFunc("/api/customer/Update", customer_api.Update).Methods("PUT")
	router.HandleFunc("/api/customer/Delete/{keyword}", customer_api.Delete).Methods("DELETE")

	err := http.ListenAndServe(":9000", router)

	if err != nil {
		fmt.Println(err)
	}

}
