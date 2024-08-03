# Google Sheets API Integration with Go and Python

This project demonstrates how to interact with Google Sheets API using both Go and Python. It covers creating, reading, updating, and handling data validation for Google Sheets, as well as sharing access to the spreadsheet.

## Features

- Create a new spreadsheet
- Read data from a spreadsheet
- Update data in a spreadsheet
- Handle data validation rules
- Share spreadsheet access (commentator, writer, owner)

## Prerequisites

1. Google Cloud account
2. Google Sheets API and Google Drive API enabled in your Google Cloud project
3. Service account credentials JSON file

## Project Structure

```
myproject/
├── go/
│   └── generate/
│       └── main.go
│   └── relist/
│       └── main.go
│   └── utils/
│       └── utils.go
│   └── go.mod
│   └── go.sum
├── python/
│   ├── generate.py
│   └── relist.py
└── credentials.json
```

## Setting Up the Environment

### Project Setup

1. Copy `credentials.json.template` to `credentials.json` and update with your Service account credentials JSON file
2. Update the path to credentials file on the code with the absolute path of your credentials file
3. Run the code

### Go Environment Setup

1. Install Go
2. Install Dependencies:
    ```
    go get -u google.golang.org/api/sheets/v4
    go get -u golang.org/x/oauth2/google
    go get -u google.golang.org/api/drive/v3
    ```

### Python Environment Setup
1. Install Python
2. Install Dependencies:
    ```
    pip install gspread oauth2client
    ```


## Running the code

### Go

#### /go/generate/main.go
The script will be used to create a new spreadsheet with the worksheet.
```
> cd path/to/myproject/go

> go run generate/main.go
```

#### /go/relist/main.go
The script will be used to get the data from one of worksheet from the specific spreadsheet. The data will be process based on the condition of a specific column (Confirmed to process) then updated another column (Process completed) when the process is done.
```
> cd path/to/myproject/go

> go run relist/main.go
```

### Python

#### /python/generate.py
The script will be used to create a new spreadsheet with the worksheet.
```
> cd path/to/myproject/python

> python generate.py
```

#### /python/relist.py
The script will be used to get the data from one of worksheet from the specific spreadsheet. The data will be process based on the condition of a specific column (Confirmed to process) then updated another column (Process completed) when the process is done.
```
> cd path/to/myproject/python

> python relist.py
```

## Contributing
If you would like to contribute, please fork the repository and use a feature branch. Pull requests are welcome.