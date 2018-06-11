package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tetsutan/git-review/src/review"
)

func init() {
	rootCmd.AddCommand(commentsCmd)
	rootCmd.AddCommand(commentsAliasCmd)
}

var commentsCmd = &cobra.Command{
	Use:   "comments",
	Short: "comments",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.Comments()
	},
}

var commentsAliasCmd = &cobra.Command{
	Use:   "l",
	Short: "comments alias",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.Comments()
	},
}
