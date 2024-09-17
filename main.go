package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	funcs "github.com/abhijaynair/metaheuristics-go/functions"
)

/*
Returns the status of the server.
** Only for Testing **å
*/
func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server healthy")
}

/*
Returns a list of available algorithms. Referred from the book by Xin-She Yang
More algorithms will be added to this list as they are implemented.
Refer to /testfunctions to see the list
of available test functions.
The response is formatted as

	{
		"algorithms":[]string
	}
*/
func algorithms(w http.ResponseWriter, r *http.Request) {

	type Algorithms struct {
		Algorithms []string `json:"algorithms"`
	}

	file, err := os.Open("algorithms.txt")
	if err != nil {
		return
	}

	// Close the file stream once returned from the function
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Initialize the response struct
	response := Algorithms{}

	// Add available algorithms to the response list
	// by reading from the file "algorithms.txt"
	for scanner.Scan() {
		line := scanner.Text()
		response.Algorithms = append(response.Algorithms, line)
	}

	// Set content type to json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseJson, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Could not encode json response")
		return
	}

	// Write the json to the http writer stream
	w.Write(responseJson)
}

func testFunctions(w http.ResponseWriter, r *http.Request){
	type Functions struct {
		Functions []string `json:"test_functions"`
	}
	
	file, err := os.Open("functions.txt")
	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScaner(file)
	response := Functions{}

	for scanner.Scan() {
		line := scanner.Text()
		response.Functions = append(response.Functions, line)	
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseJson, err := json.Marshall(response)
	if err != nil {
		fmt.Println("Could not encode json response")
		return
	}
	w.Write(responseJson)

}

func test(w http.ResponseWriter, r *http.Request) {
	funcs.Test()
}

func main() {

	// Start the Server
	server := http.NewServeMux()

	// setup request Handlers
	server.HandleFunc("/status", status)
	server.HandleFunc("/test", test)
	server.HandleFunc("/algorithms", algorithms)
	server.HandleFunc("/testfunctions", testFunctions)
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Println("Error starting server. Exiting...")
	}
}
