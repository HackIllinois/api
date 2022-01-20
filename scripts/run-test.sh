#!/bin/bash

trap cleanup INT

function cleanup {
  echo "Stopping services"
	pgrep "hackillinois" | xargs kill
	rm -rf log/
	exit 0
}

REPO_ROOT="$(git rev-parse --show-toplevel)"

export HI_CONFIG=file://$REPO_ROOT/config/test_config.json

mkdir log/
touch log/access.log

$REPO_ROOT/bin/hackillinois-api --service auth &
$REPO_ROOT/bin/hackillinois-api --service user &
$REPO_ROOT/bin/hackillinois-api --service registration &
$REPO_ROOT/bin/hackillinois-api --service decision &
$REPO_ROOT/bin/hackillinois-api --service rsvp &
$REPO_ROOT/bin/hackillinois-api --service checkin &
$REPO_ROOT/bin/hackillinois-api --service upload &
$REPO_ROOT/bin/hackillinois-api --service mail &
$REPO_ROOT/bin/hackillinois-api --service event &
$REPO_ROOT/bin/hackillinois-api --service stat &
$REPO_ROOT/bin/hackillinois-api --service notifications &
$REPO_ROOT/bin/hackillinois-api --service project &
$REPO_ROOT/bin/hackillinois-api --service profile &

$REPO_ROOT/bin/hackillinois-api --service gateway &

while true
	do sleep 60
done
