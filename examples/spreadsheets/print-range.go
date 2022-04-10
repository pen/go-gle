package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pen/go-gle/client"
	"github.com/pen/go-gle/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	jsonKeyFile := flag.String("keyfile", "json.key", "Google API JSON key file")
	ssID := flag.String("ssid", "", "Google Sheets spreadsheets ID")
	title := flag.String("title", "シート1", "Google Sheets spreadsheet title")
	rangeName := flag.String("range", "A1:B2", "Range name")

	flag.Parse()

	if *jsonKeyFile == "" {
		fmt.Fprintln(os.Stderr, "need set -keyfile")
		os.Exit(1)
	}

	if *ssID == "" {
		fmt.Fprintln(os.Stderr, "need set -ssid")
		os.Exit(1)
	}

	if err := printRange(*jsonKeyFile, *ssID, *title, *rangeName); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func printRange(jsonKeyFile, ssID, title, rangeName string) error {
	values, err := getRangeValues(jsonKeyFile, ssID, title, rangeName)
	if err != nil {
		return err
	}

	//nolint:forbidigo
	for _, row := range values {
		for i, cell := range row {
			if i > 0 {
				fmt.Print("\t")
			}

			fmt.Printf("%v", cell)
		}

		fmt.Println()
	}

	return nil
}

func getRangeValues(jsonKeyFile, ssID, title, rangeName string) ([][]interface{}, error) {
	jsonKey, err := os.ReadFile(jsonKeyFile)
	if err != nil {
		return nil, fmt.Errorf(`on os.ReadFile(): %w`, err)
	}

	client, err := client.New(
		option.FromJSON(jsonKey),
		option.WithScopes(sheets.SpreadsheetsReadonlyScope),
	)
	if err != nil {
		return nil, fmt.Errorf(`on client.New(): %w`, err)
	}

	service, err := client.NewSpreadsheetsService()
	if err != nil {
		return nil, fmt.Errorf(`on NewSpreadsheetsService(): %w`, err)
	}

	spreadsheet, err := service.GetSpreadsheetByID(ssID)
	if err != nil {
		return nil, fmt.Errorf(`on GetSpreadsheetsByID(): %w`, err)
	}

	sheet, err := spreadsheet.GetSheetByTitle(title)
	if err != nil {
		return nil, fmt.Errorf(`on GetSpreadsheetByTitle("%s"): %w`, title, err)
	}

	range1, err := sheet.GetRangeByName(rangeName)
	if err != nil {
		return nil, fmt.Errorf(`on GetRange(): %w`, err)
	}

	values, err := range1.GetValues()
	if err != nil {
		return nil, fmt.Errorf(`on GetValues(): %w`, err)
	}

	return values, nil
}
