package tables

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/remotejob/go_cv_pdf/domains"
)

// func LoadData (fileStr string) {
//     fl, err := os.Open(fileStr)
//     if err == nil {
//         scanner := bufio.NewScanner(fl)
//         var c countryType
//         for scanner.Scan() {
//             // Austria;Vienna;83859;8075
//             lineStr := scanner.Text()
//             list := strings.Split(lineStr, ";")
//             if len(list) == 4 {
//                 c.nameStr = list[0]
//                 c.capitalStr = list[1]
//                 c.areaStr = list[2]
//                 c.popStr = list[3]
//                 countryList = append(countryList, c)
//             } else {
//                 err = fmt.Errorf("error tokenizing %s", lineStr)
//             }
//         }
//         fl.Close()
//         if len(countryList) == 0 {
//             err = fmt.Errorf("error loading data from %s", fileStr)
//         }
//     }
//     if err != nil {
//         pdf.SetError(err)
//     }
// }

func BasicTable(pdf *gofpdf.Fpdf, header []string, countryList []domains.CountryType) {
	for _, str := range header {
		pdf.CellFormat(40, 7, str, "1", 0, "", false, 0, "")
	}
	pdf.Ln(-1)
	for _, c := range countryList {
		pdf.CellFormat(40, 6, c.NameStr, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 6, c.CapitalStr, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 6, c.AreaStr, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 6, c.PopStr, "1", 0, "", false, 0, "")
		pdf.Ln(-1)
	}
}

func ImprovedTable(pdf *gofpdf.Fpdf, header []string, countryList []domains.CountryType) {
	// Column widths
	w := []float64{40.0, 35.0, 40.0, 45.0}
	wSum := 0.0
	for _, v := range w {
		wSum += v
	}
	// 	Header
	for j, str := range header {
		pdf.CellFormat(w[j], 7, str, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)
	// Data
	for _, c := range countryList {
		pdf.CellFormat(w[0], 6, c.NameStr, "LR", 0, "", false, 0, "")
		pdf.CellFormat(w[1], 6, c.CapitalStr, "LR", 0, "", false, 0, "")
		pdf.CellFormat(w[2], 6, c.AreaStr,
			"LR", 0, "R", false, 0, "")
		pdf.CellFormat(w[3], 6, c.PopStr,
			"LR", 0, "R", false, 0, "")
		pdf.Ln(-1)
	}
	pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
}

// Colored table
func FancyTable(pdf *gofpdf.Fpdf, header []string, countryList []domains.CountryType) {
	// Colors, line width and bold font
	pdf.SetFillColor(255, 0, 0)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetDrawColor(128, 0, 0)
	pdf.SetLineWidth(.3)
	pdf.SetFont("", "B", 0)
	// 	Header
	w := []float64{40, 35, 40, 45}
	wSum := 0.0
	for _, v := range w {
		wSum += v
	}
	for j, str := range header {
		pdf.CellFormat(w[j], 7, str, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)
	// Color and font restoration
	pdf.SetFillColor(224, 235, 255)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("", "", 0)
	// 	Data
	fill := false
	for _, c := range countryList {
		pdf.CellFormat(w[0], 6, c.NameStr, "LR", 0, "", fill, 0, "")
		pdf.CellFormat(w[1], 6, c.CapitalStr, "LR", 0, "", fill, 0, "")
		pdf.CellFormat(w[2], 6, c.AreaStr,
			"LR", 0, "R", fill, 0, "")
		pdf.CellFormat(w[3], 6, c.PopStr,
			"LR", 0, "R", fill, 0, "")
		pdf.Ln(-1)
		fill = !fill
	}
	pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
}
