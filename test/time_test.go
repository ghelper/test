package test

import (
	"fmt"
	"testing"
	"time"
)

// GetYesterdayTime 返回昨天零点的时间戳
func GetYesterdayStartUnix() int64 {
	location, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(location).AddDate(0,0,-1)

	// now.Year(), now.Month(), now.Day() 是以本地时区为参照的年、月、日
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	yesterdayUnix := today.Unix()
	return yesterdayUnix
}

// GetYesterdayTime 返回昨天零点的时间戳
func GetYesterdayEndUnix() int64 {
	location, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(location).AddDate(0,0,-1)
	// now.Year(), now.Month(), now.Day() 是以本地时区为参照的年、月、日
	today := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, location)
	yesterdayUnix := today.Unix()
	return yesterdayUnix
}

func TestXin_Run(t *testing.T) {
	fmt.Println(GetYesterdayStartUnix())
	fmt.Println(GetYesterdayEndUnix())
}


func TestXin1_Run(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Shanghai")
	currentTime := time.Now().In(location).AddDate(0,0,-1)
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	fmt.Println(startTime)
	fmt.Println(startTime.Format("2006/01/02 15:04:05"))
}

func TestXin2_Run(t *testing.T) {
	currentTime := time.Now()
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
	fmt.Println(endTime)
	fmt.Println(endTime.Format("2006/01/02 15:04:05"))
}