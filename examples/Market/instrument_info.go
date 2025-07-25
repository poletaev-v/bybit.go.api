package main

import (
	"context"
	"fmt"
	bybit "github.com/poletaev-v/bybit.go.api"
)

func main() {
	GetInstrumentInfo()
}

func GetInstrumentInfo() {
	client := bybit.NewBybitHttpClient("", "", bybit.WithBaseURL(bybit.TESTNET), bybit.WithDebug(true))
	params := map[string]interface{}{"category": "spot", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetInstrumentInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(serverResult))
}
