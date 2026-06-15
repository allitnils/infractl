package cmd

import (
	"fmt"

	"github.com/allitnils/infractl/internal/azure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Evaluate Azure Policy compliance state",
	RunE: func(cmd *cobra.Command, args []string) error {
		sub := viper.GetString("subscription")
		if sub == "" {
			sub = "demo-subscription-id"
		}
		client := azure.NewMockClient(sub)
		states, err := client.GetPolicyStates()
		if err != nil {
			return fmt.Errorf("policy query failed: %w", err)
		}
		compliant, nonCompliant := 0, 0
		for _, s := range states {
			if s.Compliant {
				compliant++
			} else {
				nonCompliant++
			}
		}
		fmt.Printf("Policy compliance: %d compliant, %d non-compliant\n", compliant, nonCompliant)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(policyCmd)
}
