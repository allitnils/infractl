package azure

import "fmt"

type AzureClient interface {
	ListResources(resourceGroup string) ([]Resource, error)
	GetCostSummary(days int) (*CostSummary, error)
	GetPolicyStates() ([]PolicyState, error)
}

type MockClient struct {
	SubscriptionID string
}

func NewMockClient(subscriptionID string) AzureClient {
	return &MockClient{SubscriptionID: subscriptionID}
}

func (m *MockClient) ListResources(resourceGroup string) ([]Resource, error) {
	rg := resourceGroup
	if rg == "" {
		rg = "demo-rg"
	}
	return []Resource{
		{
			ID:            fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachines/vm-prod-01", m.SubscriptionID, rg),
			Name:          "vm-prod-01",
			Type:          "Microsoft.Compute/virtualMachines",
			Location:      "australiasoutheast",
			ResourceGroup: rg,
			Tags:          map[string]string{"Environment": "Production"},
			Properties:    map[string]any{"diskEncryption": "disabled"},
		},
		{
			ID:            fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/stprod01", m.SubscriptionID, rg),
			Name:          "stprod01",
			Type:          "Microsoft.Storage/storageAccounts",
			Location:      "australiasoutheast",
			ResourceGroup: rg,
			Tags:          map[string]string{},
			Properties:    map[string]any{"supportsHttpsTrafficOnly": false},
		},
		{
			ID:            fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.KeyVault/vaults/kv-prod-01", m.SubscriptionID, rg),
			Name:          "kv-prod-01",
			Type:          "Microsoft.KeyVault/vaults",
			Location:      "australiasoutheast",
			ResourceGroup: rg,
			Tags:          map[string]string{"Environment": "Production", "Owner": "platform-team", "CostCenter": "IT-001", "Project": "DataPlatform"},
			Properties:    map[string]any{"softDeleteEnabled": true},
		},
	}, nil
}

func (m *MockClient) GetCostSummary(days int) (*CostSummary, error) {
	_ = days
	return &CostSummary{Total: 4250.75, Currency: "USD"}, nil
}

func (m *MockClient) GetPolicyStates() ([]PolicyState, error) {
	return []PolicyState{
		{ResourceID: "vm-prod-01", PolicyDefID: "require-tags", Compliant: false, ComplianceState: "NonCompliant"},
		{ResourceID: "kv-prod-01", PolicyDefID: "require-tags", Compliant: true, ComplianceState: "Compliant"},
	}, nil
}
