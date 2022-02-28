import numpy as np
import pandas as pd
import requests
import time
import urllib.request
import os

REGISTRATION_URL = "http://registration.api:8004/"
UPLOAD_URL = "http://upload.api:8008/"

SLEEP_NUM_SECONDS = 2
DIR = './resumes'
OUT_FILE = './resumes.csv'

df = pd.DataFrame.from_records(requests.get(REGISTRATION_URL + 'registration/attendee/list/').json()['registrations'])
df['resume'] = ''
df = df.sort_values(by=['major', 'graduationYear'])
df = df.reset_index()

total_rows = len(df.index)
valid_pdf_counter = 0
for i in df.index:
    print(str(i) + "/" + str(total_rows))
    id = df.iloc[i].id

    # get file link from upload endopint
    link = requests.get(UPLOAD_URL + '/upload/resume/' + str(id) + '/').json()['resume']

    # download file
    try:
        filename = f'{DIR}/resume{str(valid_pdf_counter)}.pdf'
        urllib.request.urlretrieve(link, filename)
        valid_pdf_counter += 1
    except Exception as e:
        filename = ''
    
    df.at[i, 'resume'] = filename.split('/')[-1]

    time.sleep(SLEEP_NUM_SECONDS) # rate limit

# output csv
df = df[df['resume'] != '']
df = df.reset_index()
df.to_csv(OUT_FILE)
