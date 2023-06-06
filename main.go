package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := Date(2002, 1, 21)
	t2 := time.Now()
	days := t2.Sub(t1).Hours() / 24
	fmt.Println(int(days))
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
