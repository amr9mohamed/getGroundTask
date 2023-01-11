package main

import (
	"github.com/getground/tech-tasks/backend/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "GetGround Party Service",
		Short: "A BE service that allow to organize a party",
	}

	rootCmd.AddCommand(
		cmd.API(),
	)

	cobra.CheckErr(rootCmd.Execute())
}
