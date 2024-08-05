package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)


func main() {
	// Define command-line flags
	inputFile := flag.String("input", "", "Input JSON file (optional)")
	outputFile := flag.String("output", "", "Output JSON file (optional)")
	flag.Parse()

	var input []byte
	var err error

	// Read input from file or standard input
	if *inputFile != "" {
		input, err = os.ReadFile(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
			os.Exit(1)
		}
	} else {
		input, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading standard input: %v\n", err)
			os.Exit(1)
		}
	}

	// Unmarshal JSON to check for validity
	var jsonData interface{}
	err = json.Unmarshal(input, &jsonData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid JSON: %v\n", err)
		os.Exit(1)
	}

	// Marshal JSON with indentation
	formattedJSON, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting JSON: %v\n", err)
		os.Exit(1)
	}


	// Write output to file or standard output
	if *outputFile != "" {
		err = os.WriteFile(*outputFile, formattedJSON, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println(string(formattedJSON))
	}
}