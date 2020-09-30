// Package cmd implements three different commands
// - status, gets and validates all gslbs cross the contexts
// - install, install gslb to context
// - list, returns list of gslbs in particular contexts and namespaces
package cmd

import (
	"fmt"
	"os"

	"github.com/enescakir/emoji"
	"github.com/kuritka/k8gb-tools/pkg/common"
	"github.com/logrusorgru/aurora"

	"github.com/kuritka/k8gb-tools/pkg/common/guard"

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
		fmt.Printf("\n Not sure what to do %s? check out %s! %s\n", aurora.BrightGreen("next"), aurora.BrightBlue(common.GitHubUrl), emoji.BeachWithUmbrella)
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
