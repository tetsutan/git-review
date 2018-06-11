package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tetsutan/git-review/src/review"
)

func init() {
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(statusAliasCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "status",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.Status()
	},
}
var statusAliasCmd = &cobra.Command{
	Use:   "s",
	Short: "status alias",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.Status()
	},
}
