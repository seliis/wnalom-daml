package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"time"

	"github.com/adshao/go-binance/v2/futures"
)

func GetNewKlines(name string, symbol string, genesis string, duration string, interval string) error {
	file, writer, err := GetNewCSV(name)
	if err != nil {
		return err
	}

	expectedRequest, err := PrintExpectation(genesis, duration, interval)
	if err != nil {
		return err
	}

	timeBands, err := GetTimeBands(genesis, duration, interval)
	if err != nil {
		return err
	}

	writer.Write([]string{
		"OpenTime",
		"CloseTime",
		"Open",
		"Close",
		"High",
		"Low",
		"TradeNum",
		"Volume",
		"QuoteAssetVolume",
		"TakerBuyBaseAssetVolume",
		"TakerBuyQuoteAssetVolume",
	})

	futuresClient := futures.NewClient(API, SECRET)
	requestCount := 0.0

	for _, timeBand := range *timeBands {
		data, err := futuresClient.NewKlinesService().Symbol(symbol).Interval(interval).StartTime(timeBand.Start).EndTime(timeBand.End).Do(context.Background())
		if err != nil {
			return err
		}
		requestCount++
		WriteKlineToCSV(data, writer)
		time.Sleep(105 * time.Millisecond)
		fmt.Printf("\rProgress Status: %0.2f Percent\t", (requestCount/expectedRequest)*100)
	}

	defer func() {
		writer.Flush()
		file.Close()
	}()

	return nil
}

func WriteKlineToCSV(datas []*futures.Kline, writer *csv.Writer) error {
	for _, data := range datas {
		if err := writer.Write([]string{
			GetHumanTime(data.OpenTime),
			GetHumanTime(data.CloseTime),
			data.Open, data.Close,
			data.High, data.Low,
			fmt.Sprint(data.TradeNum), data.Volume,
			data.QuoteAssetVolume,
			data.TakerBuyBaseAssetVolume,
			data.TakerBuyQuoteAssetVolume,
		}); err != nil {
			return err
		}
	}
	return nil
}
