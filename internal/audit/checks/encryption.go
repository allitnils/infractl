package checks

import "github.com/allitnils/infractl/internal/azure"

type EncryptionCheck struct{}

func (c *EncryptionCheck) Evaluate(r azure.Resource) CheckResult {
	switch r.Type {
	case "Microsoft.Compute/virtualMachines":
		if enc, ok := r.Properties["diskEncryption"].(string); ok && enc == "disabled" {
			return CheckResult{
				Passed:      false,
				Severity:    "HIGH",
				Message:     "disk encryption is disabled",
				Remediation: "enable Azure Disk Encryption",
			}
		}
	case "Microsoft.Storage/storageAccounts":
		if v, ok := r.Properties["supportsHttpsTrafficOnly"].(bool); ok && !v {
			return CheckResult{
				Passed:      false,
				Severity:    "HIGH",
				Message:     "HTTPS-only traffic not enforced",
				Remediation: "set supportsHttpsTrafficOnly=true",
			}
		}
	}
	return CheckResult{Passed: true, Message: "encryption posture OK"}
}
