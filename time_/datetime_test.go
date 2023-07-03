package time_

import (
	"fmt"
	"testing"
	"time"
)

func TestDateTime(t *testing.T) {
	// now := time.Now()
	fmt.Println(time.DateTime)

	now := time.Now()
	anoNow := time.Now()
	fmt.Println(now.Compare(anoNow)) // -1

	timeStr := "2023-01-17 18:00:00"
	formatTime, _ := time.Parse(time.DateTime, timeStr)
	anoFormatTime, _ := time.Parse(time.DateTime, timeStr)
	fmt.Println(formatTime.Compare(anoFormatTime)) // 0
}
