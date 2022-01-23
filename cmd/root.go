/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/likexian/whois"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "whereis",
	Short: "A brief description of your application",
	Long:  `A longer description.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		domain := args[0]
		ResolveIP(domain)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.whereis.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("domain", "d", "", "target domain")
}

func ResolveIP(domain string) {
	addr, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		fmt.Println("Resolve error ", err.Error())
		os.Exit(1)
	}
	fmt.Printf("%+v\n", addr.String())
	result, err := whois.Whois(addr.String())
	if err == nil {
		fmt.Println(result)
	}
}
