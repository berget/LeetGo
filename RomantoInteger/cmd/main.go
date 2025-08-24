package main

import (
	"example.com/romantointeger"
	"fmt"
)

func main() {
	fmt.Println(romantointeger.RomanToInt("III"))     // 3
	fmt.Println(romantointeger.RomanToInt("I"))       // 1
	fmt.Println(romantointeger.RomanToInt("MCMXCIV")) // 1994

	fmt.Println(romantointeger.RomanToIntBasic("III"))     // 3
	fmt.Println(romantointeger.RomanToIntBasic("I"))       // 1
	fmt.Println(romantointeger.RomanToIntBasic("MCMXCIV")) // 1994

}
