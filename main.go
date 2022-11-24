package main

import (
	"note-app/infrastructure"
)

func main() {
  infrastructure.GetRouter().Run()
}