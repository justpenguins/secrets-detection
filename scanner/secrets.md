# Gitleaks+

## Overview
*Gitleaks+** is a Go-based application designed to scan directories and files for sensitive information such as API keys, AWS secret keys, and other secrets. It identifies potential security vulnerabilities by detecting hardcoded secrets in your codebase and provides a report summarizing the findings.

---

## Features
- **Recursive Directory Scanning**: Scans a specified directory and all its subdirectories for secrets.
- **Pattern Matching**: Detects secrets using customizable regex patterns (e.g., API keys, AWS secret keys).
- **Secrets Report**: Generates a detailed report of detected secrets, including their file location and remediation status.
- **Remediation Tracking**: Marks secrets as remediated when addressed.

---

## How It Works
1. **Directory Scanning**: The application starts at a specified directory and recursively scans all files.
2. **Pattern Matching**: Each file is scanned line by line, and regex patterns are applied to detect secrets.
3. **Report Generation**: A summary report is generated, showing the total number of secrets detected, remediated, and pending.

---

## Motivation
Hardcoded secrets in codebases are a common security vulnerability that can lead to data breaches and unauthorized access. This tool helps developers identify and address such issues early in the development lifecycle, improving the overall security posture of their applications.

---

## Usage

### Prerequisites
- **Go**: Ensure Go is installed on your system. You can download it from [golang.org](https://golang.org/).
- **AWS Credentials**: If you plan to use the AWS key rotation feature, ensure your AWS credentials are configured (e.g., via `~/.aws/credentials` or environment variables).

### Steps to Run
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd secrets