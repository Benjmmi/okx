package main

import (
	"context"
	"fmt"
	"github.com/Benjmmi/okx/models/config"
	"log"

	"github.com/Benjmmi/okx"
	"github.com/Benjmmi/okx/api"
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
	if len(cfg.SourceIPs) == 0 || cfg.SourceIPs[0] == "" {
		client, err = api.NewClient(ctx, cfg.OkxAPIKey, cfg.OkxSecretKey, cfg.OkxPassword, dest)
	} else if len(cfg.TargetIPs) == 0 || cfg.TargetIPs[0] == "" {
		client, err = api.NewClientWithIP(ctx, cfg.OkxAPIKey, cfg.OkxSecretKey, cfg.OkxPassword, dest, cfg.SourceIPs[0])
	} else {
		client, err = api.NewClientWithSourceAndTargetIP(ctx, cfg.OkxAPIKey, cfg.OkxSecretKey, cfg.OkxPassword, dest, cfg.SourceIPs[0], cfg.TargetIPs[0])
	}
	if err != nil {
		log.Fatal(err)
		//return false
		return
	}

	err = client.Ws.Private.Connect(true)
	if err != nil {
		fmt.Printf("[StartOkxClientPoolForTargetIPs] Fail To Connect WS, Error: %s \n", err.Error())
	} else {
		fmt.Printf("[StartOkxClientPoolForTargetIPs] Connect WS Success for ip %s %s \n", cfg.SourceIPs[0], cfg.TargetIPs[0])
	}

	err = client.Ws.Private.Login()
	if err != nil {
		fmt.Printf("[StartOkxClientPoolForTargetIPs] Fail To Login WS, Error: %s \n", err.Error())
	} else {
		fmt.Printf("[StartOkxClientPoolForTargetIPs] Login WS Success for ip %s %s \n", cfg.SourceIPs[0], cfg.TargetIPs[0])
	}

}
