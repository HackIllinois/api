'''
This script collects all registrants and exports their associated resume file. All resumes are
collected into a directory as noted by variable `DIR` and registrant data is exported to a csv under
the name contained within `OUT_FILE`.

This script is meant to be run external to the API network. It is also reccommended to run this in
a virtual env as this does require some dependencies.

First time setup:

Venv setup:
```
python3 -m venv .venv
source .venv/bin/activate
```

Dependencies:
```
pip install numpy pandas requests python-magic
```

NOTE: For python-magic, you will need to do a little bit of extra setup:
https://github.com/ahupp/python-magic#installation

To run script:
```
HI_AUTH=your_admin_auth_token_here python resume_compile.py
```
'''
import numpy as numpy
import pandas as pd
import requests
import time
import urllib.request
import os
import magic
import mimetypes
from api import make_request

SLEEP_NUM_SECONDS = 0.5
DIR = './resumes'
OUT_FILE = './resumes.csv'

res, ok = make_request("GET", '/registration/attendee/list/')

if not ok:
    print("Failed to get registration data: /registration/attendee/list/")
    exit(1)

df = pd.DataFrame.from_records(res['registrations'])
df['resume'] = ''
df = df.sort_values(by=['major', 'graduationYear'])
df = df.reset_index()

mime = magic.Magic(mime=True)

os.makedirs(DIR, exist_ok=True)

total_rows = len(df.index)
valid_pdf_counter = 0
for i in df.index:
    print(str(i) + "/" + str(total_rows))
    id = df.iloc[i].id

    # get file link from upload endopint
    res, ok = make_request("GET", '/upload/resume/' + str(id) + '/')

    if not ok:
        print(f"Failed to get resume link: /upload/resume/{str(id)}")
        exit(1)

    link = res['resume']

    # download file
    try:
        filename = f'{DIR}/resume{str(valid_pdf_counter)}'
        _, headers = urllib.request.urlretrieve(link, filename)
        mimes = mime.from_file(filename) # Get mime type
        ext = mimetypes.guess_all_extensions(mimes)[0] # Guess extension
        os.rename(filename, filename+ext) # Rename file
        filename = filename+ext
        valid_pdf_counter += 1
    except Exception as e:
        filename = ''
    
    df.at[i, 'resume'] = filename.split('/')[-1]

    time.sleep(SLEEP_NUM_SECONDS) # rate limit

# output csv
df = df[df['resume'] != '']
df = df.reset_index()
df.to_csv(OUT_FILE)
