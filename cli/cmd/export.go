/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/alydnhrealgang/moving/cli/api"
	"github.com/alydnhrealgang/moving/cli/api/moving_clients/operations"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"math"
	"strings"
	"time"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export exist item data to an Excel file",
	Long: `This command let you create an Excel file that contains item data in a spreadsheet. 
You can specify item type, field names, tags, props arguments to query and filter data. 
If you don't specify the default output fields are code,item type and its quantity.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		moving, err := api.CreateApiClient(FlagApiUrl)
		if nil != err {
			return err
		}

		params := &operations.QueryItemsParams{
			FetchSize: math.MaxInt64,
			Type:      FlagItemType,
			Keyword:   "*",
		}
		params.WithTimeout(time.Second * 60)
		resp, err := moving.Operations.QueryItems(params)

		if nil != err {
			return err
		}

		excelFile := excelize.NewFile()
		defer func() {
			err = excelFile.Close()
			if nil != err {
				cmd.PrintErrln(err)
			}
		}()

		sheetIndex, err := excelFile.NewSheet("MOVING")
		if nil != err {
			return err
		}
		excelFile.SetActiveSheet(sheetIndex)
		titles := []interface{}{"code", "item type"}
		titles = append(titles, lo.Map(FlagFields, func(s string, _ int) interface{} { return s })...)
		titles = append(titles, lo.Map(FlagProps, func(s string, _ int) interface{} { return s })...)
		titles = append(titles, lo.Map(FlagTags, func(s string, _ int) interface{} { return s })...)
		texts := [][]interface{}{titles}
		for _, item := range resp.GetPayload() {
			row := []interface{}{item.Code, item.Type}
			for _, field := range FlagFields {
				switch strings.ToLower(field) {
				case "name":
					row = append(row, item.Name)
				case "quantity":
					row = append(row, item.Count)
				case "description":
					row = append(row, item.Description)
				case "boxCode":
					row = append(row, item.BoxCode)
				default:
					row = append(row, "n/a")
				}
			}
			if nil == item.Props {
				item.Props = make(map[string]string, 0)
			}
			for _, prop := range FlagProps {
				row = append(row, item.Props[prop])
			}
			if nil == item.Tags {
				item.Tags = make(map[string]string, 0)
			}
			for _, tag := range FlagTags {
				row = append(row, item.Tags[tag])
			}
			texts = append(texts, row)
		}

		for rowIndex, row := range texts {
			for colIndex, value := range row {
				cell := ExcelIndex2Column(colIndex + 1)
				address := fmt.Sprintf("%s%d", cell, rowIndex+1)
				err = excelFile.SetCellValue("MOVING", address, value)
				if nil != err {
					cmd.PrintErrln("write content to cell: ", address, " failed.")
					return err
				}
			}
		}

		return excelFile.SaveAs(FlagFileName)
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
		&FlagTags, "tag", "t", FlagTags,
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
