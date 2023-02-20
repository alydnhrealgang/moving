/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/alydnhrealgang/moving/cli/api"
	"github.com/alydnhrealgang/moving/cli/api/models"
	"github.com/alydnhrealgang/moving/cli/api/moving_clients/operations"
	"github.com/alydnhrealgang/moving/common/utils"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "update item data from an excel file",
	Long: `This command offers you to batch update the 
fields, tags or props of items from a Excel file you previous exported.
Please note that the 'code' column must be the first column that addressed in A in Excel sheet and
The name of sheet in Excel must be named as 'MOVING'.
Also you should provide which columns you want to update, there 3 columns arguments:
1. Fields: (name|count|description)=(EXCEL COLUMN) e.g. field=name=B,count=C,description=D
2. Properties: (any property name)=(EXCEL COLUMN) e.g. prop=material=E,brand=F
3. Tags: (any tag name)=(EXCEL COLUMN) e.g. tag=category=G,color=H`,
	RunE: func(cmd *cobra.Command, args []string) error {
		moving, err := api.CreateApiClient(FlagApiUrl)
		if nil != err {
			return err
		}

		excel, err := excelize.OpenFile(FlagFileName)
		if nil != err {
			return err
		}

		itemsToUpdate := make([]*models.ItemData, 0)
		for i := 2; ; i++ {
			address := fmt.Sprintf("A%d", i)
			cmd.Println("Processing:", address)
			code, err := excel.GetCellValue("MOVING", address)
			if nil != err {
				return err
			}
			if utils.EmptyOrWhiteSpace(code) {
				cmd.Println("stop processing")
				break
			}
			params := &operations.GetItemByCodeParams{
				Code: code,
			}
			params.WithTimeout(time.Second * 60)
			resp, err := moving.Operations.GetItemByCode(params)
			if nil != err {
				return err
			}
			if len(resp.GetPayload()) <= 0 {
				return fmt.Errorf("cannot found item by code: %s", code)
			}
			item := resp.GetPayload()[0]
			itemsToUpdate = append(itemsToUpdate, item)
			cmd.Println("starting processing item:", code)
			for name, col := range FlagFieldMap {
				address := fmt.Sprintf("%s%d", strings.ToUpper(col), i)
				cmd.Printf("processing field %s at %s\n", name, address)
				value, err := excel.GetCellValue("MOVING", address)
				if nil != err {
					return err
				}
				switch strings.ToLower(name) {
				case "name":
					cmd.Println("set item name to:", value)
					item.Name = value
				case "description":
					cmd.Println("set item description to:", value)
					item.Description = value
				case "quantity":
					quantity, err := strconv.ParseInt(value, 10, 64)
					if nil != err {
						return err
					}
					cmd.Println("set item quantity to:", value)
					item.Count = quantity
				}
			}
			for name, col := range FlagPropMap {
				address := fmt.Sprintf("%s%d", strings.ToUpper(col), i)
				cmd.Printf("processing property %s at %s\n", name, address)
				value, err := excel.GetCellValue("MOVING", address)
				if nil != err {
					return err
				}
				if nil == item.Props {
					item.Props = make(map[string]string)
				}
				cmd.Println("set property value to: ", value)
				item.Props[name] = value
			}

			for name, col := range FlagTagMap {
				address := fmt.Sprintf("%s%d", strings.ToUpper(col), i)
				cmd.Printf("processing tag %s at %s\n", name, address)
				value, err := excel.GetCellValue("MOVING", address)
				if nil != err {
					return err
				}
				if nil == item.Tags {
					item.Tags = make(map[string]string)
				}
				if utils.EmptyOrWhiteSpace(value) {
					cmd.Println("delete item tag:", name)
					delete(item.Tags, name)
				} else {
					cmd.Println("set item tag value to:", value)
					item.Tags[name] = value
				}
			}
		}

		//for _, item := range itemsToUpdate {
		//	params := operations.NewSaveItemParamsWithTimeout(time.Second * 60).WithBody(item)
		//	cmd.Println("saving item %s", item.Code)
		//	_, err := moving.Operations.SaveItem(params)
		//	if nil != err {
		//		return err
		//	}
		//}

		return nil
	},
}

func init() {
	excelCmd.AddCommand(importCmd)
	importCmd.Flags().StringToStringVarP(
		&FlagFieldMap, "field", "f", FlagFieldMap, "a map of fields that you want to update")
	importCmd.Flags().StringToStringVarP(
		&FlagPropMap, "prop", "p", FlagPropMap, "a map of properties that you want to update")
	importCmd.Flags().StringToStringVarP(
		&FlagTagMap, "tag", "t", FlagTagMap, "a map of tags that you want to update")
	importCmd.Flags().StringVarP(
		&FlagFileName, "file", utils.EmptyString, utils.EmptyString, "source Excel file name or path")

	_ = importCmd.MarkFlagRequired("file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
