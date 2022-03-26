package spreadsheets

import (
	"context"

	"golang.org/x/oauth2/jwt"
	goption "google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Service struct {
	gservice *sheets.Service
}

func NewService(jwt *jwt.Config) (*Service, error) {
	ctx := context.Background()

	googleSheetsService, err := sheets.NewService(ctx, goption.WithHTTPClient(jwt.Client(ctx)))
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return &Service{
		gservice: googleSheetsService,
	}, nil
}

func (s *Service) GetSpreadsheetsByID(id string) (*Spreadsheets, error) {
	return NewSpreadsheets(s, id)
}
