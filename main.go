package main

import (
	"go-jwt/database"
	"go-jwt/router"
)

func main() {
	database.StartDB()
	r := router.StrartApp()
	r.Run(":8080")
}
