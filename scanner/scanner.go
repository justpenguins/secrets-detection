package scanner

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type Finding struct {
	File    string
	LineNum int
	Match   string
	Type    string
}

func ScanDirectory(path string) ([]Finding, error) {
	var findings []Finding
	err := filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
		if info.IsDir() || shouldSkip(file) {
			return nil
		}
		fileFindings, err := scanFile(file)
		if err == nil {
			findings = append(findings, fileFindings...)
		}
		return nil
	})
	return findings, err
}

func shouldSkip(path string) bool {
	skipDirs := []string{".git", "node_modules", "vendor"}
	for _, d := range skipDirs {
		if filepath.Base(path) == d {
			return true
		}
	}
	return false
}

func scanFile(path string) ([]Finding, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var findings []Finding
	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		for name, pattern := range SecretPatterns {
			if pattern.MatchString(line) {
				findings = append(findings, Finding{
					File:    path,
					LineNum: lineNum,
					Match:   pattern.FindString(line),
					Type:    name,
				})
			}
		}
		lineNum++
	}
	return findings, scanner.Err()
}
