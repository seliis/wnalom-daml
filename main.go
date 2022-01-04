package main

import (
	"fmt"
	"log"
)

/*
	** IMPORTANT! **
	Maximum datas of a request is 500. So, you must set calculation result less than 500 per request.
	when interval time scale is minutes, duration = (interval * 500) / 60
*/

var (
	API    = "lWTzykapQbvi7kSU6OYHuEYVvbc9f9Bi5fidxJmvYpwFvnDlLPP42fMcgtaPLDtR"
	SECRET = "IbAZ82BeiFAWtBzTQtefn9FDaVnd3mkOdQXJMkBf5tMeWIMfcDA9Mn6qWvQfbPpM"
	/*
		Genesis time is targetted cryptocurrency listed date.
		Format of this time must be YYYY-MM-DD
	*/
	CRYPTO_NAME    = "BTCUSDT"
	CRYPTO_GENESIS = "2019-09-08"
	/*
		Go-Lang time.ParseDuration Scales
		NanoSecond : ns
		MicroSecond: us
		MilliSecond: ms
		Second     : s
		Minute     : m
		Hour       : h

		if interval is minutes scale,
		duration = (interval * 500) / 60
	*/
	TIME_DURATION = "41.66666h"
	/*
		Binance Futures Intervals (23 Dec 2021)
		Minutes : 1m, 3m, 5m, 15m, 30m
		Hours   : 1h, 2h, 4h, 6h, 8h, 12h
		Days    : 1d, 3d
		Weeks   : 1w
		Months  : 1M
	*/
	TIME_INTERVAL = "5m"
	/*
		Switch of Indicators
	*/
	INDICATOR_SWITCH = map[string]bool{
		"MACD":      true,
		"BOLLINGER": true,
		"RSI":       false,
	}
)

func main() {
	if err := GetNewKlines(GetFileName(), CRYPTO_NAME, GetGenesisTime(), TIME_DURATION, TIME_INTERVAL); err != nil {
		log.Fatal(err)
	}

	if err := SetIndicators("High", "Low", "Close"); err != nil {
		if err.Error() == "no_selected_indicator" {
			fmt.Println("NOTE: No Selected Indicator")
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println("Complete")
}
