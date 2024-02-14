package cmd

import "github.com/spf13/cobra"

var secretsManagerCmd *cobra.Command

func init() {
	rootCmd.AddCommand(newCmdSecretsmanager())
}

func newCmdSecretsmanager() *cobra.Command {

	secretsManagerCmd = &cobra.Command{
		Use:   "secretsmanager",
		Short: "sm",
	}

	return secretsManagerCmd
}
