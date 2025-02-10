package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gosecscan",
	Short: "GosecScan - A comprehensive security scanner for Cloud & DevOps",
	Long: `GosecScan is a security scanner that helps identify vulnerabilities,
misconfigurations, and security risks in cloud infrastructure, containers,
and CI/CD pipelines.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(reportCmd)
}