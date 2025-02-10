package report

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ScanReport struct {
	Type      string    `json:"type"`
	Severity  string    `json:"severity"`
	Details   string    `json:"details"`
	Timestamp time.Time `json:"timestamp"`
}

type Report struct {
	Summary     ReportSummary  `json:"summary"`
	Findings    []ScanReport   `json:"findings"`
	Statistics  ReportStats    `json:"statistics"`
	GeneratedAt time.Time      `json:"generated_at"`
}

type ReportSummary struct {
	TotalIssues int `json:"total_issues"`
	Critical    int `json:"critical"`
	High        int `json:"high"`
	Medium      int `json:"medium"`
	Low         int `json:"low"`
}

type ReportStats struct {
	ScanDuration string `json:"scan_duration"`
	ResourcesScanned int `json:"resources_scanned"`
}

func SaveReport(findings []ScanReport, filename string) {
	report := prepareReport(findings)
	
	// Create report file
	file, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		fmt.Printf("❌ Error creating report: %v\n", err)
		return
	}

	// Write to file
	err = os.WriteFile(filename+".json", file, 0644)
	if err != nil {
		fmt.Printf("❌ Error saving report: %v\n", err)
		return
	}

	fmt.Printf("✅ Report saved as %s.json\n", filename)
}

func prepareReport(findings []ScanReport) Report {
	now := time.Now()
	summary := calculateSummary(findings)
	
	return Report{
		Summary:     summary,
		Findings:    findings,
		Statistics: ReportStats{
			ScanDuration: "5m30s", // Example duration
			ResourcesScanned: 150,  // Example count
		},
		GeneratedAt: now,
	}
}

func calculateSummary(findings []ScanReport) ReportSummary {
	var summary ReportSummary
	summary.TotalIssues = len(findings)
	
	for _, finding := range findings {
		switch finding.Severity {
		case "CRITICAL":
			summary.Critical++
		case "HIGH":
			summary.High++
		case "MEDIUM":
			summary.Medium++
		case "LOW":
			summary.Low++
		}
	}
	
	return summary
}