#!/usr/bin/env python3

import sys
import os

from apistat import stat_entry
from apievent import event_entry
from apiregistration import registration_entry
from apiauth import auth_entry
from apidecision import decision_entry
from helper import options_menu

MODULES = {
	'registration': registration_entry,
	'event': event_entry,
	'stat': stat_entry,
	'auth': auth_entry,
	'decision': decision_entry
}

def main():
	if os.environ.get('HI_AUTH') == None:
		print('Must set HI_AUTH environment variable')
		return

	selected_module = options_menu(MODULES)
	MODULES[selected_module]()

if __name__== '__main__':
	main()
