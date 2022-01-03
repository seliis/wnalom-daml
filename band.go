package main

import (
	"errors"
	"time"
)

type TimeBand struct {
	Start int64
	End   int64
}

func GetHumanTime(unix int64) string {
	return time.UnixMilli(unix).Format(time.RFC3339)
}

func GetTimeBands(genesis string, duration string, interval string) (*[]TimeBand, error) {
	unixGenesis, err := GetUnixGenesis(genesis)
	if err != nil {
		return nil, err
	}

	unixDuration, err := GetUnixDuration(duration)
	if err != nil {
		return nil, err
	}

	unixNow := GetUnixNow()

	responseBool, err := GetConfirm()
	if err != nil {
		return nil, err
	}

	if responseBool {
		var timeBands []TimeBand
		for i := unixGenesis; i < unixNow; i += unixDuration {
			timeBand := new(TimeBand)
			timeBand.Start = i
			timeBand.End = unixDuration + i
			timeBands = append(timeBands, *timeBand)
		}
		return &timeBands, nil
	} else {
		return nil, errors.New("terminated")
	}
}

func GetUnixNow() int64 {
	return time.Now().UnixMilli()
}

func GetUnixInterval(interval string) (int64, error) {
	if timeInterval, err := time.ParseDuration(interval); err != nil {
		return 0, err
	} else {
		return timeInterval.Milliseconds(), nil
	}
}

func GetUnixGenesis(genesis string) (int64, error) {
	if timeGenesis, err := time.Parse(time.RFC3339, genesis); err != nil {
		return 0, err
	} else {
		return timeGenesis.UnixMilli(), nil
	}
}

func GetUnixDuration(duration string) (int64, error) {
	if timeDuration, err := time.ParseDuration(duration); err != nil {
		return 0, err
	} else {
		return timeDuration.Milliseconds(), nil
	}
}
