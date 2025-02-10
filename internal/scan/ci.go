package scan

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func ScanGitHubActions() {
	fmt.Println("🔍 Scanning GitHub Actions workflows...")

	// Patterns to detect secrets
	patterns := map[string]*regexp.Regexp{
		"AWS Key":    regexp.MustCompile(`(?i)(aws_access_key_id|aws_secret_access_key)`),
		"Token":      regexp.MustCompile(`(?i)(api_token|api_key|auth_token|secret_token)`),
		"Password":   regexp.MustCompile(`(?i)(password|passwd|pwd)`),
		"Private Key": regexp.MustCompile(`(?i)(private_key|-----BEGIN [A-Z ]+ PRIVATE KEY-----)`),
	}

	// Scan .github/workflows directory
	err := filepath.Walk(".github/workflows", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && (filepath.Ext(path) == ".yml" || filepath.Ext(path) == ".yaml") {
			scanWorkflowFile(path, patterns)
		}
		return nil
	})

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("ℹ️  No GitHub Actions workflows found.")
		} else {
			fmt.Printf("❌ Error scanning workflows: %v\n", err)
		}
	}
}

func scanWorkflowFile(path string, patterns map[string]*regexp.Regexp) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("❌ Error reading file %s: %v\n", path, err)
		return
	}

	for patternName, pattern := range patterns {
		if pattern.Match(content) {
			fmt.Printf("🚨 Warning: Potential %s found in %s\n", patternName, path)
		}
	}
}