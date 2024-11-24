package main

import (
	"fmt"
	dotenv "github.com/joho/godotenv"
)

func main() {
	err := dotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	

	
}