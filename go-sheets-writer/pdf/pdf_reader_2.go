package pdf

import (
	"fmt"

	"github.com/ledongthuc/pdf"
)

type Data struct {
	CompanyName string
	Email       string
	Phone       string
}

func ParsePdf(path string) (string, error) {
	file, reader, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		file.Close()
	}()

	totalPages := reader.NumPage()
	for i := 1; i <= totalPages; i++ {
		page := reader.Page(i)
		if page.V.IsNull() {
			continue
		}

		// texts := page.Content().Text
		// var lastTextStyle pdf.Text
		// for _, text := range texts {
		// if isSameSentence(text, lastTextStyle) {
		// 	lastTextStyle.S = lastTextStyle.S + text.S
		// } else {
		// 	fmt.Printf("Font: %s, Font-size: %f, x: %f, y: %f, content: %s \n", lastTextStyle.Font, lastTextStyle.FontSize, lastTextStyle.X, lastTextStyle.Y, lastTextStyle.S)
		// 	lastTextStyle = text
		// }
		// }

		rows, err := page.GetTextByRow()
		if err != nil {
			return "", err
		}

		countId, stat, keyDesc, addr, city, state, zip := "", "", "", "", "", "", ""
		companyName, email, phone := "", "", ""
		for _, row := range rows {
			fmt.Println("**** Row => ", row.Position)
			current, count, extra := "", 0, ""
			for _, word := range row.Content {
				fmt.Println(i, " : ", word.S)
				if word.S != "" {
					current += word.S
				} else {
					if count == 0 {
						countId = current
					} else if count == 1 {
						stat = current
					} else if count == 2 {
						keyDesc = current
					} else if count == 3 {
						companyName = current
					} else if count == 4 {
						addr = current
					} else if count == 5 {
						city = current
					} else if count == 6 {
						state = current
					} else if count == 7 {
						zip = current
					} else if count == 8 {
						email = current
					} else if count == 9 {
						phone = current
					} else {
						extra += current
					}
					current = ""
					count++
				}
			}
			fmt.Printf("countId: %s stat: %s keyDesc: %s\n", countId, stat, keyDesc)
			fmt.Printf("addr: %s city: %s state: %s zip: %s\n", addr, city, state, zip)
			fmt.Printf("companyName: %s email: %s phone: %s\n", companyName, email, phone)
			fmt.Printf("^^^EXTRA => %s\n", extra)
		}
	}

	return "", nil
}

func isSameSentence(textStyle, lastTextStyle pdf.Text) {

}
