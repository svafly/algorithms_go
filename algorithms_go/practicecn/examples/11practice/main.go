package main

import "fmt"

func main() {
	strarray := []string{"I", "am", "stupid", "and", "weak"}
	replace(strarray)
	fmt.Println(strarray)
}

func replace(strarr []string) {
	for i := 0; i < len(strarr); i++ {
		if strarr[i] == "stupid" {
			strarr[i] = "smart"
		} else if strarr[i] == "weak" {
			strarr[i] = "strong"
		}
	}
}
