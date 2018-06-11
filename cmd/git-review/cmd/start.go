package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tetsutan/git-review/src/review"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("err", "argument failure")
			return
		}

		spec := args[0]

		p := review.Review{}
		if p.Start(spec) {
			return
		}

		return
	},
}

