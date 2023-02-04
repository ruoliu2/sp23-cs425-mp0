package main

import (
	"fmt"
	"time"
)

func main() {
	// print current time with precision to microseconds, convert to seconds
	// 1610688413.743385
	//	keep 6 digits after the decimal point
	fmt.Printf("%d\n", time.Now().UnixMicro())
	fmt.Printf("%f\n", float64(time.Now().UnixMicro())/1e6)
	fmt.Printf("%.6f", float64(time.Now().UnixMicro())/1e6)
}
