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
