package main

import (
	"log"
	"net/http"

	_ "shoppinglist/database"
)

func main() {
	err := http.ListenAndServe(":80", GenerateRouter())
	if err != nil {
		log.Panic(err.Error())
	}

}
