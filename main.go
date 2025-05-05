package main

import (
	"fmt"
	"os"

	"secrets/scanner"
	"secrets/utils"
)

func main() {
	findings, err := scan.ScanDirectory(".")
	if err != nil {
		fmt.Println("Scan error:", err)
		os.Exit(1)
	}

	if len(findings) == 0 {
		fmt.Println("✅ No secrets found.")
		return
	}

	for _, f := range findings {
		fmt.Printf("❌ %s:%d [%s] %s\n", f.File, f.LineNum, f.Type, f.Match)

		// Handle AWS keys
        if f.Type == "AWSAccessKey" {
            fmt.Println("🔑 Rotating AWS key:", f.Match)
            utils.Rotate(f.Match)
        } else {
            // Placeholder for future remediation logic
            fmt.Println("⚠️ Non-AWS secret detected. Remediation required.")
        }
	}

	os.Exit(1)
}
