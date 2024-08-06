package main

import(
	"io"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", echoHello)
	log.Println("start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoHello(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Hello World\n")
}