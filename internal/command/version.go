package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version",
	Long:  `Prints the version of this binary`,
	Run: func(cmd *cobra.Command, args []string) {
		version := viper.GetString("version")
		gitCommit := viper.GetString("git-commit")
		buildDate := viper.GetString("build-date")
		fmt.Printf("Version: %s\nBuildDate: %s\nGitCommit: %s\n", version, buildDate, gitCommit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
