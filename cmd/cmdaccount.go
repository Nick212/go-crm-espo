package cmd

import (
	"github.com/Nick212/go-crm-espo/app"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use: "cmdaccount",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := app.New()

		if err != nil {
			return err
		}
		defer a.Close()

		ctx := a.NewContext(a.Config.Version, "en-us", "", "")

		ctx.Logger.Println("Started GET salesPendingProvisioning")

		ctx.Logger.Println("Finished")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
}
