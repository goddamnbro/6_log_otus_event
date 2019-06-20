package main

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestLogHwAccepted(t *testing.T) {

	var hwa OtusEvent
	hwa = &HwAccepted{Id:1, Grade:10}

	var buf bytes.Buffer
	LogOtusEvent(hwa, &buf)

	expected := fmt.Sprintf("[%s] accepted id:1 grade:10", time.Now().Format("2006-01-02"))
	if buf.String() != expected {
		t.Errorf("LogOtusEvent works incorrect for the HwAccepted")
	}
}


func TestLogHwSubmitted(t *testing.T) {

	var hws OtusEvent
	hws = &HwSubmitted{Id:2, Code:"<html></html>", Comment:"my pretty code"}

	var buf bytes.Buffer
	LogOtusEvent(hws, &buf)

	expected := fmt.Sprintf(
		"[%s] submitted id:2 code:<html></html> Comment:my pretty code",
		time.Now().Format("2006-01-02"),
	)
	if buf.String() != expected {
		t.Errorf("LogOtusEvent works incorrect for the HwSubmitted")
	}
}
