package timet

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	t.Log(now)
	t.Log(now.Format("2006-01-02 15:04:05"))
	t.Log(now.Format("2006-1-2 15:4:5"))
	t.Log(now.Date())
	t.Log(now.Clock())
	t.Log(now.Year())
	t.Log(now.Month().String())
	t.Log(now.Day())
	t.Log(now.YearDay())
	t.Log(now.Hour())
	t.Log(now.Minute())
	t.Log(now.Second())
	t.Log(now.Weekday().String())
	t.Log(now.Unix())
	t.Log(now.UnixNano())
	t.Log(now.Nanosecond())
}

func TestLocation(t *testing.T) {
	now := time.Now()
	t.Log(now.Format("2006-01-02 15:04:05"))
	t.Log(now.Location().String())
	t.Log(now.UTC().Location().String())
	t.Log(now.In(now.UTC().Location()).Format("2006-01-02 15:04:05"))

	loc := time.FixedZone("UTC", 8*3600)
	t.Log(now.In(loc).Format("2006-01-02 15:04:05"))

	location, _ := time.LoadLocation("Local")
	t.Log(now.In(location).Format("2006-01-02 15:04:05"))

}

func TestTimeCal(t *testing.T) {
	now := time.Now()
	t.Log(now.Format("2006-01-02 15:04:05"))

	duration, _ := time.ParseDuration("1h")
	afterNow := now.Add(duration)
	t.Log(afterNow.Format("2006-01-02 15:04:05"))

	d, _ := time.ParseDuration("-1h")
	beforNow := now.Add(d)
	t.Log(beforNow.Format("2006-01-02 15:04:05"))

	loc := time.FixedZone("UTC", 8*3600)
	monthStartdayTime := time.Date(now.Year(), now.Month(), 0, 24, 0, 0, 0, loc)
	t.Log(monthStartdayTime.Zone())
	t.Log(monthStartdayTime.Format(time.RFC3339))

	t.Log(monthStartdayTime.IsZero())
	t.Log(now.Equal(monthStartdayTime))
	subDuration := now.Sub(monthStartdayTime)
	t.Log(subDuration.String())
	t.Log(subDuration.Seconds())

	t.Log(now.Unix() - monthStartdayTime.Unix())
}

func TestTimer(t *testing.T) {
	timer := time.NewTimer(time.Second * 2)
	defer timer.Stop()

	tt := time.After(time.Second * 5)
	var start bool = true
	for start {
		select {
		case <-tt:
			fmt.Println("over")
			start = false
		case nowTime := <-timer.C:
			fmt.Println(nowTime.Format(time.RFC3339))
			timer.Reset(time.Second * 2)
		}
	}

}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for {
		select {
		case nowTime := <-ticker.C:
			fmt.Println(nowTime.Format(time.RFC3339))
		case nowTime := <-time.Tick(time.Second):
			fmt.Println(nowTime.Format(time.RFC3339Nano))
		}
	}
}
