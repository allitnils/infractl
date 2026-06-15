package checks

import "github.com/allitnils/infractl/internal/azure"

type BackupCheck struct{}

func (c *BackupCheck) Evaluate(r azure.Resource) CheckResult {
	if r.Type == "Microsoft.Compute/virtualMachines" {
		if v, ok := r.Properties["backupEnabled"].(bool); ok && !v {
			return CheckResult{
				Passed:      false,
				Severity:    "HIGH",
				Message:     "no backup policy assigned",
				Remediation: "assign a Recovery Services vault backup policy",
			}
		}
	}
	return CheckResult{Passed: true, Message: "backup posture OK"}
}
