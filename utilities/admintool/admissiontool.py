#!/usr/bin/env python3

import os
import random
import json

import api

from registration_score import get_registration_score

def get_pending_applicants():
	data, success = api.make_request('GET', '/decision/filter/?status=PENDING')

	if not success:
		print('Failed to retreive decisions')
		return {}

	return [decision['id'] for decision in data['decisions']]

def get_registrations():
	data, success = api.make_request('GET', '/registration/filter/')

	if not success:
		print('Failed to retreive registrations')
		return {}

	return {registration['id'] : registration for registration in data['registrations']}

def get_applicant_score(registration):
	return get_registration_score(registration) + random.random(0.0, 1.0)

def get_top_applicants(count):
	applicants = get_pending_applicants()
	registrations = get_registrations()

	applicant_scores = []

	for applicant in applicants:
		applicant_scores.append((applicant, get_applicant_score(registrations[applicant])))

	applicant_scores.sort(lambda x: x[1], reverse = True)
	applicant_scores = applicant_scores[0 : count]

	return [applicant[0] for applicant in applicant_scores]

def admit_applicant(user_id, wave):
	decision = json.dumps({
		'id': user_id,
		'status': 'ACCEPTED',
		'wave': wave
	})

	data, success = api.make_request('POST', '/decision/', data = decision)

	if not success:
		print('Failed to update user decision')
		return

	finalization = json.dumps({
		'id': user_id,
		'finalized': True
	})

	data, success = api.make_request('POST', '/decision/finalize/', data = finalization)

	if not success:
		print('Failed to finalize user decision')
		return

	print('Admitted applicant')

def admit_top_applicants(count, wave):
	top_applicants = get_top_applicants(count)

	for applicant in top_applicants:
		admit_applicant(applicant, wave)

	print('Finished admitting top {} applicants'.format(count))

def main():
	if os.environ.get('HI_AUTH') == None:
		print('Must set HI_AUTH environment variable')
		return

	print('Enter number of applicants to admit')
	count = int(input('>'))

	print('Enter the wave to admit under')
	wave = int(input('>'))

	admit_top_applicants(count, wave)

if __name__== '__main__':
	main()
