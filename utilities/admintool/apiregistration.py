import json
import csv

import api

from helper import options_menu, dict_flatten

def registration_download():
	data, success = api.make_request('GET', '/registration/attendee/list/')

	if not success:
		print('Failed to download registrations')
		return

	registrations = [dict_flatten(reg) for reg in data['registrations']]
	fields = registrations[0].keys()

	print('Enter the csv file location to save registration')
	output_file = input('>')

	with open(output_file, 'w') as f:
		csv_writer = csv.DictWriter(f, fieldnames = fields)
		csv_writer.writeheader()

		for registration in registrations:
			csv_writer.writerow(registration)

	print('Saved registrations to file')

METHODS = {
	'download': registration_download
}

def registration_entry():
	selected_method = options_menu(METHODS)
	METHODS[selected_method]()
