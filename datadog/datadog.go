package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {

	CheckDatadogMonitor("132929230")

}

func CheckDatadogMonitor(MonitorOutput string) {
	MonitorID, _ := strconv.ParseInt(MonitorOutput, 10, 64)
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewMonitorsApi(apiClient)
	resp, r, err := api.GetMonitor(ctx, MonitorID, *datadogV1.NewGetMonitorOptionalParameters().WithWithDowntimes(true))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.GetMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.GetMonitor`:\n%s\n", responseContent)

	var response map[string]interface{}

	json.Unmarshal(responseContent, &response)

	formatted_string := response["options"].(map[string]interface{})["thresholds"].(map[string]interface{})["critical"]
	str := fmt.Sprintf("%v", formatted_string)
	parsed_critical_value := strings.Trim(strings.TrimPrefix(strings.TrimSuffix(str, "%"), "%!s(float64="), "=")
	var converted_critical_value int64
	parsedFloat, _ := strconv.ParseFloat(parsed_critical_value, 64)
	parsedFloat = parsedFloat * 100
	converted_critical_value = int64(parsedFloat)
	fmt.Print(converted_critical_value)

	// float, _ := strconv.ParseFloat(str, 64)
	// fmt.Print(float * 100)
	// fmt.Fprintf(os.Stdout, "parsed_critical_value: %s\n", parsed_critical_value)

}
