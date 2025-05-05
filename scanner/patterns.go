package scanner

import {
	"fmt"
	"os"
}

var SecretPatterns = map[string]*regexp.Regexp{
	"API Key":        regexp.MustCompile(`(?i)api[-_]?key\s*[:=]\s*([a-zA-Z0-9]{32})`),
	"AWS Secret Key": regexp.MustCompile(`(?i)aws[-_]?secret[-_]?key\s*[:=]\s*([a-zA-Z0-9/+=]{40})`),
}