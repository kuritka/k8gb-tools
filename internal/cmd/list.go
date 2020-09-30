package cmd

import (
	"github.com/kuritka/k8gb-tools/internal/cmd/internal/runner"
	"github.com/kuritka/k8gb-tools/pkg/common"
	"github.com/spf13/cobra"

	"github.com/kuritka/k8gb-tools/internal/cmd/list"
)

var listOptions list.Options

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "k8gb list",
	Long:  `Provides a list of gslb accessible by configs. Use gslb name in case you use default configuration or  yaml in case of multiple configurations.`,

	Run: func(cmd *cobra.Command, args []string) {
		list := list.New(listOptions.YamlConfig, listOptions.Gslb)
		runner.New(list).MustRun()
	},
}

func init() {
	listCmd.Flags().StringVarP(&listOptions.YamlConfig, "config", "c", "",
		"config yaml containing gslb operator and kube config paths. "+
			"If yaml file is not passed, the default config is chosen. "+
			"See "+common.GitHubConfigYaml)
	rootCmd.AddCommand(listCmd)
}
