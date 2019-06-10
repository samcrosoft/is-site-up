package main

import (
	"fmt"
	"os"
	"time"

	S "github.com/samcrosoft/is-site-up"
)

func getCliArgs() string {

	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s <URL> \n", os.Args[0])
		os.Exit(0)
	}

	url := os.Args[1]
	return url
}

func main() {
	url := getCliArgs()
	timeOutInSeconds := 30
	timeout := time.Duration(timeOutInSeconds) * time.Second

	isUp, response := S.CheckSite(url, timeout)

	statusMsg := "Down"
	if isUp == true {
		statusMsg = "Up"
	}
	message := fmt.Sprintf("Site - [%s] Is %s", url, statusMsg)
	fmt.Println(message)
	if isUp == true {
		fmt.Printf("Status : %s \n", response.Status)
		fmt.Printf("Status Code : %d \n", response.StatusCode)
	}

}
