import time
import random

import gspread
from oauth2client.service_account import ServiceAccountCredentials


# Define the scope
scope = ["https://spreadsheets.google.com/feeds", "https://www.googleapis.com/auth/drive"]

credentials_path = '/path/to/myproject/credentials.json'

# Add your credentials here
creds = ServiceAccountCredentials.from_json_keyfile_name(credentials_path, scope)
client = gspread.authorize(creds)

unix_id = f"{int(time.time())} Python"
print("Unix ID: ", unix_id)

# Exponential backoff
def share_with_exponential_backoff(spreadsheet, email, perm_type, role):
    max_retries = 5
    for n in range(max_retries):
        try:
            spreadsheet.share(email, perm_type=perm_type, role=role)
            print("Spreadsheet shared successfully.")
            break
        except gspread.exceptions.APIError as e:
            if 'rate limit exceeded' in str(e).lower():
                sleep_time = (2 ** n) + random.uniform(0, 1)
                print(f"Rate limit exceeded, retrying in {sleep_time:.2f} seconds...")
                time.sleep(sleep_time)
            else:
                raise

# Create a new spreadsheet
spreadsheet_name = f"New Spreadsheet {unix_id}"
spreadsheet = client.create(spreadsheet_name)
print("Spreadsheet ID:", spreadsheet.id)
print("Spreadsheet Name:", spreadsheet.title)


# Get the URL of the created spreadsheet
spreadsheet_url = f"https://docs.google.com/spreadsheets/d/{spreadsheet.id}/edit"
print("Spreadsheet URL:", spreadsheet_url)

# Share the spreadsheet
share_with_exponential_backoff(spreadsheet, 'ramdani.r@laku6.com', 'user', 'writer')  # Share with your email

# Add a new sheet
sheet_name = f"New Sheet {unix_id}"
sheet = spreadsheet.add_worksheet(title=sheet_name, rows="100", cols="20")
print("Sheet ID:", sheet.id)
print("Sheet Name:", sheet.title)

