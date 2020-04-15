package binancedex

import (
	"context"
	"fmt"
	"os"
	"sort"
	"time"

	sdk "github.com/binance-chain/go-sdk/client"
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
	"github.com/tendermint/tendermint/rpc/client"
)

var(
 //example
	mnemonic := "lock globe panda armed mandate fabric couple dove climb step stove price recall decrease fire sail ring media enhance excite deny valid ceiling arm"
)

func client()(client sdk.DexClient){
	keyManager, _ := keys.NewMnemonicKeyManager(mnemonic)
	client, _ := sdk.NewDexClient("dex.binance.org", types.ProdNetwork, keyManager)
	return client
}


func adress( client sdk.DexClient, sym1 string, sym2 string) []string {
	Trades, _ := client.GetTrades(types.NewTradesQuery(true).WithSymbol(sym1, sym2).WithLimit(1000))

	var tab []string
	var tab2 []string
	var gt []float32
	var gs []string

	la := len(Trades.Trade)
	for i := 0; i < la; i++ {
		b := Trades.Trade[i].BuyerId
		c := Trades.Trade[i].SellerId
		d := Trades.Trade[i].TradeID
		if b == c {
			tab = append(tab, b)
			tab2 = append(tab2, d)
		} else {
			tab = append(tab, b)
			tab = append(tab, c)
			tab2 = append(tab2, d)
		}
	}
	uniqueSlice2 := unique(tab2)
	fmt.Println("Nombre de Trade", len(uniqueSlice2))
	uniqueSlice := unique(tab)

	for k := 0; k < len(uniqueSlice); k++ {
		a := countInArray(tab, uniqueSlice[k])
		a2 := float32(a)
		b2 := float32(len(tab))
		g := a2 / b2 * 100
		gt = append(gt, g)
		if g > 1 {
			gs = append(gs, uniqueSlice[k])
		}
	}

	fmt.Println("Nombre d'adresse", len(uniqueSlice))
	return gs
}

func Getordre(client sdk.DexClient, gA []string) (string, string) {
	OrdreB := make([]ordre, 0)
	OrdreS := make([]ordre, 0)
	for i := 0; i < len(gA); i++ {

		openOrders, _ := client.GetOpenOrders(types.NewOpenOrdersQuery(gA[i], true).WithSymbol("BTCB-1DE_BUSD-BD1"))

		for k := 0; k < len(openOrders.Order); k++ {
			if openOrders.Order[k].Side == 1 {
				Ordrei := make([]ordre, 1)
				Ordrei[0].Adresse = gA[i]
				Ordrei[0].Price = openOrders.Order[k].Price
				OrdreB = append(OrdreB, Ordrei...)
			} else {
				Ordres := make([]ordre, 1)
				Ordres[0].Adresse = gA[i]
				Ordres[0].Price = openOrders.Order[k].Price
				OrdreS = append(OrdreS, Ordres...)
			}
		}
	}
	sort.SliceStable(OrdreB, func(i, j int) bool {
		return OrdreB[i].Price < OrdreB[j].Price
	})
	sort.SliceStable(OrdreS, func(i, j int) bool {
		return OrdreS[i].Price < OrdreS[j].Price
	})
	ab := len(OrdreB)
	//as := len(OrdreS)
	//for k := 1; k >= 0; k-- {
	//fmt.Println("S", OrdreS[k].Price)
	//}

	//a, _ := strconv.ParseFloat(OrdreB[ab-1].Price, 64)
	//b, _ := strconv.ParseFloat(OrdreS[0].Price, 64)
	//c := b - a

	//fmt.Println(" srpead = ", c)
	//for k := ab - 1; k > ab-2; k-- {
	//fmt.Println("B", OrdreB[k].Price)
	//}
	Long := OrdreB[ab-1].Price
	Short := OrdreS[0].Price

	return Long, Short
}