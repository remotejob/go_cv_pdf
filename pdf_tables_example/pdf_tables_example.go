package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/remotejob/go_cv_pdf/domains"
	"github.com/remotejob/go_cv_pdf/pdf_tables_example/tables"
)

func main() {

	pdf := gofpdf.New("P", "mm", "A4", "")
	// type countryType struct {
	//     nameStr, capitalStr, areaStr, popStr string
	// }
	countryList := make([]domains.CountryType, 0, 8)
	header := []string{"Country", "Capital", "Area (sq km)", "Pop. (thousands)"}

	fl, err := os.Open("countries.txt")
	if err == nil {
		scanner := bufio.NewScanner(fl)
		var c domains.CountryType
		for scanner.Scan() {
			// Austria;Vienna;83859;8075
			lineStr := scanner.Text()
			list := strings.Split(lineStr, ";")
			if len(list) == 4 {
				c.NameStr = list[0]
				c.CapitalStr = list[1]
				c.AreaStr = list[2]
				c.PopStr = list[3]
				countryList = append(countryList, c)
			} else {
				err = fmt.Errorf("error tokenizing %s", lineStr)
			}
		}
		fl.Close()
		if len(countryList) == 0 {
			err = fmt.Errorf("error loading data from %s", "coutry")
		}
	}
	if err != nil {
		pdf.SetError(err)
	}

	pdf.SetFont("Arial", "", 14)
	pdf.AddPage()

	tables.BasicTable(pdf, header, countryList)
	pdf.AddPage()

	tables.ImprovedTable(pdf, header, countryList)

	pdf.AddPage()

	tables.FancyTable(pdf, header, countryList)

	err = pdf.OutputFileAndClose("pdf_result/testpdf.pdf")

	if err == nil {
		fmt.Println("Successfully generated my_cv.pdf")
	} else {
		fmt.Println(err)
	}
}
