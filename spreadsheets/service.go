package spreadsheets

import (
	"context"

	"golang.org/x/oauth2/jwt"
	goption "google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Service struct {
	gService *sheets.SpreadsheetsService
}

func NewService(jwt *jwt.Config) (*Service, error) {
	ctx := context.Background()

	gService, err := sheets.NewService(ctx, goption.WithHTTPClient(jwt.Client(ctx)))
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return &Service{
		gService: gService.Spreadsheets,
	}, nil
}

func (s *Service) GetSpreadsheetByID(id string) (*Spreadsheet, error) {
	return GetSpreadsheetsByID(s.gService, id)
}
