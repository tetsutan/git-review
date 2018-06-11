package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tetsutan/git-review/src/review"
)

func init() {
	rootCmd.AddCommand(prevCmd)
	rootCmd.AddCommand(nextCmd)

	rootCmd.AddCommand(prevUCmd)
	rootCmd.AddCommand(nextUCmd)

	rootCmd.AddCommand(prevAliasCmd)
	rootCmd.AddCommand(nextAliasCmd)

	rootCmd.AddCommand(prevUAliasCmd)
	rootCmd.AddCommand(nextUAliasCmd)
}

var prevCmd = &cobra.Command{
	Use:   "prev",
	Short: "prev",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.Prev()
	},
}
var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "next",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.Next()
	},
}

var prevUCmd = &cobra.Command{
	Use:   "prev-u",
	Short: "prev uncommented",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.PrevUncommented()
	},
}
var nextUCmd = &cobra.Command{
	Use:   "next-u",
	Short: "next uncommented",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.NextUncommented()
	},
}



// alias
var prevAliasCmd = &cobra.Command{
	Use:   "p",
	Short: "prev alias",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.Prev()
	},
}
var nextAliasCmd = &cobra.Command{
	Use:   "n",
	Short: "next alias",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.Next()
	},
}
var prevUAliasCmd = &cobra.Command{
	Use:   "pu",
	Short: "prev uncommented alias",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.PrevUncommented()
	},
}
var nextUAliasCmd = &cobra.Command{
	Use:   "nu",
	Short: "next uncommented alias",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := review.Review{}
		r.NextUncommented()
	},
}
