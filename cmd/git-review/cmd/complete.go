package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tetsutan/git-review/src/review"
)

func init() {
	rootCmd.AddCommand(completeCmd)
}

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "complete",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.Complete()
	},
}
