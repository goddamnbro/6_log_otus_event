package main

import (
	"fmt"
	"io"
	"time"
)

type OtusEvent interface {
	Formatter() string
}

type HwAccepted struct {
	Id int
	Grade int
}

type HwSubmitted struct {
	Id int
	Code string
	Comment string
}

func (hw *HwAccepted) Formatter() string {
	dt := time.Now()
	return fmt.Sprintf("[%s] accepted id:%d grade:%d",
		dt.Format("2006-01-02"), hw.Id, hw.Grade)
}

func (hw *HwSubmitted) Formatter() string {
	dt := time.Now()
	return fmt.Sprintf("[%s] submitted id:%d code:%s Comment:%s",
		dt.Format("2006-01-02"), hw.Id, hw.Code, hw.Comment)
}


func LogOtusEvent(event OtusEvent, w io.Writer) {
	w.Write([]byte(event.Formatter()))
}
