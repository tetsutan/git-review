package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tetsutan/git-review/src/review"
)

func init() {
	rootCmd.AddCommand(resetCmd)
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "reset",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.Reset()
	},
}
