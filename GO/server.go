package main

import (
	"net/http"
	"fmt"
	"sync"
)
var mutex = &sync.Mutex{}

var actual_id int

type CurrentJson struct {
	Id int `json`
	Content string `json`
}

type ShoutOut struct {
	Content string `json:"content"`
	Id int `json`
}

func handler(writer http.ResponseWriter, request *http.Request) {
	mutex.Lock()
 	actual_id = actual_id + 1
	mutex.Unlock()
	fmt.Fprintf(writer, "Hello World, %s!", actual_id , request.URL.Path[1:])
}



func main() {
	actual_id = 0;
	http.HandleFunc("/home", handler)
	http.HandleFunc("/api", handler)
	http.ListenAndServe(":8077", nil)
}
