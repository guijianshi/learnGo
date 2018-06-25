package main

import "fmt"

func main() {
	var CountryCapitalMap map[string]string // 创建集合
	CountryCapitalMap = make(map[string]string)

	/* map插入key - value 对 */
	CountryCapitalMap["中国"] = "北京"
	CountryCapitalMap["日本"] = "东京"
	CountryCapitalMap["美国"] = "华盛顿"

	/* 迭代输出 */
	for country, city := range CountryCapitalMap {
		fmt.Printf("%s的首都是%s\n", country, city)
	}

	/*查看元素在集合中是否存在 */
	captial, ok := CountryCapitalMap["美国"] /*如果确定是真实的,则存在,否则不存在 */

	if (ok) {
		fmt.Println("美国的首都是", captial)
	} else {
		fmt.Println("美国的首都不存在")
	}

	delete(CountryCapitalMap, "美国")
	captial, ok = CountryCapitalMap["美国"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(captial) */
	/*fmt.Println(ok) */
	if (ok) {
		fmt.Println("美国的首都是", captial)
	} else {
		fmt.Println("美国的首都不存在")
	}
}
