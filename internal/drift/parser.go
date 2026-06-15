package drift

import (
	"encoding/json"
	"os"
)

type ARMTemplate struct {
	Schema    string        `json:"$schema"`
	Resources []ARMResource `json:"resources"`
}

type ARMResource struct {
	Type       string            `json:"type"`
	Name       string            `json:"name"`
	Location   string            `json:"location"`
	Tags       map[string]string `json:"tags"`
	Properties map[string]any    `json:"properties"`
}

func ParseTemplate(path string) (*ARMTemplate, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var t ARMTemplate
	return &t, json.Unmarshal(data, &t)
}
