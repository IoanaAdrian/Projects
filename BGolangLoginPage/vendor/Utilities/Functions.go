package Utilities

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func CreateProfile(response http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)
	if userName == "" {
		http.Redirect(response, request, "/", 302)
	}

	FName := request.FormValue("Fname")
	LName := request.FormValue("Lname")
	Country := request.FormValue("Country")
	City := request.FormValue("City")
	Telephone := request.FormValue("Telephone")
	Adress := request.FormValue("Adress")
	db, err := sql.Open("mysql", "hDwSeFYfZa:5LKhaYraEq@tcp(remotemysql.com:3306)/hDwSeFYfZa")
	if err != nil {
		panic(err.Error())
	}
	query := "INSERT IGNORE INTO cp (`FName`, `LName`, `Country`, `City`, `Telephone`, `Adress`, `username`) VALUES (?,?, ?, ?, ?, ?, ?)"
	if userName != "" {
		insert, err := db.Query(query, FName, LName, Country, City, Telephone, Adress, userName)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		defer db.Close()
	}
	bdy, err := ioutil.ReadFile("hpages/indexx.html")
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(response, string(bdy), userName, FName, LName, Country, City, Telephone, Adress)

}

func Logout(response http.ResponseWriter, request *http.Request) {

	clearSession(response)
	http.Redirect(response, request, "/", 302)

}
func Indexx(response http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)
	if userName == "" {
		http.Redirect(response, request, "/", 302)
	}
	db, err := sql.Open("mysql", "hDwSeFYfZa:5LKhaYraEq@tcp(remotemysql.com:3306)/hDwSeFYfZa")
	rows, err := db.Query("SELECT * FROM `godb`")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var username string
		var ppassword string
		var profilecreatedd int
		err = rows.Scan(&username, &ppassword, &profilecreatedd)

		if err != nil {
			panic(err.Error())
		}
		if userName == username {
			if profilecreatedd == 0 {

				q := "UPDATE godb SET profilecreated='1' WHERE username=?"
				if err != nil {
					panic(err.Error())
				}
				insert, err := db.Query(q, userName)
				if err != nil {
					panic(err.Error())
				}
				defer insert.Close()
				defer db.Close()
				defer rows.Close()
				http.Redirect(response, request, "/CProfile.html", 302)

				//http.Redirect(response, request, "/indexx", 302)
			} else {
				rows, err := db.Query("SELECT * FROM `cp`")
				defer db.Close()
				if err != nil {
					panic(err.Error())
				}
				for rows.Next() {
					var FName string
					var LName string
					var Country string
					var City string
					var Telephone string
					var Adress string
					var username string
					err = rows.Scan(&FName, &LName, &Country, &City, &Telephone, &Adress, &username)
					if username == userName {
						defer db.Close()
						defer rows.Close()
						bdy, err := ioutil.ReadFile("hpages/indexx.html")
						if err != nil {
							panic(err.Error())
						}

						fmt.Fprintf(response, string(bdy), userName, FName, LName, Country, City, Telephone, Adress)
					}
				}
				defer db.Close()
				defer rows.Close()
			}
			defer db.Close()
			defer rows.Close()
			break
		}

	}
	defer db.Close()
	defer rows.Close()

}

func Login(response http.ResponseWriter, request *http.Request) {

	name := request.FormValue("name")
	password := request.FormValue("password")
	redirectT := "/register.html"
	if len(name) > 0 && len(password) > 0 {
		db, err := sql.Open("mysql", "hDwSeFYfZa:5LKhaYraEq@tcp(remotemysql.com:3306)/hDwSeFYfZa")
		rows, err := db.Query("SELECT * FROM `godb`")

		if err != nil {
			panic(err.Error())
		}

		for rows.Next() {
			var username string
			var ppassword string
			var profilecreatedd int
			err = rows.Scan(&username, &ppassword, &profilecreatedd)

			if err != nil {
				panic(err.Error())
			}
			if name == username && password == ppassword {
				redirectT = "/indexx"
				setSession(name, response)
				break
			}

		}
		defer db.Close()
		defer rows.Close()
	}
	http.Redirect(response, request, redirectT, 302)

}

func Register(response http.ResponseWriter, request *http.Request) {

	request.ParseForm()

	name := request.FormValue("username")
	email := request.FormValue("email")
	pass := request.FormValue("password")
	confirmpass := request.FormValue("confirmPassword")

	if len(name) > 0 && len(email) > 0 && len(pass) > 0 && len(confirmpass) > 0 && pass == confirmpass {
		db, err := sql.Open("mysql", "hDwSeFYfZa:5LKhaYraEq@tcp(remotemysql.com:3306)/hDwSeFYfZa")
		if err != nil {
			panic(err.Error())
		}
		query := "INSERT IGNORE INTO godb (`Username`, `Password`, `profilecreated`) VALUES (?, ?, ?)"

		var pc int = 0
		insert, err := db.Query(query, name, pass, pc)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		defer db.Close()
		http.Redirect(response, request, "/", 302)

	} else {
		http.Redirect(response, request, "/register.html", 302)

	}

}
func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func EditFirstName(response http.ResponseWriter, request *http.Request) {

	userName := getUserName(request)
	if userName == "" {
		http.Redirect(response, request, "/", 302)
	}

	editpart := request.FormValue("EditH")
	fmt.Println(editpart)
	db, err := sql.Open("mysql", "hDwSeFYfZa:5LKhaYraEq@tcp(remotemysql.com:3306)/hDwSeFYfZa")
	if err != nil {
		panic(err.Error())
	}
	q := "UPDATE cp SET FName=? WHERE username=?"
	if err != nil {
		panic(err.Error())
	}
	insert, err := db.Query(q, editpart, userName)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	defer db.Close()
	http.Redirect(response, request, "/logout", 302)

}
func EditLastName(response http.ResponseWriter, request *http.Request) {

	userName := getUserName(request)
	if userName == "" {
		http.Redirect(response, request, "/", 302)
	}

	editpart := request.FormValue("EditL")
	fmt.Println(editpart)
	db, err := sql.Open("mysql", "hDwSeFYfZa:5LKhaYraEq@tcp(remotemysql.com:3306)/hDwSeFYfZa")
	if err != nil {
		panic(err.Error())
	}
	q := "UPDATE cp SET LName=? WHERE username=?"
	if err != nil {
		panic(err.Error())
	}
	insert, err := db.Query(q, editpart, userName)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	defer db.Close()
	http.Redirect(response, request, "/logout", 302)

}
func EditCountry(response http.ResponseWriter, request *http.Request) {

	userName := getUserName(request)
	if userName == "" {
		http.Redirect(response, request, "/", 302)
	}

	editpart := request.FormValue("EditC")
	fmt.Println(editpart)
	db, err := sql.Open("mysql", "hDwSeFYfZa:5LKhaYraEq@tcp(remotemysql.com:3306)/hDwSeFYfZa")
	if err != nil {
		panic(err.Error())
	}
	q := "UPDATE cp SET Country=? WHERE username=?"
	if err != nil {
		panic(err.Error())
	}
	insert, err := db.Query(q, editpart, userName)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	defer db.Close()
	http.Redirect(response, request, "/logout", 302)

}
func EditCity(response http.ResponseWriter, request *http.Request) {

	userName := getUserName(request)
	if userName == "" {
		http.Redirect(response, request, "/", 302)
	}

	editpart := request.FormValue("EditCi")
	fmt.Println(editpart)
	db, err := sql.Open("mysql", "hDwSeFYfZa:5LKhaYraEq@tcp(remotemysql.com:3306)/hDwSeFYfZa")
	if err != nil {
		panic(err.Error())
	}
	q := "UPDATE cp SET City=? WHERE username=?"
	if err != nil {
		panic(err.Error())
	}
	insert, err := db.Query(q, editpart, userName)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	defer db.Close()
	http.Redirect(response, request, "/logout", 302)

}
func EditTelephone(response http.ResponseWriter, request *http.Request) {

	userName := getUserName(request)
	if userName == "" {
		http.Redirect(response, request, "/", 302)
	}

	editpart := request.FormValue("EditT")
	fmt.Println(editpart)
	db, err := sql.Open("mysql", "hDwSeFYfZa:5LKhaYraEq@tcp(remotemysql.com:3306)/hDwSeFYfZa")
	if err != nil {
		panic(err.Error())
	}
	q := "UPDATE cp SET Telephone=? WHERE username=?"
	if err != nil {
		panic(err.Error())
	}
	insert, err := db.Query(q, editpart, userName)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	defer db.Close()
	http.Redirect(response, request, "/logout", 302)

}
func EditAdress(response http.ResponseWriter, request *http.Request) {

	userName := getUserName(request)
	if userName == "" {
		http.Redirect(response, request, "/", 302)
	}

	editpart := request.FormValue("EditA")
	fmt.Println(editpart)
	db, err := sql.Open("mysql", "hDwSeFYfZa:5LKhaYraEq@tcp(remotemysql.com:3306)/hDwSeFYfZa")
	if err != nil {
		panic(err.Error())
	}
	q := "UPDATE cp SET Adress=? WHERE username=?"
	if err != nil {
		panic(err.Error())
	}
	insert, err := db.Query(q, editpart, userName)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	defer db.Close()
	http.Redirect(response, request, "/logout", 302)

}
