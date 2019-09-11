package main

import (
	dbmanagaer "Super-market/Database"
	"Super-market/Inventory"
	"Super-market/Order"
	"Super-market/Registration"
	"log"
	"net/http"

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
	r.HandleFunc("/", Registration.Homepage)
	r.HandleFunc("/signup", Registration.SignUp)
	r.HandleFunc("/login", Registration.Login)
	r.HandleFunc("/inventory", Inventory.ReceiveOrderpage)
	r.HandleFunc("/inventory/VegPizza", Inventory.VegPizza)
	r.HandleFunc("/inventory/NonVegPizza", Inventory.NonvegPizza)
	r.HandleFunc("/order/{id}", Order.Ordervegpizza)
	r.HandleFunc("/order/{id}", Order.Ordernonvegpizza)

	dbmanagaer.InitDB()
	log.Println("Starting http server.....")
	log.Fatal(http.ListenAndServe(":9595", r))

}
