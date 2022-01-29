package cmd

import (
	"os"

	Whereis "github.com/harakeishi/whereis/whereis"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "whereis",
	Short: "A brief description of your application",
	Long:  `A longer description.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		domain := args[0]
		Whereis.Resolve(domain)
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("domain", "d", "", "target domain")
}
