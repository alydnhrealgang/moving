/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/alydnhrealgang/moving/common/utils"
	"github.com/spf13/cobra"
)

// excelCmd represents the Excel command
var excelCmd = &cobra.Command{
	Use:   "excel",
	Short: "Excel command for moving",
	Long:  `Use for maintenance the data of moving items`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(excelCmd)
	excelCmd.PersistentFlags().StringVarP(
		&FlagItemType, "item-type", utils.EmptyString, FlagItemType, "type for querying items")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// excelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// excelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
