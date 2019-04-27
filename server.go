package main

import (
	"net/http"

	"github.com/yusufpapurcu/GoDatabaseExample/database"

	"github.com/yusufpapurcu/GoDatabaseExample/route"
)

func main() {
	database.Connect("mongodb://localhost:27017") //MongoDB Connection Started
	router := route.SetRouter()                   // Api Router's Set
	http.ListenAndServe(":8080", router)          // Start Server
}
