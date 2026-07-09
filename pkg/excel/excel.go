package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
	"time"
)

// ExportExcel generates an Excel file from data and returns the file path
// headers: column headers
// rows: [][]interface{} data rows
// sheetName: sheet name
func ExportExcel(headers []string, rows [][]interface{}, sheetName string) (string, error) {
	f := excelize.NewFile()
	defer f.Close()

	sheet := "Sheet1"
	if sheetName != "" {
		sheet = sheetName
	}
	f.SetSheetName("Sheet1", sheet)

	// Set headers
	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue(sheet, cell, header)
	}

	// Set data rows
	for i, row := range rows {
		for j, value := range row {
			cell := fmt.Sprintf("%c%d", 'A'+j, i+2)
			f.SetCellValue(sheet, cell, value)
		}
	}

	// Save to temp file
	fileName := fmt.Sprintf("export_%d.xlsx", time.Now().UnixNano())
	filePath := filepath.Join(os.TempDir(), fileName)
	if err := f.SaveAs(filePath); err != nil {
		return "", err
	}
	return filePath, nil
}
