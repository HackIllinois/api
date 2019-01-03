import json

import api

from helper import options_menu

def stat_download():
	data, success = api.make_request('GET', '/stat/')

	if not success:
		print('Failed to download stats')
		return

	print(json.dumps(data, indent=4, sort_keys=True))

METHODS = {
	'download': stat_download
}

def stat_entry():
	selected_method = options_menu(METHODS)
	METHODS[selected_method]()
