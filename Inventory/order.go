package Inventory

import (
	"net/http"
)

func ReceiveOrderpage(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "Inventory/menubar.html")
}

func VegPizza(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "Inventory/vegpizaa.html")
}

func NonvegPizza(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "Inventory/non-vegpizza.html")
}
