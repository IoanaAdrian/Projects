package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"

	"os"

	_ "github.com/go-sql-driver/mysql"
)

var K string

func main() {
	http.Handle("/", http.FileServer(http.Dir("./hpages")))
	http.HandleFunc("/trimite", TrimiteF)
	http.HandleFunc("/arhiva", ArhivaF)
	db, err := sql.Open("mysql", "hDwSeFYfZa:5LKhaYraEq@tcp(remotemysql.com:3306)/hDwSeFYfZa")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("SELECT * FROM `buzidbuser`")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {

		var name string
		var pcr int
		err = rows.Scan(&name, &pcr)
		if err != nil {
			panic(err.Error())
		}
		http.HandleFunc(name, ArhivaUser)
	}
	defer db.Close()
	defer rows.Close()
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
	//http.ListenAndServe(":8080", nil)

}
func TrimiteF(response http.ResponseWriter, request *http.Request) {
	a := request.FormValue("CodeTitle")
	b := request.FormValue("CodeSourcee")
	db, err := sql.Open("mysql", "hDwSeFYfZa:5LKhaYraEq@tcp(remotemysql.com:3306)/hDwSeFYfZa")
	if err != nil {
		panic(err.Error())
	}
	query := "INSERT IGNORE INTO buzidb (`NumeSursa`, `Sursa`) VALUES (?,?)"
	insert, err := db.Query(query, a, b)
	if err != nil {
		panic(err.Error())
	}
	n := len(a)
	var s string = "/"
	for i := 0; i < n; i++ {
		if a[i] == '/' || a[i] == '\\' {
			break
		}
		s = s + string(a[i])
	}
	K = s
	defer insert.Close()
	rows, err := db.Query("SELECT * FROM `buzidbuser`")
	if err != nil {
		panic(err.Error())
	}
	var ok = 0
	for rows.Next() {

		var name string
		var source string
		err = rows.Scan(&name, &source)
		if err != nil {
			panic(err.Error())
		}
		if name == s {
			ok = 1
			break
		}

	}
	defer rows.Close()
	if ok == 0 {
		queryy := "INSERT IGNORE INTO buzidbuser (`NumeSursa`, `profilecreated`) VALUES (?,?)"
		insertt, err := db.Query(queryy, s, 0)
		if err != nil {
			panic(err.Error())
		}
		http.HandleFunc(string(s), ArhivaUser)
		defer insertt.Close()
	}
	defer db.Close()
	http.Redirect(response, request, "/arhiva", 302)

}
func ArhivaF(response http.ResponseWriter, request *http.Request) {
	db, err := sql.Open("mysql", "hDwSeFYfZa:5LKhaYraEq@tcp(remotemysql.com:3306)/hDwSeFYfZa")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("SELECT * FROM `buzidb`")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {

		var name string
		var source string
		err = rows.Scan(&name, &source)
		if err != nil {
			panic(err.Error())
		}
		bdy, err := ioutil.ReadFile("hpages/arhiva.html")
		if err != nil {
			panic(err.Error())
		}
		n := len(name)
		var s string = ""
		for i := 0; i < n; i++ {
			if name[i] == '/' || name[i] == '\\' {
				break
			}
			s = s + string(name[i])
		}
		fmt.Fprintf(response, string(bdy), s, name, source)

	}
	defer db.Close()
	defer rows.Close()

}
func ArhivaUser(response http.ResponseWriter, request *http.Request) {

	insf := request.URL.String()
	db, err := sql.Open("mysql", "hDwSeFYfZa:5LKhaYraEq@tcp(remotemysql.com:3306)/hDwSeFYfZa")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("SELECT * FROM `buzidb`")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {

		var name string
		var source string
		err = rows.Scan(&name, &source)
		if err != nil {
			panic(err.Error())
		}
		bdy, err := ioutil.ReadFile("hpages/arhiva.html")
		if err != nil {
			panic(err.Error())
		}
		n := len(name)
		var s string = "/"
		for i := 0; i < n; i++ {
			if name[i] == '/' || name[i] == '\\' {
				break
			}
			s = s + string(name[i])
		}
		if s == insf {
			fmt.Fprintf(response, string(bdy), s, name, source)
		}

	}
	defer db.Close()
	defer rows.Close()

}
