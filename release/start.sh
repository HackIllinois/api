#!/bin/bash

# Start gateway and api services
./hackillinois-api-auth &
./hackillinois-api-user &
./hackillinois-api-registration &
./hackillinois-api-decision &
./hackillinois-api-rsvp &
./hackillinois-api-checkin &
./hackillinois-api-upload &
./hackillinois-api-mail &
./hackillinois-api-event &
./hackillinois-api-stat &
./hackillinois-api-notifications &
./hackillinois-api-gateway -u
