package main

import (
	"gowithtests/di"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(di.MyGreeterHandler)))

}
