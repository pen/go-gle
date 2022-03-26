package spreadsheets

type Spreadsheet struct {
	spreadSheets *Spreadsheets
	title        string
}

func NewSpreadsheet(s *Spreadsheets, title string) (*Spreadsheet, error) {
	return &Spreadsheet{
		spreadSheets: s,
		title:        title,
	}, nil
}

func (s *Spreadsheet) GetRange(name string) (*Range, error) {
	return NewRange(s, name)
}
