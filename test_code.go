package main

import (
	"fmt"
	"learnGo/DictionaryCode"
)

func main()  {
	region := DictionaryCode.NewGBT2260()
	localCode := region.SearchGBT2260("130104")
	fmt.Println(localCode)
}
