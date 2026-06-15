package drift

import (
	"fmt"

	"github.com/allitnils/infractl/internal/azure"
)

type DriftChange struct {
	ResourceName string
	ResourceType string
	Field        string
	Expected     any
	Actual       any
	ChangeType   string
}

type Detector struct {
	client azure.AzureClient
}

func NewDetector(client azure.AzureClient) *Detector {
	return &Detector{client: client}
}

func (d *Detector) Compare(templatePath, resourceGroup string) ([]DriftChange, error) {
	template, err := ParseTemplate(templatePath)
	if err != nil {
		return nil, fmt.Errorf("parsing template: %w", err)
	}

	liveResources, err := d.client.ListResources(resourceGroup)
	if err != nil {
		return nil, fmt.Errorf("listing resources: %w", err)
	}

	liveByName := make(map[string]azure.Resource)
	for _, r := range liveResources {
		liveByName[r.Name] = r
	}

	var changes []DriftChange
	for _, tr := range template.Resources {
		live, exists := liveByName[tr.Name]
		if !exists {
			changes = append(changes, DriftChange{
				ResourceName: tr.Name,
				ResourceType: tr.Type,
				ChangeType:   "missing",
				Field:        "resource",
				Expected:     "exists",
				Actual:       "not found",
			})
			continue
		}
		for k, v := range tr.Tags {
			if liveVal, ok := live.Tags[k]; !ok {
				changes = append(changes, DriftChange{
					ResourceName: tr.Name,
					ResourceType: tr.Type,
					ChangeType:   "modified",
					Field:        "tag:" + k,
					Expected:     v,
					Actual:       "(missing)",
				})
			} else if liveVal != v {
				changes = append(changes, DriftChange{
					ResourceName: tr.Name,
					ResourceType: tr.Type,
					ChangeType:   "modified",
					Field:        "tag:" + k,
					Expected:     v,
					Actual:       liveVal,
				})
			}
		}
	}
	return changes, nil
}
