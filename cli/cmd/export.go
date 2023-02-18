/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export exist item data to an Excel file",
	Long: `This command let you create an Excel file that contains item data in a spreadsheet. 
You can specify item type, field names, tags, props arguments to query and filter data. 
If you don't specify the default output fields are code,item type and its quantity.`,
	RunE: func(cmd *cobra.Command, args []string) {
		fmt.Println("export called")
	},
}

func init() {
	excelCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringSliceVarP(
		&FlagFields, "field", "f", FlagFields,
		"a list of fields that you want to export in excel")
	exportCmd.Flags().StringSliceVarP(
		&FlagProps, "prop", "p", FlagProps,
		"a list of property names that you want to export in excel")
	exportCmd.Flags().StringSliceVarP(
		&FlagTags, "tag", "p", FlagTags,
		"a list of tag names that you want to export in excel")
	exportCmd.Flags().StringVarP(
		&FlagFileName, "output-file", "o", FlagFileName,
		"the filename or path you want save you excel file")

	_ = exportCmd.MarkFlagRequired("output-file")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
