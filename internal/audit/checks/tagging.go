package checks

import "github.com/allitnils/infractl/internal/azure"

var requiredTags = []string{"Environment", "Owner", "CostCenter", "Project"}

type TaggingCheck struct{}

func (c *TaggingCheck) Evaluate(r azure.Resource) CheckResult {
	var missing []string
	for _, tag := range requiredTags {
		if _, ok := r.Tags[tag]; !ok {
			missing = append(missing, tag)
		}
	}
	if len(missing) == 0 {
		return CheckResult{Passed: true, Message: "all required tags present"}
	}
	return CheckResult{
		Passed:      false,
		Severity:    "MEDIUM",
		Message:     "missing required tags: " + joinStrings(missing),
		Remediation: "add tags: Environment, Owner, CostCenter, Project",
	}
}
