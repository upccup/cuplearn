#!/bin/sh


check_omega_agent() {
  if ps ax | grep -v grep | grep "omega-agent" > /dev/null
  then
    echo "Omega Agent service is running now... "
    echo "Wraning!!! Continue installation will overwrite the original version"
    while true; do
		echo "Are you sure to continue install [Y/N]"
		read answer
		case $answer in 
			Y|y|yes|YES)
			service omega-agent stop > /dev/null 2>&1
			break
			;;
			N|n|NO|no)
			exit 1
			;;
		esac
	done
  fi
}

check_omega_agent

