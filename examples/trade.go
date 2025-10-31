package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Benjmmi/okx"
	"github.com/Benjmmi/okx/api"
	"github.com/Benjmmi/okx/api/rest"
	"github.com/Benjmmi/okx/models/config"
	"log"
)

func main() {

	cfg, _ := config.LoadConfig("./config/config_test.json")

	var dest okx.Destination
	if cfg.Demo {
		dest = okx.DemoServer
	} else if cfg.Colo {
		dest = okx.ColoServer
	} else {
		dest = okx.NormalServer
	}

	ctx := context.Background()

	var client *api.Client
	var err error
	client, err = api.NewClient(ctx, cfg.OkxAPIKey, cfg.OkxSecretKey, cfg.OkxPassword, dest)

	if err != nil {
		log.Fatal(err)
		//return false
		return
	}
	tradeClient := rest.NewTrade(client.Rest)
	accountRateLimit, err := tradeClient.GetAccountRateLimit()
	if err != nil {
		fmt.Println("Error fetching order history:", err)
		return
	}
	accountRateLimitJSON, err := json.MarshalIndent(accountRateLimit, "", "  ")
	if err != nil {
		fmt.Println("Error converting order history to JSON:", err)
		return
	}
	fmt.Println("Account Rate Limit:", string(accountRateLimitJSON))
}
