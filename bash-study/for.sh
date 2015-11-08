#!/bin/sh                                                                                                                                                                                                   


check_omega_agent() {
      if ps ax | grep -v grep | grep "omega-agent" > /dev/null
        then
        echo "Omega Agent service is running now... "
        echo "Wraning!!! Continue installation will overwrite the original version"
        for ((i=10;i>=0;--i))
			do 
			    echo "new omega-agent will install after ${i}s" 
			    sleep 1s
			done
	   fi
}

check_omega_agent
                       
