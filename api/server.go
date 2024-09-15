package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ncephamz/efishery-be-test/api/controllers"
)

var server = controllers.Server{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("STEIN_HQ_URL"), os.Getenv("STEIN_HQ_USERNAME"), os.Getenv("STEIN_HQ_PASSWORD"))

	server.Run(":8080")

}
