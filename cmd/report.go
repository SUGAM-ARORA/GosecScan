package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yourusername/gosecscan/internal/report"
)

var (
	format   string
	filename string
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate security reports",
	Run: func(cmd *cobra.Command, args []string) {
		results := []report.ScanReport{
			{
				Type:     "aws",
				Severity: "HIGH",
				Details:  "Public S3 bucket found",
			},
		}
		report.SaveReport(results, filename)
	},
}

func init() {
	reportCmd.Flags().StringVarP(&format, "format", "f", "json", "Report format (json/markdown)")
	reportCmd.Flags().StringVarP(&filename, "output", "o", "security-report", "Output filename")
}