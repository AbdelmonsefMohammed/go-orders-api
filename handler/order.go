package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all Orders")
}
func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Order")
}

func (o *Order) Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an Order by id")
}

func (o *Order) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an Order by id")
}

func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an Order by id")
}
