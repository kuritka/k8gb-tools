package cmd

import (
	"fmt"
	"github.com/kuritka/k8gb-tools/internal/cmd/internal/config"
	"github.com/spf13/cobra"

	"github.com/kuritka/k8gb-tools/internal/cmd/internal/runner"
	"github.com/kuritka/k8gb-tools/internal/cmd/status"
	"github.com/kuritka/k8gb-tools/pkg/common/guard"
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
		status := status.New(statusOptions)
		cfg,err := config.NewConfigFactory(statusOptions.YamlConfig,statusOptions.Gslb)
		if err != nil {
			guard.FailOnError(err,"")
		}
		fmt.Println(cfg)
		runner.New(status).MustRun()
	},
}

func init() {
	//TODO: fix description
	statusCmd.Flags().StringVarP(&statusOptions.Gslb, "gslb", "g", "", "name of gslb operator. List operators if not specified")
	statusCmd.Flags().StringVarP(&statusOptions.YamlConfig, "config", "c", "", "config yaml containing gslb operator and kube config paths. If yaml file is not passed, the default config is chosen")
	//err := statusCmd.MarkFlagRequired("gslb")
	//guard.FailOnError(err, "namespace required")
	rootCmd.AddCommand(statusCmd)
}
