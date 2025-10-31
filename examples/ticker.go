package main

import (
	"context"
	"fmt"
	"github.com/Benjmmi/okx"
	"github.com/Benjmmi/okx/api"
	"github.com/Benjmmi/okx/events/public"
	"github.com/Benjmmi/okx/models/config"
	wsRequestPublic "github.com/Benjmmi/okx/requests/ws/public"
	"log"
)

func main() {
	cfg, _ := config.LoadConfig("./config/config_test.json")
	// 监听okx行情信息
	instID := "USDT-USDC"
	okxSpotTickerChan := make(chan *public.Tickers)
	ctx := context.Background()
	dest := okx.DemoServer
	client, err := api.NewClient(ctx, cfg.OkxAPIKey, cfg.OkxSecretKey, cfg.OkxPassword, dest)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ws.Public.Tickers(wsRequestPublic.Tickers{
		InstID: instID,
	}, okxSpotTickerChan)
	if err != nil {
		log.Fatal(err)
	}
	for {
		s := <-okxSpotTickerChan
		for _, t := range s.Tickers {
			fmt.Println(t)
		}
	}
}
