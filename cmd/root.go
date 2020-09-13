// Package cmd implements three different commands
// - status
// - test
// - install
package cmd

import (
	"fmt"
	"os"

	"github.com/kuritka/k8gb-tools/pkg/common/guard"

	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
)

var (
	//Verbose output
	Verbose bool
)

var rootCmd = &cobra.Command{
	Short: "k8gb plugins",
	//TODO: Long description
	//Long:  `load balancer demo`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			guard.Message("No parameters included")
			_ = cmd.Help()
			os.Exit(0)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n\n bye! %v\n\n", emoji.CrossedFingers)
		//logger.Info().Msg("done..")
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

//Execute runs concrete command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
