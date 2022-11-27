package main

import (
	"github.com/joho/godotenv"
	"note-app/infrastructure/db"
	"note-app/infrastructure/router"
	r "note-app/repository"
)

func init() {
	godotenv.Load()
}

func main() {
	r.Repository = db.GetConnection()
	router.GetRouter().Run()
}
