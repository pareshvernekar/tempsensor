package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	router:=NewRouter()
	fmt.Println("Starting Temperature Sensor API")
	log.Fatal(http.ListenAndServe(":5000", router))
}