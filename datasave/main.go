package main

import (
	"datasave/model"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/upload",model.Upload)
	http.HandleFunc("downlaod",model.Download)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d",Host,Port),nil))
}