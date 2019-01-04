import collections

def options_menu(options):
	print('Select a option to use:')
	for option_name in options.keys():
		print(option_name)

	while True:
		print('-' * 20)
		selected_option = input('>')

		if selected_option in options:
			return selected_option
			return

		print('Invalid option. Enter a valid option')

def dict_flatten(data, parent_key = '', seperator = '_'):
	items = []
	for k, v in data.items():
		new_key = parent_key + seperator + k if parent_key else k
		if isinstance(v, collections.MutableMapping):
			items.extend(dict_flatten(v, new_key, seperator = seperator).items())
		elif isinstance(v, list):
			items.append((new_key, ','.join(v)))
		else:
			items.append((new_key, v))
	return dict(items)
