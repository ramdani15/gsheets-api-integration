package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Config struct {
	ClientOption option.ClientOption
}

func LoadConfig(credentialsFile string) (*Config, error) {
	credentialsData, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		return nil, err
	}

	_, err = google.JWTConfigFromJSON(credentialsData, sheets.SpreadsheetsScope, drive.DriveScope)
	if err != nil {
		return nil, err
	}

	clientOption := option.WithCredentialsJSON(credentialsData)

	return &Config{ClientOption: clientOption}, nil
}

func NewDriveService(clientOption option.ClientOption) (*drive.Service, error) {
	ctx := context.Background()
	return drive.NewService(ctx, clientOption)
}

func CreateSpreadsheet(sheetsService *sheets.Service, title string) string {
	spreadsheet := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: title,
		},
	}

	createSpreadsheetResponse, err := sheetsService.Spreadsheets.Create(spreadsheet).Do()
	if err != nil {
		log.Fatalf("Unable to create spreadsheet: %v", err)
	}

	spreadsheetID := createSpreadsheetResponse.SpreadsheetId

	fmt.Println("Spreadsheet ID:", spreadsheetID)
	fmt.Println("Spreadsheet Title:", createSpreadsheetResponse.Properties.Title)

	return spreadsheetID
}

func CreateWorksheet(sheetsService *sheets.Service, spreadsheetID, title string) string {
	addSheetRequest := &sheets.Request{
		AddSheet: &sheets.AddSheetRequest{
			Properties: &sheets.SheetProperties{
				Title: title,
			},
		},
	}
	batchUpdateSpreadsheetRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{addSheetRequest},
	}
	_, err := sheetsService.Spreadsheets.BatchUpdate(spreadsheetID, batchUpdateSpreadsheetRequest).Do()
	if err != nil {
		log.Fatalf("Unable to add sheet: %v", err)
	}

	fmt.Printf("Sheet Title: %s\n", title)

	return title
}

func ShareSpreadsheet(driveService *drive.Service, spreadsheetID, email, role string) error {
	permission := &drive.Permission{
		Type:         "user",
		Role:         role,
		EmailAddress: email,
	}

	_, err := driveService.Permissions.Create(spreadsheetID, permission).Do()
	if err != nil {
		return fmt.Errorf("unable to create permission: %v", err)
	}

	return nil
}

func GetSpreadsheetById(sheetsService *sheets.Service, spreadsheetID string) (*sheets.Spreadsheet, error) {
	spreadsheet, err := sheetsService.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to get spreadsheet: %v", err)
	}

	return spreadsheet, nil
}

func GetSheetByTitle(spreadsheet *sheets.Spreadsheet, sheetTitle string) (*sheets.Sheet, error) {
	for _, sheet := range spreadsheet.Sheets {
		if sheet.Properties.Title == sheetTitle {
			return sheet, nil
		}
	}

	return nil, fmt.Errorf("sheet with title %s not found", sheetTitle)
}

func GetSheetValues(sheetsService *sheets.Service, spreadsheetID, sheetTitle string) (*sheets.ValueRange, error) {
	resp, err := sheetsService.Spreadsheets.Values.Get(spreadsheetID, sheetTitle).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to get sheet values: %v", err)
	}

	if len(resp.Values) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	return resp, nil
}
