package cmd

import (
	"fmt"
	"os"

	Whereis "github.com/harakeishi/whereis/whereis"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "whereis",
	Short: "`whereis` is Displays management information for IPs associated with the domain.",
	Long:  `"whereis" outputs the target domain and IP from the input domain, as well as the administrator information for that IP (administrator name, network name, range of IPs to be managed, and country name).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		domain := args[0]
		v, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := Whereis.Resolve(domain, v); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
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
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose output")
}
