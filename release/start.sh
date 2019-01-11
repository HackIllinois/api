#!/bin/bash

# Start gateway and api services
./hackillinois-api --service auth &
./hackillinois-api --service user &
./hackillinois-api --service registration &
./hackillinois-api --service decision &
./hackillinois-api --service rsvp &
./hackillinois-api --service checkin &
./hackillinois-api --service upload &
./hackillinois-api --service mail &
./hackillinois-api --service event &
./hackillinois-api --service stat &
./hackillinois-api --service notifications &
./hackillinois-api --service gateway
