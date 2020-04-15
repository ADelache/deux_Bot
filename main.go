package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/adshao/go-binance"
	"github.com/adshao/go-binance/futures"
	"github.com/dghubble/oauth1"
)

var (
	apiKey    = ""
	secretKey = ""

	Long2  string
	Short2 string
	Long   string
	Short  string
	TABL   []int
	L      int
	S      int
)

func twitter() {
	config := oauth1.NewConfig("", "")
	token := oauth1.NewToken("", "")
	httpClient := config.Client(oauth1.NoContext, token)
	clienttw := twitter.NewClient(httpClient)
	return clienttw
}

func main() {

	/////Initialisation
	f, _ := os.Create("test.txt")
	c, _ := os.Create("Ordre.txt")
	r, _ := os.Create("res.txt")

	futuresClient := binance.NewFuturesClient(apiKey, secretKey)
	clientdex =client()

	gA :=adress(clientdex, "BTCB-1DE", "BUSD-BD1")

	////Boucle

	for true {
		Long, Short = Getordre(client, gA)
		Lg, _ := futuresClient.NewCreateOrderService().Symbol("BTCUSDT").Side(futures.SideTypeBuy).Type(futures.OrderTypeLimit).Quantity("0.1").Price(Long).TimeInForce(futures.TimeInForceTypeGTC).Do(context.Background())
		Sl, _ := futuresClient.NewCreateOrderService().Symbol("BTCUSDT").Side(futures.SideTypeSell).Type(futures.OrderTypeLimit).Quantity("0.1").Price(Short).TimeInForce(futures.TimeInForceTypeGTC).Do(context.Background())
		fmt.Println(Lg)
		fmt.Println(Sl)
		Tw := "LONG : " + Long + "\nSHORT : " + Short + "\n#Bitcoin #BTC" + "\n#GoLang" + "\n#BinanceFutures " + "#BOT " + "#BOT_TRADING "
		clienttw.Statuses.Update(Tw, nil)
		fmt.Println("Le long est à ", Long)
		fmt.Println("Le short est à", Short)
		c.WriteString("Le long est à ")
		fmt.Fprintln(c, Long)
		c.WriteString("Le short est à ")
		fmt.Fprintln(c, Short)
		doneC := make(chan struct{})
		L = 0
		S = 0

		for true {
			Long2, Short2 = Getordre(client, gA)
			if Long2 >= Short {
				//openOrders, _ := futuresClient.NewListOpenOrdersService().Symbol("BTCUSDT").Do(context.Background())
				Lg2, _ := futuresClient.NewCreateOrderService().Symbol("BTCUSDT").Side(futures.SideTypeBuy).Type(futures.OrderTypeMarket).Quantity("0.1").Price(Short2).TimeInForce(futures.TimeInForceTypeGTC).Do(context.Background())
				_, err := futuresClient.NewCancelOrderService().Symbol("BTCUSDT").OrderID(Lg.OrderID).Do(context.Background())
				fmt.Println(Lg2)
				fmt.Println(err)
				fmt.Println("Perdu")
				f.WriteString("Perdu //")
				TABL = append(TABL, 0)
				b :=countInArrayI(TABL, 0)
				r.WriteString("Nombre de trade raté ")
				fmt.Fprintln(r, b)
				c :=countInArrayI(TABL, 1)
				r.WriteString("Nombre de trade réussit ")
				fmt.Fprintln(r, c)
				time.Sleep(300 * time.Second)
				break
			}

			if Short2 <= Long {
				//Faire annuler l'ordre
				//openOrders, _ := futuresClient.NewListOpenOrdersService().Symbol("BTCUSDT").Do(context.Background())
				St2, _ := futuresClient.NewCreateOrderService().Symbol("BTCUSDT").Side(futures.SideTypeSell).Price(Long2).Type(futures.OrderTypeLimit).Quantity("0.1").TimeInForce(futures.TimeInForceTypeGTC).Do(context.Background())
				_, err := futuresClient.NewCancelOrderService().Symbol("BTCUSDT").OrderID(Sl.OrderID).Do(context.Background())
				fmt.Println(St2)
				fmt.Println(err)
				fmt.Println("Perdu")
				f.WriteString("Perdu //")
				TABL = append(TABL, 0)
				b :=countInArrayI(TABL, 0)
				r.WriteString("Nombre de trade raté")
				fmt.Fprintln(r, b)
				c :=countInArrayI(TABL, 1)
				r.WriteString("Nombre de trade réussit ")
				fmt.Fprintln(r, c)
				time.Sleep(300 * time.Second)
				break
			}

			wsTradeHandler := func(event *binance.WsTradeEvent) {
				if L == 0 && S == 0 {
					fmt.Println("Le Prix ", event.Price)
				}

				if event.Price <= Long && L == 0 {
					fmt.Println("L'ordre de Long est passé", event.Price)
					L = 1
					c.WriteString("L'ordre long est passé au prix")
					fmt.Fprintln(c, event.Price)
				}
				if Short <= event.Price && S == 0 {
					fmt.Println("L'odre de short est passé", event.Price)
					S = 1
					c.WriteString("L'ordre Short est passé au prix")
					fmt.Fprintln(c, event.Price)
				}

				if L == 0 && S == 1 {
					fmt.Println("Prix ", event.Price)
					fmt.Println("Prix de vente Short ", Long)
					fmt.Println("Prix actuelle de la Perte (Long) ", Long2)
					if event.Price <= Long {
						S = 1
						doneC <- struct{}{}
					}
				}

				if L == 1 && S == 0 {
					fmt.Println("Prix ", event.Price)
					fmt.Println("Prix de vente Long ", Short)
					fmt.Println("Prix actuelle de la Perte (Short) ", Short2)
					if event.Price >= Short {
						S = 1
						doneC <- struct{}{}
					}
				}

			}
			errHandler := func(err error) {
				fmt.Println(err)
			}

			doneC, stopC, _ := binance.WsTradeServe("BTCUSDT", wsTradeHandler, errHandler)
			go func() {
				time.Sleep(20 * time.Second)
				stopC <- struct{}{}
			}()

			<-doneC

			if S == 1 && L == 1 {
				TABL = append(TABL, 1)
				a :=countInArrayI(TABL, 1)
				b :=countInArrayI(TABL, 0)
				r.WriteString("Nombre de trade réussit ")
				fmt.Fprintln(r, a)
				r.WriteString("Nombre de trade raté ")
				fmt.Fprintln(r, b)
				f.WriteString("Gagné //")
				time.Sleep(160 * time.Second)
				break
			}

		}

	}
}
