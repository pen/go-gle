package spreadsheets

type Spreadsheets struct {
	service *Service
	id      string
}

func NewSpreadsheets(s *Service, id string) (*Spreadsheets, error) {
	return &Spreadsheets{
		service: s,
		id:      id,
	}, nil
}

func (s *Spreadsheets) GetSpreadsheetByTitle(title string) (*Spreadsheet, error) {
	return NewSpreadsheet(s, title)
}
