package record

import (
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

type Record struct {
	ID            int64
	TimeStamp     time.Time
	Latency       time.Duration
	UploadSpeed   float64
	DownloadSpeed float64
	Distance      float64
	PingOk        bool
}

func Now() (Record, error) {

	user, e := speedtest.FetchUserInfo()
	if e != nil {
		return Record{}, e
	}
	serverList, e := speedtest.FetchServers(user)
	if e != nil {
		return Record{}, e
	}
	targets, e := serverList.FindServer([]int{})
	if e != nil {
		return Record{}, e
	}

	var record Record

	for _, s := range targets {
		pingError := s.PingTest()
		upErr := s.UploadTest(false)
		if upErr != nil {
			return Record{}, upErr
		}

		downErr := s.DownloadTest(false)
		if downErr != nil {
			return Record{}, downErr
		}

		record = Record{
			TimeStamp:     time.Now(),
			Latency:       s.Latency,
			UploadSpeed:   s.ULSpeed,
			DownloadSpeed: s.DLSpeed,
			Distance:      s.Distance,
			PingOk:        pingError == nil,
		}
	}
	return record, nil
}
