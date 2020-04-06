package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for {
		runCommand(menu())
	}
}

// Menu
func menu() int {
	fmt.Println("COMMANDS")
	fmt.Println("  1 - Healthcheck [Sandbox]")
	fmt.Println("  2 - Healthcheck [Production]")
	fmt.Println("  0 - End program")

	return getCommand()
}

func getCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println()
	return command
}

func runCommand(command int) {
	switch command {
	case 1:
		fmt.Println("Running: Healthcheck [Sandbox]")
		healthcheckSandbox()
	case 2:
		fmt.Println("Running: Healthcheck [Production]")
		healthcheckProduction()
	case 0:
		os.Exit(1)
	default:
		fmt.Println("Unknown command")
	}

	fmt.Println()
}

// Commands
func healthcheckSandbox() {
	urls := read("urls.txt")
	runHealthchecks(urls)
}

func healthcheckProduction() {
	urls := read("urls.txt")
	runHealthchecks(urlsRemoveSandbox(urls))
}

// Functions
func urlsRemoveSandbox(urls []string) []string {
	for i, url := range urls {
		urls[i] = strings.Replace(url, "sandbox", "", 1)
	}
	return urls
}

func runHealthchecks(urls []string) {
	for _, url := range urls {
		check := healtcheck(url)
		if check {
			fmt.Println("[ OK ] " + url)
		} else {
			fmt.Println("[FAIL] " + url)
		}
	}
}

// File Management
func read(filename string) []string {
	var content []string

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error reading file '"+filename+"':", err)
		return content
	}

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading file '"+filename+"':", err)
			return content
		}
		content = append(content, strings.TrimSpace(string(line)))
	}

	return content
}

// Healthcheck
type healtcheckResponse struct {
	IsHealthy bool
}

func healtcheck(url string) bool {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error calling '"+url+"':", err)
		return false
	}

	if response.StatusCode != 200 {
		return false
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error reading '"+url+"' response:", err)
		return false
	}

	var bodyObject healtcheckResponse
	json.Unmarshal(body, &bodyObject)

	return bodyObject.IsHealthy
}
