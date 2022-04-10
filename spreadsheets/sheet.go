package spreadsheets

import (
	"errors"

	"google.golang.org/api/sheets/v4"
)

type Sheet struct {
	gService     *sheets.SpreadsheetsService
	gSpreadsheet *sheets.Spreadsheet
	gSheet       *sheets.Sheet
}

func GetSheetByTitle(
	gService *sheets.SpreadsheetsService,
	gSpreadsheet *sheets.Spreadsheet,
	title string,
) (*Sheet, error) {
	for _, gSheet := range gSpreadsheet.Sheets {
		if gSheet.Properties.Title == title {
			return &Sheet{
				gService:     gService,
				gSpreadsheet: gSpreadsheet,
				gSheet:       gSheet,
			}, nil
		}
	}

	return nil, errors.New("sheet not found")
}

func (s *Sheet) GetRangeByName(name string) (*Range, error) {
	return GetRangeByName(s.gService, s.gSpreadsheet, s.gSheet, name)
}
