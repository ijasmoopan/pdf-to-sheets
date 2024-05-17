package main

import (
	"fmt"
	"os"

	"github.com/ijasmoopan/pdf-reader/sheets"
)

func main() {
	// read from json file
	fmt.Println("Welcome to Sheets Writer by GO..")
	pdfData, err := sheets.ReadFromJsonFile()
	if err != nil {
		fmt.Println("Error in ReadPdfFile", err)
		os.Exit(1)
	}

	// write to google sheets
	err = sheets.WriteToSheet(pdfData)
	if err != nil {
		fmt.Println("Error in ReadPdfFile", err)
		os.Exit(1)
	}
	fmt.Println("Signing off from GO!")
}
