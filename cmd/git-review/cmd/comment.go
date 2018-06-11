package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/tetsutan/git-review/src/review"
	"strings"
)

func init() {
	rootCmd.AddCommand(commentCmd)
	rootCmd.AddCommand(goodCmd)
	rootCmd.AddCommand(badCmd)
	rootCmd.AddCommand(skipCmd)

	rootCmd.AddCommand(goodAliasCmd)
	rootCmd.AddCommand(badAliasCmd)
}

var commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "comment",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Comment string needed")
			return
		}

		message := strings.Join(args, " ")


		r := review.Review{}
		if r.Comment(message) {
			return
		}

		return
	},
}


var goodCmd = &cobra.Command{
	Use:   "good",
	Short: "good",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.Good()
	},
}
var badCmd = &cobra.Command{
	Use:   "bad",
	Short: "bad",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.Bad()
	},
}
var skipCmd = &cobra.Command{
	Use:   "skip",
	Short: "skip",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.Skip()
	},
}


// alias
var goodAliasCmd = &cobra.Command{
	Use:   "g",
	Short: "good alias",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.Good()
	},
}
var badAliasCmd = &cobra.Command{
	Use:   "b",
	Short: "bad alias",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.Bad()
	},
}
