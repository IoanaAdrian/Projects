package main

import (
	"Utilities"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.Handle("/", http.FileServer(http.Dir("./hpages")))
	http.HandleFunc("/login", Utilities.Login)
	http.HandleFunc("/editfirstname", Utilities.EditFirstName)
	http.HandleFunc("/editlastname", Utilities.EditLastName)
	http.HandleFunc("/editcountry", Utilities.EditCountry)
	http.HandleFunc("/editcity", Utilities.EditCity)
	http.HandleFunc("/edittelephone", Utilities.EditTelephone)
	http.HandleFunc("/editadress", Utilities.EditAdress)
	http.HandleFunc("/createprofile", Utilities.CreateProfile)
	http.HandleFunc("/indexx", Utilities.Indexx) //not letting the user acces the profile without logging in
	//http.HandleFunc("/indexx.html", Utilities.Indexx) //not letting the user acces the profile without logging in
	http.HandleFunc("/logout", Utilities.Logout)
	http.HandleFunc("/register", Utilities.Register)
	http.ListenAndServe(":"+port, nil)
	//http.ListenAndServe(":8080", nil)

}
