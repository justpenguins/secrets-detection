package scanner

import (
    "regexp"
)

var SecretPatterns = map[string]*regexp.Regexp{
    "API Key":        regexp.MustCompile(`(?i)(api[-_]?key=)([a-zA-Z0-9]{5,32})`), // Matches API keys in URLs
    "AWS Secret Key": regexp.MustCompile(`(?i)aws[-_]?secret[-_]?key\s*[:=]\s*([a-zA-Z0-9/+=]{40})`),
}