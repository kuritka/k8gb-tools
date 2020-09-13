package cmd

import (
	"github.com/kuritka/k8gb-tools/cmd/status"
	"github.com/kuritka/k8gb-tools/pkg/common/guard"

	"github.com/spf13/cobra"
)

var statusOptions status.Options

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "k8gb status",
	//TODO: long description
	//Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		//var err error
		//
		//guard.FailOnError(err, "error when building command context")
		//status := status.New(statusOptions)
		//runner.New(status).MustRun()
	},
}

func init() {
	//TODO: fix description
	statusCmd.Flags().StringVarP(&statusOptions.Gslb, "gslb", "g", "", "name of gslb operator")
	statusCmd.Flag().StringVarP()
	err := statusCmd.MarkFlagRequired("namespace")
	guard.FailOnError(err, "namespace required")
	rootCmd.AddCommand(statusCmd)
}
