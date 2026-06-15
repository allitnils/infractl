package audit

import (
	"github.com/allitnils/infractl/internal/audit/checks"
	"github.com/allitnils/infractl/internal/azure"
)

type Severity string

const (
	SeverityLow      Severity = "LOW"
	SeverityMedium   Severity = "MEDIUM"
	SeverityHigh     Severity = "HIGH"
	SeverityCritical Severity = "CRITICAL"
)

type Finding struct {
	ResourceID   string
	ResourceName string
	ResourceType string
	CheckName    string
	Severity     Severity
	Message      string
	Remediation  string
	Passed       bool
}

type Runner struct {
	client azure.AzureClient
}

func NewRunner(client azure.AzureClient) *Runner {
	return &Runner{client: client}
}

func (r *Runner) Run(checkNames []string) ([]Finding, error) {
	resources, err := r.client.ListResources("")
	if err != nil {
		return nil, err
	}

	allChecks := map[string]checks.Check{
		"tagging":    &checks.TaggingCheck{},
		"encryption": &checks.EncryptionCheck{},
		"networking": &checks.NetworkingCheck{},
		"identity":   &checks.IdentityCheck{},
		"backup":     &checks.BackupCheck{},
	}

	runAll := len(checkNames) == 1 && checkNames[0] == "all"

	var findings []Finding
	for name, check := range allChecks {
		shouldRun := runAll
		if !shouldRun {
			for _, requested := range checkNames {
				if requested == name {
					shouldRun = true
					break
				}
			}
		}
		if !shouldRun {
			continue
		}
		for _, res := range resources {
			result := check.Evaluate(res)
			findings = append(findings, Finding{
				ResourceID:   res.ID,
				ResourceName: res.Name,
				ResourceType: res.Type,
				CheckName:    name,
				Severity:     Severity(result.Severity),
				Message:      result.Message,
				Remediation:  result.Remediation,
				Passed:       result.Passed,
			})
		}
	}
	return findings, nil
}
