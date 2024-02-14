package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "aws-helper",
	Short: "CMS AWS HELPER",
	Long:  "Do funny stuff in different AWS services",
}

func Execute() int {

	if rootCmd.Execute() != nil {
		return 1
	}

	return 0
}
