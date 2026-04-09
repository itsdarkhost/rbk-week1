package main

import (
	"fmt"
	"os"

	"github.com/itsdarkhost/rbk-week1/internal/service"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input> <output>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	tc := service.TextConverter{
		Input: string(data),
	}

	result := tc.Convert()

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Printf("File saved at: %s\n", outputFile)
		return
	}

	fmt.Println("File saved at: ", outputFile)
}
