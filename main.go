package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var config = Config{}

func main() {
	// Open the log file for writing
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	defer logFile.Close()
	readConfig(&config)

	// Start the server
	fmt.Printf("Starting server on port %s", config.Addr)
	log.Printf("Starting server on port %s", config.Addr)
	router := server()
	log.Fatal(http.ListenAndServe(config.Addr, router))
}

func readConfig(config *Config) {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// Read the file
	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
}
