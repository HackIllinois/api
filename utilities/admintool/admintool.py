#!/usr/bin/env python3

import sys
import os

from apistat import stat_entry
from helper import options_menu

MODULES = {
	'registration': None,
	'event': None,
	'stat': stat_entry
}

def main():
	if os.environ.get('HI_AUTH') == None:
		print('Must set HI_AUTH environment variable')
		return

	selected_module = options_menu(MODULES)
	MODULES[selected_module]()

if __name__== '__main__':
	main()
