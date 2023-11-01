/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ylanzinhoy/services"

	"github.com/spf13/cobra"
)

// expressCmd represents the express command
var expressCmd = &cobra.Command{
	Use:   "express",
	Short: "instala o express",
	Long: `Instala o express totalmente configurado com typescript ja com o arquivo server.ts`,
	Run: func(cmd *cobra.Command, args []string) {
		services.CreateExpressApp()
	},
}

func init() {
	rootCmd.AddCommand(expressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// expressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// expressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
