package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	showIntroduction()

	for {
		showMenu()
		option := readOption()
		processOption(option)
	}

	/**
	if option == 1 {
		fmt.Println("Starting monitoring...")
	} else if option == 2 {
		fmt.Println("Showing logs...")
	} else if option == 0 {
		fmt.Println("Exiting...")
	} else {
		fmt.Println("Invalid option. Please try again.")
	}
		**/

}

func showIntroduction() {
	userName := "Denilson"
	fmt.Println("========================================")
	fmt.Println("Hello", userName, "!Welcome to our site monitor. Choose some options below:")
	fmt.Println("========================================")
}

func showMenu() {
	fmt.Println("1. Start Monitoring")
	fmt.Println("2. Show Logs")
	fmt.Println("0. Exit")

	fmt.Println("========================================")
}

func readOption() int {
	var option int
	fmt.Scanf("%d", &option)
	fmt.Println("You chose option:", option)
	return option
}

func processOption(option int) {
	switch option {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Showing logs...")
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid option. Please try again.")
		os.Exit(-1)
	}

}

func startMonitoring() {
	fmt.Println("Starting monitoring...")
	site := "http://www.alura.com.br"
	response, _ := http.Get(site)
	fmt.Println("Response status code:", response.StatusCode)

	responseHttpStatusCodce := response.StatusCode == 200

	if responseHttpStatusCodce {
		fmt.Println("Site", site, "is up!")
	} else {
		fmt.Println("Site ", site, "is down!")
	}
}
