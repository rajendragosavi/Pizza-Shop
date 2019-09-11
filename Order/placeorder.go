package Order

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var index string

type order struct {
	Id    string
	Name  string
	Price string
}

var tmpl = template.Must(template.ParseGlob("Order/*"))

func Ordervegpizza(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	index = params["id"]
	if index == "Margherit" {
		n := "Margherit"
		pizzoder := order{Id: index, Name: n, Price: "350 INR"}
		tmpl.ExecuteTemplate(w, "order.html", pizzoder)
	} else if index == "Peppy_Paneer" {
		n := "Peppy_Paneer"
		pizzoder := order{Id: index, Name: n, Price: "450 INR"}
		tmpl.ExecuteTemplate(w, "order.html", pizzoder)
	} else {
		n := "Farmhouse"
		pizzoder := order{Id: index, Name: n, Price: "550 INR"}
		tmpl.ExecuteTemplate(w, "order.html", pizzoder)
	}

}

/*
 in order struct we dont need to have name field , Id is enogh to do whatever we wan to do

*/

func Ordernonvegpizza(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	index = params["id"]

	if index == "Chicken_Sausage" {
		n := "Chicken_Sausage"
		pizzoder := order{Id: index, Name: n, Price: "350 INR"}
		tmpl.ExecuteTemplate(w, "order.html", pizzoder)
	} else if index == "Chicken_Golden_Delight" {
		n := "Chicken_Golden_Delight"
		pizzoder := order{Id: index, Name: n, Price: "450 INR"}
		tmpl.ExecuteTemplate(w, "order.html", pizzoder)
	} else {
		n := "Pepper_Barbeque"
		pizzoder := order{Id: index, Name: n, Price: "550 INR"}
		tmpl.ExecuteTemplate(w, "order.html", pizzoder)
	}
}
