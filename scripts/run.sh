#!/bin/bash

trap cleanup INT

function cleanup {
	echo "Cleaning up services"
	pgrep "hackillinois-" | xargs kill
	rm -rf log/
	exit 0
}

export HI_CONFIG=file://${GOPATH}/src/github.com/HackIllinois/api/config/dev_config.json

mkdir log/
touch log/access.log

hackillinois-api-auth &
hackillinois-api-user &
hackillinois-api-registration &
hackillinois-api-decision &
hackillinois-api-rsvp &
hackillinois-api-checkin &
hackillinois-api-upload &
hackillinois-api-mail &
hackillinois-api-event &
hackillinois-api-stat &
hackillinois-api-notifications &

hackillinois-api-gateway -u &

sleep infinity
