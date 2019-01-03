import json

import api

from helper import options_menu

def auth_queryrole():
	print('Enter the email for the user')
	email = input('>')

	user_id, success = api.get_id_by_email(email)

	if not success:
		print('Failed to find user by email')
		return

	data, success = api.make_request('GET', '/auth/roles/{}/'.format(user_id))

	if not success:
		print('Failed to query user roles')
		return

	print('Roles: {}'.format(','.join(data['roles'])))

def auth_addrole():
	print('Enter the email for the user')
	email = input('>')

	user_id, success = api.get_id_by_email(email)

	if not success:
		print('Failed to find user by email')
		return

	print('Enter the role to add')
	role = input('>')

	role_modification = json.dumps({
		'id': user_id,
		'role': role
	})

	data, success = api.make_request('PUT', '/auth/roles/add/', data = role_modification)

	if not success:
		print('Failed to update user roles')
		return

	print('Updated roles: {}'.format(','.join(data['roles'])))

def auth_removerole():
	print('Enter the email for the user')
	email = input('>')

	user_id, success = api.get_id_by_email(email)

	if not success:
		print('Failed to find user by email')
		return

	print('Enter the role to remove')
	role = input('>')

	role_modification = json.dumps({
		'id': user_id,
		'role': role
	})

	data, success = api.make_request('PUT', '/auth/roles/remove/', data = role_modification)

	if not success:
		print('Failed to update user roles')
		return

	print('Updated roles: {}'.format(','.join(data['roles'])))

METHODS = {
	'queryrole': auth_queryrole,
	'addrole': auth_addrole,
	'removerole': auth_removerole
}

def auth_entry():
	selected_method = options_menu(METHODS)
	METHODS[selected_method]()
