package main

import (
	"github.com/joho/godotenv"
	"note-app/infrastructure/router"
	"note-app/infrastructure/db"
	r "note-app/repository"
)

func init() {
	godotenv.Load()
}

func main() {
	r.Repository = db.GetConnection()
	router.GetRouter().Run()
}
