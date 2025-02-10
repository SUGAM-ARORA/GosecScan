package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yourusername/gosecscan/internal/scan"
)

var (
	dockerImage string
	provider    string
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for security vulnerabilities",
	Long:  `Scan cloud infrastructure, containers, and CI/CD pipelines for security issues.`,
}

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Scan AWS resources",
	Run: func(cmd *cobra.Command, args []string) {
		scan.CheckS3Buckets()
	},
}

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Scan Docker images",
	Run: func(cmd *cobra.Command, args []string) {
		scan.ScanContainers(dockerImage)
	},
}

var ciCmd = &cobra.Command{
	Use:   "ci",
	Short: "Scan CI/CD configurations",
	Run: func(cmd *cobra.Command, args []string) {
		scan.ScanGitHubActions()
	},
}

var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Scan Kubernetes clusters",
	Run: func(cmd *cobra.Command, args []string) {
		scan.ScanKubernetes()
	},
}

func init() {
	scanCmd.AddCommand(awsCmd)
	scanCmd.AddCommand(dockerCmd)
	scanCmd.AddCommand(ciCmd)
	scanCmd.AddCommand(k8sCmd)

	dockerCmd.Flags().StringVarP(&dockerImage, "image", "i", "", "Docker image to scan")
	dockerCmd.MarkFlagRequired("image")
}