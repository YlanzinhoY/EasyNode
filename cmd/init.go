/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ylanzinhoy/services"

	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "init",
	Short: "criar a estrutura inicial para uma aplicacao com node + typescript",
	Long:  "cria uma estrutura inicial para uma aplicacao com node + typescript ",
	Run: func(cmd *cobra.Command, args []string) {
		services.CreateNodeApp()
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
