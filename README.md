# Bonjour voici le deux_bot.
La stratégie du deux bot est simple 
Dans un premier temps nous observons le nombre d'adresse qui trade sur une monnaie binancedex.
# Adress Binance dex.
```golang
	Tabadress := binancedex.adress(clientdex, "BTCB-1DE", "BUSD-BD1")

```
Ensuite nous voyons les trades actuelles de ces adresses pour pourvoir déterminer notre long et short.
Nous avons deux moyens : -Websocket
                         - API
                      
# GetOrder.

```golang
Long, Short := Getordre(client, gA)

```
# Passer des ordres 

```golang
Lg, _ := futuresClient.NewCreateOrderService().Symbol("BTCUSDT").Side(futures.SideTypeBuy).Type(futures.OrderTypeLimit).Quantity("0.1")
.Price(Long).TimeInForce(futures.TimeInForceTypeGTC).Do(context.Background())

```
