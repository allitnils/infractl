package output

import (
	"encoding/json"
	"fmt"

	"github.com/allitnils/infractl/internal/audit"
	"github.com/allitnils/infractl/internal/drift"
	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7C3AED"))
	passStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#22C55E"))
	failStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#EF4444"))
)

type Printer interface {
	PrintAuditResults(findings []audit.Finding) error
	PrintDriftResults(changes []drift.DriftChange) error
}

type TablePrinter struct{}

func (p *TablePrinter) PrintAuditResults(findings []audit.Finding) error {
	fmt.Println(headerStyle.Render("infractl audit results"))
	passed, failed := 0, 0
	for _, f := range findings {
		if f.Passed {
			passed++
			continue
		}
		failed++
		fmt.Printf("%s %s (%s): %s\n", failStyle.Render("[FAIL]"), f.ResourceName, f.CheckName, f.Message)
		if f.Remediation != "" {
			fmt.Printf("       → %s\n", f.Remediation)
		}
	}
	fmt.Printf("\n%s %d passed  %s %d failed\n", passStyle.Render("✓"), passed, failStyle.Render("✗"), failed)
	return nil
}

func (p *TablePrinter) PrintDriftResults(changes []drift.DriftChange) error {
	fmt.Println(headerStyle.Render("infractl drift results"))
	if len(changes) == 0 {
		fmt.Println(passStyle.Render("✓ no drift detected"))
		return nil
	}
	for _, c := range changes {
		fmt.Printf("[%s] %s.%s: expected=%v actual=%v\n", c.ChangeType, c.ResourceName, c.Field, c.Expected, c.Actual)
	}
	return nil
}

type JSONPrinter struct{}

func (p *JSONPrinter) PrintAuditResults(findings []audit.Finding) error {
	b, err := json.MarshalIndent(findings, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func (p *JSONPrinter) PrintDriftResults(changes []drift.DriftChange) error {
	b, err := json.MarshalIndent(changes, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func NewPrinter(format string) Printer {
	if format == "json" {
		return &JSONPrinter{}
	}
	return &TablePrinter{}
}
