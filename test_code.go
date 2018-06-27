package main

import (
	"fmt"
	"learnGo/DictionaryCode"
	"time"
)

func main()  {
	start_time := time.Now()
	region := DictionaryCode.NewGBT2260()
	localCode := region.SearchGBT2260("110101")
	fmt.Println(localCode)
	end_time := time.Now()
	var cost_time  float64
	fmt.Println(testStringParse("110101"))
	cost_time = (float64(end_time.Nanosecond() - start_time.Nanosecond()) / (float64(1000000000)))
	fmt.Println(cost_time)
}

func testStringParse(code string) []string  {
	return DictionaryCode.StringParse(code)
}
