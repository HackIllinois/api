#!/bin/bash

trap cleanup INT

STATUS=2
function cleanup {
  echo "Stopping services"
	pgrep "hackillinois" | xargs kill
	rm -rf log/
	exit $STATUS
}

REPO_ROOT="$(git rev-parse --show-toplevel)"

export HI_CONFIG=file://$REPO_ROOT/config/test_config.json
export BASE_PACKAGE=github.com/HackIllinois/api

mkdir log/
touch log/access.log

echo "Checking if another API is running on 8000 ...";
curl --silent --output /dev/null localhost:8000
if [ $? -eq 0 ]
then
	echo "Another API is running on port 8000. Please make sure to stop the process in order to run integration tests."
	exit $STATUS
fi

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

sleep 2

echo "Beginning integration tests";
echo "Checking if the API is running...";
curl --silent --output /dev/null localhost:8000 || (echo "Failed to connect to the API. Is it running? If it's not, start it with 'make run-test'"; exit 1;)
echo "Running end-to-end tests";
HI_CONFIG=file://$REPO_ROOT/config/test_config.json go test $BASE_PACKAGE/tests/e2e/... -v -count 1;
STATUS=$?

cleanup

