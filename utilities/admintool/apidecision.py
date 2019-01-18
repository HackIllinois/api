import json

import api

from helper import options_menu

def decision_query():
	print('Enter the email for the user')
	email = input('>')

	user_id, success = api.get_id_by_email(email)

	if not success:
		print('Failed to find user by email')
		return

	data, success = api.make_request('GET', '/decision/{}/'.format(user_id))

	if not success:
		print('Failed to query user decision')
		return

	print(json.dumps(data, indent=4, sort_keys=True))

def decision_update():
	print('Enter the email for the user')
	email = input('>')

	user_id, success = api.get_id_by_email(email)

	if not success:
		print('Failed to find user by email')
		return

	print('Enter decision status')
	status = input('>')

	wave = 0
	if status == 'ACCEPTED':
		print('Enter acceptance wave')
		wave = input('>')

	decision = json.dumps({
		'id': user_id,
		'status': status,
		'wave': wave
	})

	data, success = api.make_request('POST', '/decision/', data = decision)

	if not success:
		print('Failed to update user decision')
		return

	print('Updated user decision')

def decision_finalize():
	print('Enter the email for the user')
	email = input('>')

	user_id, success = api.get_id_by_email(email)

	if not success:
		print('Failed to find user by email')
		return

	print('Enter finalize or unfinalize')
	finalize = (input('>') == 'finalize')

	finalization = json.dumps({
		'id': user_id,
		'finalized': finalize
	})

	data, success = api.make_request('POST', '/decision/finalize/', data = finalization)

	if not success:
		print('Failed to finalize user decision')
		return

	print('Finalized user decision')

METHODS = {
	'query': decision_query,
	'update': decision_update,
	'finalize': decision_finalize
}

def decision_entry():
	selected_method = options_menu(METHODS)
	METHODS[selected_method]()
