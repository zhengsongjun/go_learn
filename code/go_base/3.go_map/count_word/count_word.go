package main

import (
	"fmt"
	"strings"
)

func count_word(str string) map[string]int {
	strSlice := strings.Split(str, " ")
	fmt.Println(strSlice)
	strMap := make(map[string]int)
	for _, v := range strSlice {
		value, exists := strMap[v]
		if exists {
			strMap[v] = value + 1
		} else {
			strMap[v] = 1
		}
	}
	return strMap
}

func main() {
	fmt.Println(count_word("because our freedom is at stake our freedom to create to influence and to use the power of money"))

}
