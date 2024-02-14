package cmd

import "github.com/spf13/cobra"

var acmCmd *cobra.Command

func init() {
	rootCmd.AddCommand(newCmdAcm())
}

func newCmdAcm() *cobra.Command {

	acmCmd = &cobra.Command{
		Use:     "acm",
		Aliases: []string{"certs", "certificates"},
	}

	return acmCmd
}
