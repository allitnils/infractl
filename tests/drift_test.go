package tests

import (
	"os"
	"testing"

	"github.com/allitnils/infractl/internal/azure"
	"github.com/allitnils/infractl/internal/drift"
)

func TestDriftDetector(t *testing.T) {
	templateJSON := `{"$schema":"https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#","resources":[{"type":"Microsoft.KeyVault/vaults","name":"kv-prod-01","location":"australiasoutheast","tags":{"Environment":"Production"},"properties":{}}]}`
	f, err := os.CreateTemp("", "arm-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	_, _ = f.WriteString(templateJSON)
	f.Close()

	client := azure.NewMockClient("test-sub")
	detector := drift.NewDetector(client)
	changes, err := detector.Compare(f.Name(), "my-rg")
	if err != nil {
		t.Fatalf("drift detection failed: %v", err)
	}
	_ = changes
}
