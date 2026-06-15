package checks

import "github.com/allitnils/infractl/internal/azure"

type IdentityCheck struct{}

func (c *IdentityCheck) Evaluate(r azure.Resource) CheckResult {
	if v, ok := r.Properties["managedIdentityEnabled"].(bool); ok && !v {
		return CheckResult{
			Passed:      false,
			Severity:    "MEDIUM",
			Message:     "managed identity not enabled",
			Remediation: "enable system-assigned managed identity",
		}
	}
	return CheckResult{Passed: true, Message: "identity posture OK"}
}
