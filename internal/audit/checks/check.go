package checks

import (
	"strings"

	"github.com/allitnils/infractl/internal/azure"
)

type Check interface {
	Evaluate(r azure.Resource) CheckResult
}

type CheckResult struct {
	Passed      bool
	Severity    string
	Message     string
	Remediation string
}

func joinStrings(ss []string) string {
	return strings.Join(ss, ", ")
}
