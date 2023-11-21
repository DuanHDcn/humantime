package main

import (
	"flag"
	"fmt"
)

func main() {
	secondPtr := flag.Int64("second", -1, "-second=digit")
	flag.Parse()

	if secondPtr == nil || *secondPtr < 0 {
		fmt.Println("[Usage] humantime -second=digit")
		return
	}

	seconds := *secondPtr
	second := seconds % 60
	minute := seconds / 60 % 60
	hour := seconds / 3600 % 24
	day := seconds / 3600 / 24

	output := ""
	output += ternaryString(day > 0, fmt.Sprintf("%dd ", day), "")
	output += ternaryString(hour > 0, fmt.Sprintf("%dh ", hour), "")
	output += ternaryString(minute > 0, fmt.Sprintf("%dm ", minute), "")
	output += ternaryString(second >= 0, fmt.Sprintf("%ds ", second), "")

	fmt.Println(output)
}

func ternaryString(cond bool, t, f string) string {
	if cond {
		return t
	}
	return f
}
