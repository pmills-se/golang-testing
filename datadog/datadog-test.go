package test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

var monitor_ids string
monitor_ids := []

func TestingDatadog() {
	for _, element := range monitor_ids {
		MonitorID := int64(element)
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
	}
}
