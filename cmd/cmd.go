package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "mim",
	Short: "Mim is a cli for mirrororrim",
	Long: "Mim is a cli for mirrororrim, the site parallel mirroring tool written in golang",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Exec() {

}