package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Calendar struct {
	Month          int          //目前在幾月
	weekdate       int          //從禮拜幾開始
	monthLenth     int          //本月長度
	currentday     int          //目前日期
	extraLine      int          //指定的額外表格長度
	CalendarWriter bytes.Buffer //用來寫入到檔案的 buffer
}

//TODO:add a unblocking constructor
func NewCalendar() *Calendar {
	Calendar := getInput()
	makeHeader(Calendar)
	return Calendar
}
func (instance *Calendar) AddMonth() {
	instance.CalendarWriter.Write([]byte(">|"))
	if instance.weekdate != 0 {
		FillExtra(&instance.CalendarWriter, instance.weekdate)
	}
	for temp := 0; temp < instance.monthLenth; temp++ {
		addDay(instance)
	}
	if instance.weekdate != 0 {
		FillExtra(&instance.CalendarWriter, instance.extraLine+7-instance.weekdate)
	}
	instance.CalendarWriter.Write([]byte("\n"))
	instance.Month++
	instance.currentday = 1
	instance.monthLenth = GetMonthLength(instance.Month)
}

func addDay(instance *Calendar) {
	if instance.weekdate == 0 {
		instance.CalendarWriter.Write([]byte(">|"))
	}
	if instance.weekdate == 0 || instance.weekdate == 6 {
		instance.CalendarWriter.Write([]byte("`"))
	} else {
		instance.CalendarWriter.Write([]byte(" "))
	}

	temp := strconv.Itoa(instance.Month)
	instance.CalendarWriter.Write([]byte(temp))
	instance.CalendarWriter.Write([]byte("/"))
	instance.CalendarWriter.Write([]byte(fmt.Sprintf("%02d", instance.currentday)))

	if instance.weekdate == 0 || instance.weekdate == 6 {
		instance.CalendarWriter.Write([]byte("`"))
	} else {
		instance.CalendarWriter.Write([]byte(" "))
	}

	instance.CalendarWriter.Write([]byte("|"))
	instance.currentday++
	instance.weekdate++
	if instance.weekdate == 7 {
		FillExtra(&instance.CalendarWriter, instance.extraLine)
		instance.CalendarWriter.Write([]byte("\n"))
		instance.weekdate = 0
	}
}

func makeHeader(instance *Calendar) {
	instance.CalendarWriter.Write([]byte(">| `SUN` |  MON  |  TUE  |  WED  |  THU  |  FRI  | `SAT` |"))
	FillExtra(&instance.CalendarWriter, instance.extraLine)
	instance.CalendarWriter.Write([]byte("\n"))
	instance.CalendarWriter.Write([]byte(">|-------|-------|-------|-------|-------|-------|-------|"))
	temp := instance.extraLine
	for ; temp > 0; temp-- {
		instance.CalendarWriter.Write([]byte("-------|"))
	}
	instance.CalendarWriter.Write([]byte("\n"))
	return
}

func FillExtra(buffer *bytes.Buffer, LineCounter int) {
	for ; LineCounter > 0; LineCounter-- {
		buffer.Write([]byte("       |"))
	}
	return
}

func GetMonthLength(Month int) int {
	if Month == 2 {
		return 28
	}
	if Month > 7 {
		return 30 + ((Month % 2) ^ 1)
	}
	return 30 + (Month % 2)
}

// 0 = sun 6 = sunday
func GetLastDay(Start int, Month int) int {
	AddDay := (GetMonthLength(Month)%7 + Start) % 7
	return AddDay
}

func ErrorCheck(err error) {
	if err != nil {
		fmt.Println("while reading %v", err)
		return
	}
}

func getInput() *Calendar {
	var InputDate Calendar
	InputDate.currentday = 1
	var TempBuffer string
	fmt.Println("The month of the calendar is: ")
	fmt.Scanln(&TempBuffer)
	{
		temp, err := strconv.Atoi(TempBuffer)
		ErrorCheck(err)
		InputDate.Month = temp
	}
	fmt.Println("The " + TempBuffer + "/01 is (0 = sunday ... 6 = saturday)")
	{
		fmt.Scanln(&TempBuffer)
		temp, err := strconv.Atoi(TempBuffer)
		ErrorCheck(err)
		InputDate.weekdate = temp
	}
	InputDate.monthLenth = GetMonthLength(InputDate.Month)
	fmt.Println("How many extraline do you want: ")
	{
		fmt.Scanln(&TempBuffer)
		temp, err := strconv.Atoi(TempBuffer)
		ErrorCheck(err)
		InputDate.extraLine = temp
	}
	return &InputDate
}

func main() {
	fmt.Printf("hello\n")
	ACalendar := NewCalendar()
	ACalendar.AddMonth()
	FileName := strconv.Itoa(ACalendar.Month-1) + "_Calendar.md"
	ioutil.WriteFile(FileName, ACalendar.CalendarWriter.Bytes(), 0644)
}
