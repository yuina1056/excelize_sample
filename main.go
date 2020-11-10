package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	createfile()
	openfile()
}

func createfile() {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "B2", 100)
	err := f.SaveAs("./test.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func openfile() {
	f, err := excelize.OpenFile("./test.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
}
