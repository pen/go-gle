package spreadsheets

import (
	"fmt"
	"strings"

	"github.com/pen/go-gle/option"
	"google.golang.org/api/sheets/v4"
)

type Range struct {
	spreadSheet *Spreadsheet
	name        string
}

func NewRange(s *Spreadsheet, name string) (*Range, error) {
	return &Range{
		spreadSheet: s,
		name:        rangeAbsName(s.title, name),
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
	ss := r.spreadSheet.spreadSheets
	gs := ss.service.gservice

	value, err := gs.Spreadsheets.Values.
		Get(ss.id, r.name).
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

	ss := r.spreadSheet.spreadSheets
	gs := ss.service.gservice

	res, err := gs.Spreadsheets.Values.
		Update(
			ss.id,
			r.name,
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
	ss := r.spreadSheet.spreadSheets
	gs := ss.service.gservice

	res, err := gs.Spreadsheets.Values.
		Clear(
			ss.id,
			r.name,
			&sheets.ClearValuesRequest{},
		).
		Do()
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return res, nil
}
