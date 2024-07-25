package main

import (
	"context"
	"fmt"
	"github.com/LIJI-Max/okx"
	"github.com/LIJI-Max/okx/api"
	"github.com/LIJI-Max/okx/events/public"
	wsRequestPublic "github.com/LIJI-Max/okx/requests/ws/public"
	"log"
)

func main() {
	// 监听okx行情信息
	instID := "USDT-USDC"
	okxSpotTickerChan := make(chan *public.Tickers)
	ctx := context.Background()
	dest := okx.DemoServer
	client, err := api.NewClient(ctx, "42198db1-f9d2-4884-89f7-5f913924a41c", "5A5C4CA179FD4A32000C31C0EA751717", "wosiSMA1)", dest)
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
