package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	testChar := `
     {
	    "allocator/event_queue_dispatches":0,
	    "master/cpus_percent":0,
	    "master/cpus_revocable_percent":0,
	    "master/cpus_revocable_total":0,
	    "master/cpus_revocable_used":0,
	    "master/cpus_total":0,
	    "master/cpus_used":0,
	    "childJson" : {
	    	"aaa" :0,
	    	"bbb" :"c",
	    	"childchJson2" : {
	    		"aasda": "bbbbb",
	    		"wwdsda": "sasdasda"
	    	},
	    	"frddasd": ["aaaa",1112313]
	    }
	}
   `

	byt := []byte(testChar)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(dat)

}
