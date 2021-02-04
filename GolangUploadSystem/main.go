package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func HandleError(e error){
	if e != nil{
		fmt.Println(e)
	}
}
func Upload(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("myFile")

	HandleError(err)

	defer file.Close()
	fmt.Print("The file name is: ", handler.Filename, "\n")
	fmt.Print("The file size is: ", handler.Size, "\n")

	fileBytes, err:= ioutil.ReadAll(file)

	HandleError(err)

	err = ioutil.WriteFile(string(handler.Filename),fileBytes,0644)

	HandleError(err)

}

func main() {
	http.HandleFunc("/upload", Upload)
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":8080", nil)
}
