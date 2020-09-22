#! /bin/bash

services=( auth checkin decision event mail notifications registration rsvp stat upload user )

if [ "$TRAVIS_PULL_REQUEST" != "false" ] ; then
    # get all changed files in PR
    changed_files=`git diff --name-only $TRAVIS_BRANCH...HEAD`

    # figure out which services might be a problem
    n_changed_services=0
    changed_services_list=""
    for service in ${services[*]} ; do
	service_upper=${service^}
	
	# check if any file in the service has been changed and if the doc page has not been changed
	if [[ $changed_files =~ "services/${service}" ]] \
	       && ! [[ $changed_files =~ "documentation/docs/reference/services/${service_upper}.md" ]]; then
	    n_changed_services=$(($n_changed_services + 1))
	    changed_services_list="${changed_services_list}- [${service_upper}](https://github.com/HackIllinois/api/blob/master/documentation/docs/reference/services/${service_upper}.md)\n"
	fi
    done

    pr_message="It looks like your PR makes changes to the following services, but does not include changes \
to the corresponding documentation:\n${changed_services_list}\nIf your changes do not require documentation \
updates, feel free to ignore this message."

    if (( $n_changed_services > 0 )); then
	echo "a service has changed"
	curl -H "Authorization: token $GITHUB_ACCESS_TOKEN" -X POST -d "{\"body\": \"${pr_message}\"}" "https://api.github.com/repos/\
${TRAVIS_REPO_SLUG}/issues/${TRAVIS_PULL_REQUEST}/comments"
    fi
fi
