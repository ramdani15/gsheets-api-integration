package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"ram-go-sheets-api/utils"

	"google.golang.org/api/sheets/v4"
)

func main() {
	ctx := context.Background()
	credentialsFile := "/path/to/myproject/credentials.json"

	// Load the config
	cfg, err := utils.LoadConfig(credentialsFile)
	if err != nil {
		log.Fatalf("Unable to load config: %v", err)
	}

	// Initialize the Sheets service
	sheetsService, err := sheets.NewService(ctx, cfg.ClientOption)
	if err != nil {
		log.Fatalf("Unable to initialize Sheets service: %v", err)
	}

	// Configurations
	spreadsheetID := "1jQgH46iRhUhFNxxx"
	sheetTitle := "Relisting"
	confirmColumn := "Confirmed to process" // Tickbox/Checkbox column
	completedColumn := "Process completed"

	// Get existing spreadsheet
	spreadsheet, err := utils.GetSpreadsheetById(sheetsService, spreadsheetID)
	if err != nil {
		log.Fatalf("Unable to get spreadsheet: %v", err)
	}

	// Get existing sheet
	sheet, err := utils.GetSheetByTitle(spreadsheet, sheetTitle)
	if err != nil {
		log.Fatalf("Unable to get sheet: %v", err)
	}

	// Get data from the sheet
	resp, err := utils.GetSheetValues(sheetsService, spreadsheetID, sheet.Properties.Title)
	if err != nil {
		log.Fatalf("Unable to retrieve data: %v", err)
	}

	headers := resp.Values[0]
	confirmedIndex := -1
	completedIndex := -1

	// Find the index of the relevant columns
	for i, header := range headers {
		if header == confirmColumn {
			confirmedIndex = i
		}
		if header == completedColumn {
			completedIndex = i
		}
	}

	if confirmedIndex == -1 || completedIndex == -1 {
		log.Fatalf("Required columns not found")
	}

	// Filter and update data
	var dataToUpdate []*sheets.ValueRange
	for rowIndex, row := range resp.Values[1:] {
		if len(row) <= confirmedIndex {
			continue
		}

		// set raw value as bool
		if _, ok := row[confirmedIndex].(bool); !ok {
			row[confirmedIndex] = row[confirmedIndex] == "TRUE"
		}

		// update completed column if confirmed
		if row[confirmedIndex].(bool) {
			if len(row) <= completedIndex {
				row = append(row, make([]interface{}, completedIndex-len(row)+1)...)
			}
			row[completedIndex] = time.Now().Format("2006-01-02 15:04:05")
			updateRange := fmt.Sprintf("%s!A%d", sheetTitle, rowIndex+2)
			fmt.Printf("Updating row %d: %v\n", rowIndex+2, row)
			dataToUpdate = append(dataToUpdate, &sheets.ValueRange{
				Range:  updateRange,
				Values: [][]interface{}{row},
			})
		}
	}

	// Update the data
	if len(dataToUpdate) > 0 {
		_, err = sheetsService.Spreadsheets.Values.BatchUpdate(spreadsheetID, &sheets.BatchUpdateValuesRequest{
			ValueInputOption: "RAW",
			Data:             dataToUpdate,
		}).Do()
		if err != nil {
			log.Fatalf("Unable to update data: %v", err)
		}
	}

	fmt.Println("Processing and updating completed.")
}
