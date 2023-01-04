package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"time"
)

type Entry struct {
	StartTime float64 `json:"startTime"`
	Title string `json:"title"`
}

type JSONFile struct {
	Version   string   `json:"version"`
	Chapters  []Entry `json:"chapters"`
}


func main() {
	// Check if a filename was provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as the first argument.")
		os.Exit(1)
	}

	// Open the file
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()


	var entries []Entry

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into the time and the text
		parts := strings.SplitN(scanner.Text(), " ", 2)
		if len(parts) != 2 {
			fmt.Printf("Skipping invalid line: %v\n", scanner.Text())
			continue
		}

		// Parse the time
		t, err := time.Parse("15:04:05.000", parts[0])
		if err != nil {
			fmt.Printf("Skipping invalid line: %v\n", scanner.Text())
			continue
		}

		// Convert the time to seconds
		milliseconds := float64(t.Hour())*3600.0 + float64(t.Minute())*60.0 + float64(t.Second()) + float64(t.Nanosecond())/1000000000.0

		// Create a new entry and append it to the slice
		entry := Entry{StartTime: milliseconds, Title: parts[1]}
		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}


	jsonFile := &JSONFile{
		Version:  "1.2.0",
		Chapters: entries,
	}

	outputFilename := filename + ".json"
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	b, err := json.MarshalIndent(jsonFile, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n\n", b)

	_, err = outputFile.Write(b)
	if err != nil {
		panic(err)
	}
	fmt.Println("All done\n\n")
}






