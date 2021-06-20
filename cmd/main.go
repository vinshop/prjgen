package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vinshop/prjgen/cmd/prjgen"
	"github.com/vinshop/prjgen/pkg/logger"
	"os"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "vin",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	rootCmd.AddCommand(prjgen.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Errorw("Error when execute command", "error", err)
		os.Exit(0)
	}
}