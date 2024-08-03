from datetime import datetime

import gspread
from oauth2client.service_account import ServiceAccountCredentials


# Define the scope
scope = ["https://spreadsheets.google.com/feeds", "https://www.googleapis.com/auth/drive"]

credentials_path = '/path/to/myproject/credentials.json'

# Add your credentials here
creds = ServiceAccountCredentials.from_json_keyfile_name(credentials_path, scope)
client = gspread.authorize(creds)

# Spreadsheet URL: https://docs.google.com/spreadsheets/d/1jQgH46iRhUhFNxxx/edit
# Configurations
spreadsheet_id = '1jQgH46iRhUhFNxxx'
sheet_title = 'Relisting'
confirm_column = 'Confirmed to process' # Tickbox/Checkbox column
completed_column = 'Process completed'

spreadsheet = client.open_by_key(spreadsheet_id)
print("Spreadsheet ID: ", spreadsheet.id)

# Get sheet by Title
sheet = spreadsheet.worksheet(sheet_title)
print("Sheet Title: ", sheet.title)

# Read from a sheet - Get all records
data = sheet.get_all_records()
print("Sheet Data: ", data)

# Filter data based on "Confirmed to process" column (bool/checkbox)
rows_to_process = [row for row in data if row.get(confirm_column) == 'TRUE']
print("Rows to process length: ", len(rows_to_process))

# Process the data and update "Process completed" column
diff_row = 2 # +2 to account for header row and 0-indexing
for row in rows_to_process:
    row_index = data.index(row) + diff_row
    print(f"Processing row {row_index}...")

    # Perform your processing here
    sheet.update_cell(row_index, sheet.find(completed_column).col, str(datetime.now()))

print("Processing and updating completed.")