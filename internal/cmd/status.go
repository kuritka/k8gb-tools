package cmd

import (
	"github.com/spf13/cobra"

	"github.com/kuritka/k8gb-tools/internal/cmd/internal/runner"
	"github.com/kuritka/k8gb-tools/internal/cmd/status"
)

var statusOptions status.Options

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "k8gb status",
	//TODO: long description
	//Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		status := status.New(statusOptions.YamlConfig, statusOptions.Gslb)
		runner.New(status).MustRun()
	},
}

func init() {
	//TODO: fix description
	statusCmd.Flags().StringVarP(&statusOptions.Gslb, "gslb", "g", "", "name of gslb operator. List operators if not specified")
	statusCmd.Flags().StringVarP(&statusOptions.YamlConfig, "config", "c", "", "config yaml containing gslb operator and kube config paths. If yaml file is not passed, the default config is chosen")
	rootCmd.AddCommand(statusCmd)
}
