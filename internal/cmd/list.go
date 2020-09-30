package cmd

import (
	"github.com/kuritka/k8gb-tools/internal/cmd/internal/runner"
	"github.com/spf13/cobra"

	"github.com/kuritka/k8gb-tools/internal/cmd/list"
)

var listOptions list.Options

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "k8gb list",
	//TODO: long description
	//Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		list := list.New(listOptions.YamlConfig, listOptions.Gslb)
		runner.New(list).MustRun()
	},
}

func init() {
	//TODO: fix description
	listCmd.Flags().StringVarP(&listOptions.YamlConfig, "config", "c", "", "config yaml containing gslb operator and kube config paths. If yaml file is not passed, the default config is chosen")
	rootCmd.AddCommand(listCmd)
}
