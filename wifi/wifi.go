package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var ssidregex = regexp.MustCompile(`ESSID:\"(.+)\"\s`)
var qualityregex = regexp.MustCompile(`Link Quality=(\d+)\D(\d+)\s`)
var ipregex = regexp.MustCompile(`inet (.+)\/`)
var ssid, ip string
var strength float64

func main() {
	dev := "wlp1s0"
	data, err := ioutil.ReadFile("/sys/class/net/" + dev + "/operstate")
	if err != nil {
		fmt.Printf("error\nerror\n#ff0000\n")
		os.Exit(33)
	}
	if strings.HasPrefix(string(data), "down") {
		fmt.Printf("down\ndown\n#ff0000\n")
		os.Exit(33)
	}

	out, _ := exec.Command("iwconfig", dev).Output()
	if ssidregex.MatchString(string(out)) {
		ssid = ssidregex.FindStringSubmatch(string(out))[1]
	}
	if qualityregex.MatchString(string(out)) {
		quality := qualityregex.FindStringSubmatch(string(out))[1]
		qval, _ := strconv.ParseInt(quality, 10, 64)
		if err != nil {
			fmt.Println(err.Error())
		}
		qualitymax := qualityregex.FindStringSubmatch(string(out))[2]
		qmaxval, _ := strconv.ParseInt(qualitymax, 10, 64)
		strength = float64(qval) / float64(qmaxval) * 100

	}
	out, _ = exec.Command("ip", "addr", "show", dev).Output()
	if ipregex.MatchString(string(out)) {
		ip = ipregex.FindStringSubmatch(string(out))[1]
	}
	color := "ff0000"
	if strength > 25.0 {
		color = "ffff00"
	}
	if strength > 50.0 {
		color = "00ff00"
	}
	fmt.Printf("%s (%.0f%%) %s\n%s (%.0f%%) %s\n#%s\n", ssid, strength, ip, ssid, strength, ip, color)
	//	fmt.Println(string(out))
}
