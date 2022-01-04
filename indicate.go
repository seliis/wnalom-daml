package main

import (
	"errors"
	"os"

	"github.com/cinar/indicator"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func SetIndicators(nameHigh, nameLow, nameClose string) error {
	if err := checkIndicators(); err != nil {
		return err
	}

	newFile, err := os.Create(CSV_FILE_NAME + "_indicated.csv")
	if err != nil {
		return err
	}

	file, err := os.Open(CSV_FILE_NAME + ".csv")
	if err != nil {
		return err
	}

	master := dataframe.ReadCSV(file)

	if INDICATOR_SWITCH["MACD"] {
		GetMACD(&master, nameClose)
	}

	if INDICATOR_SWITCH["BOLLINGER"] {
		GetBollingerBands(&master, nameClose)
	}

	if INDICATOR_SWITCH["RSI"] {
		GetRSI(&master, nameClose)
	}

	master.WriteCSV(newFile)

	defer func() {
		newFile.Close()
		file.Close()
	}()

	return nil
}

func checkIndicators() error {
	for _, v := range INDICATOR_SWITCH {
		if v {
			return nil
		}
	}
	return errors.New("no_selected_indicator")
}

func GetMACD(master *dataframe.DataFrame, nameClose string) {
	macd, signal := indicator.Macd(master.Col(nameClose).Float())

	series1 := series.New(macd, series.Float, "MacdMacd")
	series2 := series.New(signal, series.Float, "MacdSignal")

	*master = master.CBind(dataframe.New(series1, series2))
}

func GetBollingerBands(master *dataframe.DataFrame, nameClose string) {
	middle, upper, lower := indicator.BollingerBands(master.Col(nameClose).Float())

	series1 := series.New(middle, series.Float, "BollingerMiddle")
	series2 := series.New(upper, series.Float, "BollingerUpper")
	series3 := series.New(lower, series.Float, "BollingerLower")

	*master = master.CBind(dataframe.New(series1, series2, series3))
}

func GetRSI(master *dataframe.DataFrame, nameClose string) {
	rs, rsi := indicator.Rsi(master.Col(nameClose).Float())

	series1 := series.New(rs, series.Float, "RsiRs")
	series2 := series.New(rsi, series.Float, "RsiRsi")

	*master = master.CBind(dataframe.New(series1, series2))
}
