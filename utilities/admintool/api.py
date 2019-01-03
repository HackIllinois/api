import os
import requests

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
