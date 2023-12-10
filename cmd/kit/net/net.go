package net

import "github.com/spf13/cobra"

var _net = &cobra.Command{
	Use:   "net",
	Short: "net command kit",
}

func Setup(kit *cobra.Command) {
	kit.AddCommand(_net)
}
