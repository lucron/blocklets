package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	data, err := ioutil.ReadFile("/sys/class/backlight/intel_backlight/brightness")
	if err != nil {
		fmt.Printf("%s\n%s\nFF0000")
		return
	}
	intval, err := strconv.ParseInt(string(data)[:len(string(data))-1], 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	ret := float64(intval) / 4648.0 * 100.0
	fmt.Printf("%.0f%%\n%.0f%%\n", ret, ret)
	os.Exit(0)
}
