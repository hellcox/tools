package timetool

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"time"
	"tools/util"
)

var (
	data            map[string]binding.String //绑定数据
	dataKey         []string                  //绑定数据KEY
	nowTime         time.Time                 //当前时间
	nowTicker       *time.Ticker              //当前时间定时器
	tickerRunStatus int                       //0.停止 1.运行
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
	k1               = "k1"
	k2               = "k2"
	k3               = "k3"
	k4               = "k4"
	k5               = "k5"
	v1               = "v1"
	v2               = "v2"
	v3               = "v3"
	v4               = "v4"
	v5               = "v5"
)

// 初始化时间
func initData() {
	nowTime = time.Now()
	data = make(map[string]binding.String)
	dataKey = append(dataKey, []string{k1, k2, k3, k4, k5}...)
	dataKey = append(dataKey, []string{v1, v2, v3, v4, v5}...)
	for _, key := range dataKey {
		value := ""
		data[key] = binding.BindString(&value)
	}

	_ = data[k1].Set(strconv.FormatInt(nowTime.Unix(), 10))
	_ = data[v1].Set(nowTime.Format(TimeFormatA))

	go runTime()
}

func MakeGridLayout(_ fyne.Window) fyne.CanvasObject {
	initData()

	border := util.MakeBorder(5, 5)
	border2 := util.MakeBorder(5, 5)
	border3 := util.MakeBorder(5, 5)
	border4 := util.MakeBorder(5, 5)

	ch := float32(4)

	nowTimeLayout := container.NewGridWrap(fyne.Size{Width: 650, Height: 38}, nowTimeLayout())
	h := ch + float32(38)

	h += 20
	timeToDate := container.NewGridWrap(fyne.Size{Width: 650, Height: 38}, timeToDateLayout())
	timeToDate.Move(fyne.NewPos(0, h))
	h += ch + 38

	timeToDate2 := container.NewGridWrap(fyne.Size{Width: 650, Height: 38}, timeToDateLayout())
	timeToDate2.Move(fyne.NewPos(0, h))
	h += ch + 38

	h += 20

	dateToTime := container.NewGridWrap(fyne.Size{Width: 650, Height: 38}, dateToTimeLayout())
	dateToTime.Move(fyne.NewPos(0, h))
	h += ch + 38

	dateToTime2 := container.NewGridWrap(fyne.Size{Width: 650, Height: 38}, dateToTimeLayout())
	dateToTime2.Move(fyne.NewPos(0, h))
	h += ch + 38

	middle := container.NewWithoutLayout(
		nowTimeLayout,
		timeToDate,
		timeToDate2,
		dateToTime,
		dateToTime2,
	)

	return container.NewBorder(border, border2, border3, border4, middle)
}

// 运行时间
func runTime() {
	tickerRunStatus = 1
	nowTime = time.Now()
	_ = data[k1].Set(strconv.FormatInt(nowTime.Unix(), 10))
	_ = data[v1].Set(nowTime.Format(TimeFormatA))
	nowTicker = time.NewTicker(1 * time.Second)
	for range nowTicker.C {
		nowTime = nowTime.Add(1 * time.Second)
		_ = data[k1].Set(strconv.FormatInt(nowTime.Unix(), 10))
		_ = data[v1].Set(nowTime.Format(TimeFormatA))

	}
}

// 当前时间
func nowTimeLayout() fyne.CanvasObject {
	input := widget.NewEntryWithData(data[k1])
	output := widget.NewEntryWithData(data[v1])

	button := widget.NewButton("暂停/开始", func() {
		fmt.Println(tickerRunStatus)
		if tickerRunStatus == 1 {
			nowTicker.Stop()
			tickerRunStatus = 0
		} else if tickerRunStatus == 0 {
			go runTime()
			tickerRunStatus = 1
		}
	})
	res := container.NewGridWithColumns(3, input, button, output)
	return res
}

// 时间转日期
func timeToDateLayout() fyne.CanvasObject {
	inVal, outVal := "", ""
	in := binding.BindString(&inVal)
	out := binding.BindString(&outVal)
	input := widget.NewEntryWithData(in)
	input.OnChanged = func(s string) {
		in.Set(s)
		out.Set("")
	}
	output := widget.NewEntryWithData(out)

	button := widget.NewButton("时间 => 日期", func() {
		input, _ := in.Get()
		str := util.GetNum(input)
		if len(str) != len(input) {
			_ = out.Set("输入错误")
			return
		}
		num, _ := strconv.ParseInt(input, 10, 64)
		a := time.Unix(num, 0)
		_ = out.Set(a.Format(TimeFormatA))
	})
	res := container.NewGridWithColumns(3, input, button, output)
	return res
}

// 日期转时间
func dateToTimeLayout() fyne.CanvasObject {
	inVal, outVal := "", ""
	in := binding.BindString(&inVal)
	out := binding.BindString(&outVal)
	input := widget.NewEntryWithData(in)
	input.OnChanged = func(s string) {
		in.Set(s)
		out.Set("")
	}
	output := widget.NewEntryWithData(out)

	button := widget.NewButton("日期 => 时间", func() {
		input, _ := in.Get()
		str := util.GetNum(input)
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
		}

		times, err := time.Parse(layout, str)
		if err != nil {
			_ = out.Set(err.Error())
			return
		}
		_ = out.Set(strconv.FormatInt(times.Unix(),10))
	})
	res := container.NewGridWithColumns(3, input, button, output)
	return res
}
