package main

import (
	"log"
	todo "todo-app"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
