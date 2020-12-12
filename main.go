package main

import (
	"log"
	"net/http"
	"pizza-shop/Inventory"
	"pizza-shop/dbmanager"
	"pizza-shop/order-placement"
	"pizza-shop/registration"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const hashConst = 8

type Credentials struct {
	Username string `json:"Username,db:"username"`
	Password string `json:"Password,db:"password"`
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", registration.Homepage)
	r.HandleFunc("/signup", registration.SignUp)
	r.HandleFunc("/login", registration.Login)
	r.HandleFunc("/inventory", Inventory.ReceiveOrderpage)
	r.HandleFunc("/inventory/VegPizza", Inventory.VegPizza)
	r.HandleFunc("/inventory/NonVegPizza", Inventory.NonvegPizza)
	r.HandleFunc("/order/{id}", order.Ordervegpizza)
	r.HandleFunc("/order/{id}", order.Ordernonvegpizza)

	dbmanager.InitDB()
	log.Println("Starting http server.....")
	log.Fatal(http.ListenAndServe(":9595", r))

}
