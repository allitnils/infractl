package cmd

import (
	"fmt"

	"github.com/allitnils/infractl/internal/audit"
	"github.com/allitnils/infractl/internal/azure"
	"github.com/allitnils/infractl/internal/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var auditChecks []string

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Run infrastructure compliance checks",
	RunE: func(cmd *cobra.Command, args []string) error {
		sub := viper.GetString("subscription")
		if sub == "" {
			sub = "demo-subscription-id"
		}
		client := azure.NewMockClient(sub)
		runner := audit.NewRunner(client)
		results, err := runner.Run(auditChecks)
		if err != nil {
			return fmt.Errorf("audit failed: %w", err)
		}
		printer := output.NewPrinter(outputFormat)
		return printer.PrintAuditResults(results)
	},
}

func init() {
	rootCmd.AddCommand(auditCmd)
	auditCmd.Flags().StringSliceVarP(&auditChecks, "checks", "c", []string{"all"}, "checks: tagging,encryption,networking,identity,backup,all")
}
