package main

import (
	"strconv"
	"time"
	"tools/util"
)

const (
	TimeFormatA      = "2006-01-02 15:04:05" //时间格式：全
	TimeFormatB      = "2006-01-02"          //时间格式：全
	TimeFormatC      = "20060102150405"      //时间格式：全
	TimeFormatSecond = "20060102150405"      //时间格式：全
	TimeFormatMinute = "200601021504"        //时间格式：全
	TimeFormatHour   = "2006010215"          //时间格式：全
	TimeFormatDay    = "20060102"            //时间格式：全
	TimeFormatMonth  = "200601"              //时间格式：全
	TimeFormatYear   = "2006"                //时间格式：全
)

func getNowTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func time2Date(val string) string {
	timestamp, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return "解析失败"
	}
	now := time.Unix(timestamp, 0)
	res := now.Format(TimeFormatA)
	return res
}

func date2Time(val string) string {
	str := util.GetNum(val)
	if str == "" {
		return "解析失败"
	}
	layout := ""
	switch len(str) {
	case 14:
		layout = TimeFormatSecond
	case 12:
		layout = TimeFormatMinute
	case 10:
		layout = TimeFormatHour
	case 8:
		layout = TimeFormatDay
	case 6:
		layout = TimeFormatMonth
	case 4:
		layout = TimeFormatYear
	default:
		return "解析失败"
	}

	times, err := time.ParseInLocation(layout, str, time.Local)
	if err != nil {
		return "解析失败"
	}
	return strconv.FormatInt(times.Unix(), 10)
}
