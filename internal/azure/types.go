package azure

type Resource struct {
	ID             string
	Name           string
	Type           string
	Location       string
	ResourceGroup  string
	SubscriptionID string
	Tags           map[string]string
	Properties     map[string]any
}

type CostSummary struct {
	Total    float64
	Currency string
}

type PolicyState struct {
	ResourceID      string
	PolicyDefID     string
	Compliant       bool
	ComplianceState string
}
