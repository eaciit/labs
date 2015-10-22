package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

const filename = "/users/ariefdarmawan/Temp/test.xlsx"

func main() {
	_, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Printf("Error loading file: %s \n", err.Error())
		return
	}
	fmt.Printf("Loading file %s successfully \n", filename)
	/*
		for _, sheet := range xlFile.Sheets {
			for _, row := range sheet.Rows {
				for _, cell := range row.Cells {
					fmt.Printf("%s\n", cell.String())
				}
			}
		}
	*/
}
