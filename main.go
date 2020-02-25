package main

import (
	"log"

	"github.com/joho/godotenv"
	natural "github.com/solrac97gr/ibm-cloud-course/naturallenguajeclassifier"
	naturalu "github.com/solrac97gr/ibm-cloud-course/naturallenguajeunderstanding"
	personality "github.com/solrac97gr/ibm-cloud-course/personalityInsights"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	personality.Personality()
	natural.Natural("Como estas")
	naturalu.Naturalu("www.carlosgrowth.com")
}
