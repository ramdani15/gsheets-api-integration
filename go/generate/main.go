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

	// Set the unique spreadsheet title
	unixId := fmt.Sprintf("%d Golang", time.Now().Unix())
	fmt.Println("unixId:", unixId)

	// Create a new spreadsheet
	spreadsheetTitle := "New Spreadsheet " + unixId
	spreadsheetID := utils.CreateSpreadsheet(sheetsService, spreadsheetTitle)
	fmt.Println("Spreadsheet URL: https://docs.google.com/spreadsheets/d/" + spreadsheetID + "/edit")

	// Create a new worksheet
	sheetTitle := "New Sheet " + unixId
	utils.CreateWorksheet(sheetsService, spreadsheetID, sheetTitle)

	// Share the spreadsheet
	driveService, err := utils.NewDriveService(cfg.ClientOption)
	if err != nil {
		log.Fatalf("Unable to initialize Drive service: %v", err)
	}

	err = utils.ShareSpreadsheet(driveService, spreadsheetID, "ramdani.r@laku6.com", "writer")
	if err != nil {
		log.Fatalf("Unable to share spreadsheet: %v", err)
	}

	log.Printf("Shared spreadsheet with ID: %s", spreadsheetID)
}
