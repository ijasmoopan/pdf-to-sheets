package pdf

import "fmt"

func FilterContent(pages []string) []string {
	for i, page := range pages {
		fmt.Printf("\"%s\"\n => %d", page, i)
	}
	return pages
}
