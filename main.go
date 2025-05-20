package main

import (
	"fmt"
	"os"

	"secrets/scanner"
)

func main() {
	findings, err := scanner.ScanDirectory("./test/test.go")
	if err != nil {
		fmt.Println("Scan error:", err)
		os.Exit(1)
	}

	// Example: Mark AWS keys as remediated
	for i, f := range findings {
		if f.Type == "AWSAccessKey" {
			findings[i].Remediated = true
		}
	}

	// Generate and print the report
	report := scanner.GenerateReport(findings)
	fmt.Println(report)
}
