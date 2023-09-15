package main

import (
	"fmt"
	"math/rand"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func main() {
	outputs := json.Marshal({
		"deadlock_output": {
		  "sensitive": false,
		  "type": [
			"tuple",
			[]
		  ],
		  "value": []
		},
		"idle_transactions_output": {
		  "sensitive": false,
		  "type": [
			"tuple",
			[]
		  ],
		  "value": []
		},
		"server_response_time_output": {
		  "sensitive": false,
		  "type": [
			"tuple",
			[
			  "string"
			]
		  ],
		  "value": [
			"130261141"
		  ]
		},
		"stored_proc_cache_hit_output": {
		  "sensitive": false,
		  "type": [
			"tuple",
			[
			  "string"
			]
		  ],
		  "value": [
			"130261138"
		  ]
		},
		"user_connections_output": {
		  "sensitive": false,
		  "type": [
			"tuple",
			[
			  "string"
			]
		  ],
		  "value": [
			"130261137"
		  ]
		}
	  }
	terraformOptions := &terraform.Options{

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"deadlock": map[string]interface{}{

				"critical_threshold": rand.Intn(20-6) + 6,
				"warning_threshold":  rand.Intn(6-0) + 0,
			},
			"stored_proc_cache_hit": map[string]interface{}{

				"critical_threshold": rand.Intn(20-10) + 10,
				"warning_threshold":  rand.Intn(80-21) + 21,
			},
			"idle_transactions": map[string]interface{}{

				"critical_threshold": rand.Intn(10-6) + 6,
				"warning_threshold":  rand.Intn(5-0) + 0,
			},
			// "rds_engine": engine_strings[index],
			// "rds_engine": "postgres",
			"server_response_time": map[string]interface{}{

				"critical_threshold": rand.Intn(1000-501) + 501,
				"warning_threshold":  rand.Intn(500-0) + 0,
			},
			"user_connections": map[string]interface{}{

				"critical_threshold": 15,
				"warning_threshold":  10,
			},
		},
	})
	for key := range terraformOptions.Vars {
		// fmt.Printf("%s is %s\n", key, value)
		monitorid := ("151515")
		critical_threshold := int64(terraformOptions.Vars[key].(map[string]interface{})["critical_threshold"].(int))
		warning_threshold := int64(terraformOptions.Vars[key].(map[string]interface{})["warning_threshold"].(int))

		fmt.Printf("Running Datadog Monitor check with the following values: %s, %s, %d, %d\n", key, monitorid, critical_threshold, warning_threshold)
	}

}
