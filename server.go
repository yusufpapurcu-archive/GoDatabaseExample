package main

import (
	"net/http"

	"github.com/YTPSourceCode/DatabaseTest/database"

	"github.com/YTPSourceCode/DatabaseTest/route"
)

func main() {
	database.Connect("mongodb://localhost:27017") //MongoDB Connection Started
	router := route.SetRouter()                   // Api Router's Set
	http.ListenAndServe(":8080", router)          // Start Server
}
