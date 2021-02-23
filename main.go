package main

import (
	"os"
	"github.com/joho/godotenv"
	"url_shortener/src/url"
)

import "log"

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	config := ConfigFromFile(os.Getenv("CONFIG_PATH"))
	a := *NewApp(&config)
	url.NewUrlController(&a)
	a.Run()
	log.Println("lol")
}