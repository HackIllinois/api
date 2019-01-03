import os
import requests

from helper import options_menu

API_URL_BASE = 'https://api.hackillinois.org'
AUTHORIZATION_TOKEN = os.environ.get('HI_AUTH')

def make_request(method, endpoint, data = None):
	headers = {
		'Authorization': AUTHORIZATION_TOKEN,
		'Content-Type': 'application/json'
	}

	response = requests.request(method, API_URL_BASE + endpoint, data = data, headers = headers)

	if response.status_code != 200:
		return {}, False

	return response.json(), True

def get_id_by_email(email):
	data, success = make_request('GET', '/user/filter/?email={}'.format(email))

	if not success:
		return '', success

	user_ids = [user['id'] for user in data['users']]

	if len(user_ids) == 0:
		return '', False

	user_id = user_ids[0]

	if len(user_ids) > 1:
		options = {option : option for option in user_ids}
		user_id = options_menu(options)

	return user_id, True
