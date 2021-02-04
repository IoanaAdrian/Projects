package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	// "time"
)

var a string
var b string

func main() {
	// url := "https://www.worldometers.info/coronavirus//"
	// // fmt.Printf("HTML code of %s ...\n", url)
	// tr := &http.Transport{
	// 	MaxIdleConns:       10,
	// 	IdleConnTimeout:    30 * time.Second,
	// 	DisableCompression: true,
	// }
	// client := &http.Client{Transport: tr}
	// resp, err := client.Get(url)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// // reads html as a slice of bytes
	// html, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// // show the HTML code as a string %s
	// htmll := string(html)
	// // fmt.Printf("%s\n", html)
	// fmt.Printf(htmll)

	http.Handle("/", http.FileServer(http.Dir("./hpages")))
	http.HandleFunc("/login", Login)
	http.HandleFunc("/userinfo", UserInfo)
	http.HandleFunc("/csv", HandleCsv)
	// csvReaderRow()
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
	//http.ListenAndServe(":8080", nil)
}
func HandleCsv(response http.ResponseWriter, request *http.Request) {
	http.Redirect(response, request, "/data.html", 302)
}
func Login(response http.ResponseWriter, request *http.Request) {
	username := request.FormValue("Username")
	password := request.FormValue("Password")
	a = username
	b = password
	if username == "s.hoarders@gmail.com" && password == "ro031" {
		http.Redirect(response, request, "/userinfo", 302)
		fmt.Println("Logged in")
	} else {
		fmt.Println(username)
		fmt.Println(password)
	}
}
func UserInfo(response http.ResponseWriter, request *http.Request) {
	if a != "s.hoarders@gmail.com" || b != "ro031" {
		http.Redirect(response, request, "/", 302)
	}
	bdy, err := ioutil.ReadFile("hpages/userinfo.html")
	if err != nil {
		panic(err.Error())
	}
	recordFile, err := os.Open("hpages/data.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return
	}

	// Setup the reader
	reader := csv.NewReader(recordFile)

	// Read the records
	header, err := reader.Read()
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return
	}
	fmt.Printf("Headers : %v \n", header)
	var x string
	var y string
	for i := 0; ; i = i + 1 {
		record, err := reader.Read()
		if err == io.EOF {
			break // reached end of the file
		} else if err != nil {
			fmt.Println("An error encountered ::", err)
			return
		}
		x = record[2]
		y = record[3]
	}
	fmt.Println(x)
	fmt.Println(y)
	fmt.Fprintf(response, string(bdy), "", "", x, y, "stabila")

}

// func csvReaderRow() {
// 	// Open the file
// 	recordFile, err := os.Open("hpages/data.csv")
// 	if err != nil {
// 		fmt.Println("An error encountered ::", err)
// 		return
// 	}

// 	// Setup the reader
// 	reader := csv.NewReader(recordFile)

// 	// Read the records
// 	header, err := reader.Read()
// 	if err != nil {
// 		fmt.Println("An error encountered ::", err)
// 		return
// 	}
// 	fmt.Printf("Headers : %v \n", header)

// 	for i := 0; ; i = i + 1 {
// 		record, err := reader.Read()
// 		if err == io.EOF {
// 			break // reached end of the file
// 		} else if err != nil {
// 			fmt.Println("An error encountered ::", err)
// 			return
// 		}
// 		fmt.Println(record[2])
// 		fmt.Println(record[3])
// 	}

// 	// Note: Each time Read() is called, it reads the next line from the file
// 	// r1, _ := reader.Read() // Reads the first row, useful for headers
// 	// r2, _ := reader.Read() // Reads the second row
// }
