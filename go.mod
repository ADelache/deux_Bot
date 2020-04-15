module github.com/Binance

require (
	github.com/adshao/go-binance v0.0.0-20200326152909-7314295d8a33
	github.com/binance-chain/go-sdk v1.2.2
	github.com/dghubble/go-twitter v0.0.0-20190719072343-39e5462e111f // indirect
	github.com/dghubble/oauth1 v0.6.0 // indirect
)

replace github.com/tendermint/go-amino => github.com/binance-chain/bnc-go-amino v0.14.1-binance.1

go 1.14
