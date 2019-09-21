package Registration

import (
	dbmanager "Super-market/Database"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var err error

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func Homepage(w http.ResponseWriter, req *http.Request) {
	fmt.Println("HOMEPAGE HANDLER ")
	http.ServeFile(w, req, "Registration/homepage.html")
}

func SignUp(w http.ResponseWriter, req *http.Request) {
	fmt.Println("what is r.method ", req.Method)
	if req.Method != "POST" {
		fmt.Println("serving html")
		http.ServeFile(w, req, "Registration/signup.html")

	}
	username := req.FormValue("username")
	password := req.FormValue("password")
	fmt.Println("username received is", username)
	fmt.Println("password received is", password)
	var user string

	_, err = dbmanager.Db.Exec("USE userlogindata")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Selected the Database Successfully!")
	}

	rows, err := dbmanager.Db.Query("SELECT username FROM users WHERE username=?", username)
	if err != nil {
		panic(err.Error())
	}
	count := 0
	for rows.Next() {
		count++
	}
	if count > 0 {
		w.Write([]byte("USER already existed"))
		return
	}
	err = dbmanager.Db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)
	switch {
	case err == sql.ErrNoRows:

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Unable to crate accout right now (ONE)", http.StatusInternalServerError)
			return
		}
		_, err = dbmanager.Db.Exec("INSERT INTO users(username,password) VALUES(?,?)", username, hashedPassword)
		if err != nil {
			http.Error(w, "Unable to create your account (TWO)", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("User Accout Created"))

	case err != nil:
		http.Error(w, "Unable to create your account (THREE)", http.StatusInternalServerError)
		return
	default:
		http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
	}

	http.Redirect(w, req, "/login", http.StatusOK)
}

func Login(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(w, req, "Registration/login.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var dbusername string
	var dbpassword string

	err = dbmanager.Db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&dbusername, &dbpassword)
	if err != nil {
		w.Write([]byte("User Does Not Exist"))
		http.Redirect(w, req, "/login", http.StatusTemporaryRedirect)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbpassword), []byte(password))
	if err != nil {
		w.Write([]byte("Credentials Did not Match!"))
		http.Redirect(w, req, "/login", http.StatusTemporaryRedirect)
		return
	}

	//create a new session with random token
	sessionToken := uuid.NewV4().String()
	// Set the token in the cache, along with the user whom it represents
	// The token has an expiry time of 120 seconds
	fmt.Println("session token is ", sessionToken)
	_, err = dbmanager.Cache.Do("SETEX", sessionToken, "300", username)
	fmt.Println("Error is", err)
	if err != nil {
		fmt.Println("err occured")
		// If there is an error in setting the cache, return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds, the same as the cache
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(300 * time.Second),
	})

	fmt.Println("DONE!")
	//w.Write([]byte("Log In Successful"))
	http.Redirect(w, req, "/inventory", 301)
}
