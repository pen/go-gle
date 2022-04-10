package spreadsheets

import (
	"google.golang.org/api/sheets/v4"
)

type Spreadsheet struct {
	gService     *sheets.SpreadsheetsService
	gSpreadsheet *sheets.Spreadsheet
}

func GetSpreadsheetsByID(gService *sheets.SpreadsheetsService, id string) (*Spreadsheet, error) {
	gSpreadsheet, err := gService.Get(id).Do()
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return &Spreadsheet{
		gService:     gService,
		gSpreadsheet: gSpreadsheet,
	}, nil
}

func (s *Spreadsheet) GetSheetByTitle(title string) (*Sheet, error) {
	return GetSheetByTitle(s.gService, s.gSpreadsheet, title)
}
