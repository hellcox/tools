package main

import (
	"strconv"
	"time"
	"wails-vue/util"
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

type Time struct {
	NowTime int64  `json:"nowTime"`
	NowDate string `json:"nowDate"`
	Input   string `json:"input"`
	Output  string `json:"output"`
	Action  string `json:"action"`
}

type response struct {
	code int
	msg  string
	data struct {
		in  string
		out string
	}
}

func NewTime() *Time {
	return &Time{}
}

func (t *Time) GetNowTime() interface{} {
	timestamp := time.Now().Unix()
	date := time.Now().Format(TimeFormatA)

	res := map[string]interface{}{
		"timestamp": timestamp,
		"date":      date,
	}
	return res
}

func (t *Time) TimeToDate(val string) interface{} {
	res := map[string]interface{}{}
	if val == "" {
		val = "0"
	}
	res["in"] = val
	timestamp, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return ResponseError(0, "解析失败:"+err.Error(), nil)
	}
	now := time.Unix(timestamp, 0)
	res["out"] = now.Format(TimeFormatA)
	return ResponseSuccess("", res)
}

func (t *Time) DateToTime(val string) interface{} {
	str := util.GetNum(val)
	if str == "" {
		return ResponseError(0, "请输入数据", nil)
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
		return ResponseError(0, "日期格式错误", nil)
	}

	times, err := time.ParseInLocation(layout, str, time.Local)
	if err != nil {
		return ResponseError(0, "解析失败"+err.Error(), nil)
	}
	return ResponseSuccess("", map[string]interface{}{
		"in":  val,
		"out": strconv.FormatInt(times.Unix(), 10),
	})
}
