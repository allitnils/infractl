package tests

import (
	"testing"

	"github.com/allitnils/infractl/internal/audit"
	"github.com/allitnils/infractl/internal/azure"
)

func TestAuditRunner(t *testing.T) {
	client := azure.NewMockClient("test-sub")
	runner := audit.NewRunner(client)
	findings, err := runner.Run([]string{"all"})
	if err != nil {
		t.Fatalf("runner failed: %v", err)
	}
	if len(findings) == 0 {
		t.Error("expected findings, got none")
	}
}

func TestTaggingCheck(t *testing.T) {
	client := azure.NewMockClient("test-sub")
	runner := audit.NewRunner(client)
	findings, err := runner.Run([]string{"tagging"})
	if err != nil {
		t.Fatalf("tagging check failed: %v", err)
	}
	hasFailure := false
	for _, f := range findings {
		if !f.Passed {
			hasFailure = true
			break
		}
	}
	if !hasFailure {
		t.Error("expected at least one failing tagging check")
	}
}
