package pdf

import (
	"fmt"
	"os"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func ReadPdfFile(file string) ([]string, error) {
	pagesWithContent := []string{}
	f, err := os.Open(file)
	if err != nil {
		return pagesWithContent, err
	}
	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return pagesWithContent, err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return pagesWithContent, err
	}

	fmt.Printf("--------------------\n")
	fmt.Printf("PDF to text extraction: %d pages\n", numPages)
	fmt.Printf("--------------------\n")

	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return pagesWithContent, err
		}

		ex, err := extractor.New(page)
		if err != nil {
			return pagesWithContent, err
		}

		text, err := ex.ExtractText()
		if err != nil {
			return pagesWithContent, err
		}

		fmt.Println("------------------------------")
		// fmt.Printf("Page %d:\n", pageNum)
		// fmt.Printf("\"%s\"\n", text)
		pagesWithContent = append(pagesWithContent, text)
		fmt.Println("------------------------------")
	}

	return pagesWithContent, nil
}
