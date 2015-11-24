#!/bin/sh

set -e

select_iface()
{
	echo "Omega-agent use default network interface is etho."
	echo "Do you want to change it? [Y/N]"
	if read -t 5 change_ifcae 
		then
		case $change_ifcae  in 
			Y|y|YES|yes)
			while true; do 
				echo "Please choose network interface from below list: "
				echo `ls /sys/class/net/`
				read iface
				check_cmd="ls /sys/class/net/${iface}"
				if ${check_cmd} > /dev/null
					then
					break
				else
					echo "Network interface ${iface} not find"
				fi
			done
			;;
			N|n|NO|No|*)
			echo "Network interface use eth0"
			;;
		esac
	else 
		echo "Network interface use eth0"
	fi
}

select_iface