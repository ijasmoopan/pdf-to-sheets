package sheets

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ijasmoopan/pdf-reader/pdf"
)

type JsonObject struct {
	Options struct {
		HasTitles    bool          `json:"hasTitles"`
		Threshold    float64       `json:"threshold"`
		MaxStrLength int           `json:"maxStrLength"`
		IgnoreTexts  []interface{} `json:"ignoreTexts"`
	} `json:"_options"`
	NumPages int `json:"numPages"`
	Pages    []struct {
		PageNumber int     `json:"pageNumber"`
		Tables     []Table `json:"tables"`
	} `json:"pages"`
}

type Table struct {
	TableNumber int        `json:"tableNumber"`
	NumRows     int        `json:"numrows"`
	NumCols     int        `json:"numcols"`
	Data        [][]string `json:"data"`
}

func ReadFromJsonFile() ([]pdf.Data, error) {
	pdfFile := "test-response.json"
	if len(os.Args) > 1 {
		pdfFile = os.Args[1]
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	filePath := wd + "/../" + pdfFile

	fmt.Println("Reading json file: ", filePath)
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var jsonObject JsonObject
	err = json.Unmarshal(bytes, &jsonObject)
	if err != nil {
		return nil, err
	}

	var listData []pdf.Data
	for i, page := range jsonObject.Pages {
		// if you want table data from specific pages, uncomment the below one to add condition
		// add 3 on each sides
		// if i >= 80 && i <= 90 { // ponnoo
		if i >= 70 && i <= 72 { // ponnoo
			// if i >= 42 && i <= 44 { // Vindhya
			for _, table := range page.Tables {
				for k, v := range table.Data {
					if k != 0 {
						if len(v) > 9 {
							listData = append(listData, pdf.Data{
								CompanyName: v[3], Email: v[8], Phone: v[9],
							})
						} else if len(v) > 8 {
							listData = append(listData, pdf.Data{
								CompanyName: v[3], Email: v[8], Phone: "",
							})
						} else if len(v) > 3 {
							listData = append(listData, pdf.Data{
								CompanyName: v[3], Email: "", Phone: "",
							})
						} else {
							listData = append(listData, pdf.Data{
								CompanyName: "", Email: "", Phone: "",
							})
						}
					}
				}
			}
		}
	}

	// for i, d := range listData {
	// 	fmt.Println(i, " => ", d.CompanyName, "  |  ", d.Email, "  |  ", d.Phone)
	// }

	fmt.Println("-----------------------------------------------><----------------------------------------------------")
	for _, d := range listData {
		fmt.Println(d.CompanyName)
	}

	fmt.Println("-----------------------------------------------><----------------------------------------------------")
	for _, d := range listData {
		fmt.Println(d.Email)
	}

	fmt.Println("-----------------------------------------------><----------------------------------------------------")
	for _, d := range listData {
		fmt.Println(d.Phone)
	}

	fmt.Println("-----------------------------------------------><----------------------------------------------------")
	fmt.Println("Length of rows", len(listData))

	return listData, err
}
