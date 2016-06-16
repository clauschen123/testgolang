package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func test_xls() {
	read_xls()
}

func read_xls() {

	excelFileName := `D:\zsDoc\Code\master\ExcelTool\table\DailyActive.xlsx`
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				value, _ := cell.String()
				fmt.Printf("%s\n", value)
			}
		}
	}
}

func write_xls() {

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "I am a cell!"
	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
