package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const delay = 1

func main() {
	showIntroduction()

	for {
		showMenu()
		option := readOption()
		processOption(option)
	}
}

func showIntroduction() {
	userName := "Denilson"
	fmt.Println("========================================")
	fmt.Println("Hello", userName, "!Welcome to our site monitor. Choose some options below:")
	fmt.Println("========================================")
}

func showMenu() {
	fmt.Println("Choose an option:")
	fmt.Println()
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

	sites := getSitesToBeMonitored()
	sites = append(sites, "http://www.github.com")

	showLogsBeforeTestSites(sites)
	processMonitoring(sites)

}

func showLogsBeforeTestSites(sites []string) {
	fmt.Println("Total sites to be monitored:", len(sites))
	fmt.Println("The slice capacity is:", cap(sites))
	fmt.Println("Sites to be monitored:", sites)

	for i, site := range sites {
		fmt.Println("Site", i, ":", site)
	}
}

func getSitesToBeMonitored() []string {
	// return []string{
	// 	"http://www.alura.com.br",
	// 	"http://www.caelum.com.br",
	// 	"http://www.google.com",
	// 	"http://www.uol.com.br",
	// 	"http://www.terra.com.br",
	// }

	return readSitesFromFile()
}

func readSitesFromFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("An error occurred while trying to read the file:", err)
	}
	// file, err := ioutil.ReadFile("sites.txt")
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()

	//fmt.Println(string(file))
	//fmt.Println(file)

	return sites
}

func processMonitoring(sites []string) {
	var responseHttpStatusCodce bool

	for i := 0; i < len(sites); i++ {
		fmt.Println("========================================")
		response, err := http.Get(sites[i])
		if err != nil {
			fmt.Println("An error occurred while trying to access the site:", err)
		}
		responseHttpStatusCodce = response.StatusCode == 200
		fmt.Println("Response status code:", response.StatusCode)

		if responseHttpStatusCodce {
			fmt.Println("Site", sites[i], "is up!")
		} else {
			fmt.Println("Site ", sites[i], "is down!")
		}
		time.Sleep(delay * time.Second)
		fmt.Println("========================================")
	}
}
