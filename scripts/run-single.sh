#!/bin/bash

trap cleanup INT

function cleanup {
	echo "Cleaning up services"
	pgrep "hackillinois" | xargs kill
	rm -rf log/
	exit 0
}

REPO_ROOT="$(git rev-parse --show-toplevel)"

export HI_CONFIG=file://$REPO_ROOT/config/dev_config.json

mkdir log/
touch log/access.log

$REPO_ROOT/bin/hackillinois-api --service all &

sleep infinity
