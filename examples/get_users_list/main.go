package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	om "github.com/grokify/oauth2more"

	"github.com/zocollo/go-slack/client"
)

func main() {
	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		panic(err)
	}

	httpClient := om.NewClientBearerTokenSimple(
		os.Getenv("SLACK_BOT_TOKEN"),
	)

	ctx := context.Background()
	emptyMSI := map[string]interface{}{}

	apiConfig := slack.NewConfiguration()
	apiConfig.BasePath = "https://slack.com/api"
	apiConfig.HTTPClient = httpClient
	apiClient := slack.NewAPIClient(apiConfig)

	info, resp, err := apiClient.UsersApi.UsersList(ctx, emptyMSI)
	if err != nil {
		panic(err)
	} else if resp.StatusCode >= 300 {
		panic(fmt.Sprintf("Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}
