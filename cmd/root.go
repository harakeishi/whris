package cmd

import (
	"fmt"
	"os"

	curver "github.com/harakeishi/curver"
	whris "github.com/harakeishi/whris/whris"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "whris",
	Args:  cobra.MinimumNArgs(1),
	Short: "`whris` is Displays management information for IPs associated with the domain.",
	Long:  `"whris" outputs the target domain and IP from the input domain, as well as the administrator information for that IP (administrator name, network name, range of IPs to be managed, and country name).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		version, err := cmd.Flags().GetBool("version")
		if err != nil {
			return err
		}
		if version {
			curver.EchoVersion()
			return nil
		}

		v, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			return err
		}

		for n, domain := range args {
			if n > 0 {
				fmt.Println()
			}
			if _, err := whris.Resolve(domain, v); err != nil {
				return err
			}
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
	rootCmd.Flags().BoolP("version", "V", false, "echo version")
}
