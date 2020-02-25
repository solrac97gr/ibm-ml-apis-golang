package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/solrac97gr/ibm-cloud-course/discovery"
	natural "github.com/solrac97gr/ibm-cloud-course/naturallenguageclassifier"
	naturalu "github.com/solrac97gr/ibm-cloud-course/naturallenguageunderstanding"
	personality "github.com/solrac97gr/ibm-cloud-course/personalityInsights"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Select the api you need to test the success conection")
	fmt.Println("1.Personality Insights")
	fmt.Println("2.Natural Lenguage Classifier")
	fmt.Println("3.Natural Lenguage Understanding")
	fmt.Println("4.Discovery")
	fmt.Println("5.Exit")
	option, _ := reader.ReadString('\n')
	option = strings.TrimSuffix(option, "\n")
	optionint, err := strconv.Atoi(option)
	if err != nil {
		panic("Error in convert option")
	}

	if optionint == 1 {
		fmt.Println("entre")
		personality.Personality()
	}
	if optionint == 2 {
		natural.Natural("Como estas")
	}
	if optionint == 3 {
		naturalu.Naturalu("www.carlosgrowth.com")
	}
	if optionint == 4 {
		discovery.Discovery()
	}
	if optionint == 5 {
		fmt.Println("Thanks for use")
	}

}
