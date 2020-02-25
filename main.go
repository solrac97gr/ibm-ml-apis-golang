package main

import (
	"log"

	"github.com/joho/godotenv"
	natural "github.com/solrac97gr/ibm-cloud-course/naturallenguajeclassifier"
	personality "github.com/solrac97gr/ibm-cloud-course/personalityInsights"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	personality.Personality()
	natural.Natural()
}
