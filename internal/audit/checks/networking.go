package checks

import "github.com/allitnils/infractl/internal/azure"

type NetworkingCheck struct{}

func (c *NetworkingCheck) Evaluate(r azure.Resource) CheckResult {
	if r.Type == "Microsoft.Network/networkSecurityGroups" {
		if v, ok := r.Properties["allowAllInbound"].(bool); ok && v {
			return CheckResult{
				Passed:      false,
				Severity:    "CRITICAL",
				Message:     "NSG allows unrestricted inbound traffic",
				Remediation: "remove allow-all rules; use least-privilege allowlist",
			}
		}
	}
	return CheckResult{Passed: true, Message: "networking posture OK"}
}
