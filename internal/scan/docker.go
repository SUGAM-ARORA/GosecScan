package scan

import (
	"fmt"
	"os/exec"
)

func ScanContainers(image string) {
	fmt.Printf("🔍 Scanning image: %s\n", image)
	
	// Check if Trivy is installed
	if !checkTrivy() {
		fmt.Println("❌ Error: Trivy is not installed. Please install Trivy first.")
		return
	}

	// Run Trivy scan
	cmd := exec.Command("trivy", "image",
		"--severity", "HIGH,CRITICAL",
		"--format", "table",
		image)
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Error scanning image: %v\n", err)
		return
	}

	fmt.Println(string(output))
}

func checkTrivy() bool {
	cmd := exec.Command("trivy", "--version")
	return cmd.Run() == nil
}