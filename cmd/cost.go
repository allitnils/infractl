package cmd

import (
	"fmt"

	"github.com/allitnils/infractl/internal/azure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var costDays int

var costCmd = &cobra.Command{
	Use:   "cost",
	Short: "Analyse cost trends",
	RunE: func(cmd *cobra.Command, args []string) error {
		sub := viper.GetString("subscription")
		if sub == "" {
			sub = "demo-subscription-id"
		}
		client := azure.NewMockClient(sub)
		summary, err := client.GetCostSummary(costDays)
		if err != nil {
			return fmt.Errorf("cost query failed: %w", err)
		}
		fmt.Printf("Total cost (last %d days): %.2f %s\n", costDays, summary.Total, summary.Currency)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(costCmd)
	costCmd.Flags().IntVarP(&costDays, "days", "d", 30, "number of days to analyse")
}
