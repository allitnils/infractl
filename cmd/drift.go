package cmd

import (
	"fmt"

	"github.com/allitnils/infractl/internal/azure"
	"github.com/allitnils/infractl/internal/drift"
	"github.com/allitnils/infractl/internal/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var templateFile string
var resourceGroup string

var driftCmd = &cobra.Command{
	Use:   "drift",
	Short: "Detect drift between ARM template and live state",
	RunE: func(cmd *cobra.Command, args []string) error {
		if templateFile == "" {
			return fmt.Errorf("--template is required")
		}
		sub := viper.GetString("subscription")
		if sub == "" {
			sub = "demo-subscription-id"
		}
		client := azure.NewMockClient(sub)
		detector := drift.NewDetector(client)
		changes, err := detector.Compare(templateFile, resourceGroup)
		if err != nil {
			return fmt.Errorf("drift detection failed: %w", err)
		}
		printer := output.NewPrinter(outputFormat)
		return printer.PrintDriftResults(changes)
	},
}

func init() {
	rootCmd.AddCommand(driftCmd)
	driftCmd.Flags().StringVarP(&templateFile, "template", "t", "", "path to ARM template JSON (required)")
	driftCmd.Flags().StringVarP(&resourceGroup, "resource-group", "g", "", "resource group name")
	_ = driftCmd.MarkFlagRequired("template")
}
