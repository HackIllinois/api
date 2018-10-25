#!/bin/bash

trap cleanup INT

function cleanup {
	echo "Cleaning up services"
	pgrep "hackillinois-" | xargs kill
	rm -rf log/
	exit 0
}

function auth {
	export GITHUB_CLIENT_ID=
	export GITHUB_CLIENT_SECRET=
	export GOOGLE_CLIENT_ID=
	export GOOGLE_CLIENT_SECRET=
	export LINKEDIN_CLIENT_ID=
	export LINKEDIN_CLIENT_SECRET=
	export TOKEN_SECRET=secret_string
	export AUTH_REDIRECT_URI=https://hackillinois.org/auth/
	export AUTH_DB_HOST=localhost
	export AUTH_DB_NAME=auth
	export AUTH_PORT=:8002
	export USER_SERVICE=http://localhost:8003

	hackillinois-api-auth &
}

function user {
	export USER_DB_HOST=localhost
	export USER_DB_NAME=user
	export USER_PORT=:8003

	hackillinois-api-user &
}

function registration {
	export REGISTRATION_DB_HOST=localhost
	export REGISTRATION_DB_NAME=registration
	export REGISTRATION_PORT=:8004
	export AUTH_SERVICE=http://localhost:8002
	export USER_SERVICE=http://localhost:8003
	export DECISION_SERVICE=http://localhost:8005
	export MAIL_SERVICE=http://localhost:8009

	hackillinois-api-registration &
}

function decision {
	export DECISION_DB_HOST=localhost
	export DECISION_DB_NAME=decision
	export DECISION_PORT=:8005
	export MAIL_SERVICE=http://localhost:8009

	hackillinois-api-decision &
}

function rsvp {
	export RSVP_DB_HOST=localhost
	export RSVP_DB_NAME=rsvp
	export RSVP_PORT=:8006
	export AUTH_SERVICE=http://localhost:8002
	export DECISION_SERVICE=http://localhost:8005
	export MAIL_SERVICE=http://localhost:8009

	hackillinois-api-rsvp &
}

function checkin {
	export CHECKIN_DB_HOST=localhost
	export CHECKIN_DB_NAME=checkin
	export CHECKIN_PORT=:8007
	export RSVP_SERVICE=http://localhost:8006
	export REGISTRATION_SERVICE=http://localhost:8004

	hackillinois-api-checkin &
}

function upload {
	export UPLOAD_PORT=:8008
	export S3_REGION=us-east-1
	export S3_BUCKET=hackillinois-upload-2019

	hackillinois-api-upload &
}

function mail {
	export SPARKPOST_API=https://api.sparkpost.com/api/v1
	export MAIL_DB_HOST=localhost
	export MAIL_DB_NAME=mail
	export USER_SERVICE=http://localhost:8003
	export MAIL_PORT=:8009

	hackillinois-api-mail &
}

function event {
	export EVENT_DB_HOST=localhost
	export EVENT_DB_NAME=event
	export EVENT_PORT=:8010
	export CHECKIN_SERVICE=http://localhost:8007

	hackillinois-api-event &
}

function stat {
	export STAT_DB_HOST=localhost
	export STAT_DB_NAME=stat
	export STAT_PORT=:8011

	hackillinois-api-stat &
}

function notifications {
	export NOTIFICATIONS_DB_HOST=localhost
	export NOTIFICATIONS_DB_NAME=notifications
	export NOTIFICATIONS_PORT=:8012
	export SNS_REGION=us-east-2

	hackillinois-api-notifications &
}

function gateway {
	export GATEWAY_PORT=8000
	export TOKEN_SECRET=secret_string
	export AUTH_SERVICE=http://localhost:8002
	export USER_SERVICE=http://localhost:8003
	export REGISTRATION_SERVICE=http://localhost:8004
	export DECISION_SERVICE=http://localhost:8005
	export RSVP_SERVICE=http://localhost:8006
	export CHECKIN_SERVICE=http://localhost:8007
	export UPLOAD_SERVICE=http://localhost:8008
	export MAIL_SERVICE=http://localhost:8009
	export EVENT_SERVICE=http://localhost:8010
	export STAT_SERVICE=http://localhost:8011
	export NOTIFICATIONS_SERVICE=http://localhost:8012

	mkdir log/
	touch log/access.log

	hackillinois-api-gateway -u &
}

auth
user
registration
decision
rsvp
checkin
upload
mail
event
stat
notifications
gateway

sleep infinity
