import json
import csv
from datetime import datetime
import time

import api

from helper import options_menu

def event_upload():
	print('Enter event description csv to upload')
	event_file = input('>')

	with open(event_file, 'r') as f:
		events = csv.DictReader(f, delimiter = ',')

		for event in events:
			event['startTime'] = int(time.mktime(datetime.strptime(event['startTime'], '%m/%d/%Y %H:%M').timetuple()))
			event['endTime'] = int(time.mktime(datetime.strptime(event['endTime'], '%m/%d/%Y %H:%M').timetuple()))

			event['latitude'] = float(event['latitude'])
			event['longitude'] = float(event['longitude'])

			event_data = json.dumps(event)

			response, success = api.make_request('POST', '/event/', data = event_data)

			if not success:
				print('Failed to upload event {}'.format(event_data))

	print('Finished uploading events')

METHODS = {
	'upload': event_upload
}

def event_entry():
	selected_method = options_menu(METHODS)
	METHODS[selected_method]()
