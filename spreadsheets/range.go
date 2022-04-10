package spreadsheets

import (
	"fmt"
	"strings"

	"github.com/pen/go-gle/option"
	"google.golang.org/api/sheets/v4"
)

type Range struct {
	gService      *sheets.SpreadsheetsService
	spreadsheetID string
	absName       string
}

func GetRangeByName(
	gService *sheets.SpreadsheetsService,
	gSpreadsheet *sheets.Spreadsheet,
	gSheet *sheets.Sheet,
	name string,
) (*Range, error) {
	return &Range{
		gService:      gService,
		spreadsheetID: gSpreadsheet.SpreadsheetId,
		absName:       rangeAbsName(gSheet.Properties.Title, name),
	}, nil
}

func rangeAbsName(title, name string) string {
	if strings.ContainsAny(name, "'!") {
		return name
	}

	if title[0] != '\'' {
		title = fmt.Sprintf("'%s'", title)
	}

	if name == "" {
		return title
	}

	return fmt.Sprintf("%s!%s", title, name)
}

func (r *Range) GetValues() ([][]interface{}, error) {
	value, err := r.gService.Values.
		Get(r.spreadsheetID, r.absName).
		Do()
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return value.Values, nil
}

func (r *Range) UpdateValues(values [][]interface{}, options ...option.Option) (*sheets.UpdateValuesResponse, error) {
	opt := option.Apply(options, &option.Default{
		ValueInputOption: "USER_ENTERED",
	})

	res, err := r.gService.Values.
		Update(
			r.spreadsheetID,
			r.absName,
			&sheets.ValueRange{Values: values},
		).
		ValueInputOption(opt.ValueInputOption).
		Do()
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return res, nil
}

func (r *Range) ClearValues() (*sheets.ClearValuesResponse, error) {
	res, err := r.gService.Values.
		Clear(
			r.spreadsheetID,
			r.absName,
			&sheets.ClearValuesRequest{},
		).
		Do()
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return res, nil
}
