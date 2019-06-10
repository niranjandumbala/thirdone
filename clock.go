package clock

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"time"
)

var seconds float64 = 0
var tick string = "tick"
var tock string = "tock"
var bong string = "bong"

func RunClock() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			seconds = seconds + 1
			Update("input.txt")
			fmt.Println(GetPrintLabel(seconds))
		}
	}()

	time.Sleep(3 * time.Hour)
	ticker.Stop()
}

func GetPrintLabel(secs float64) string {
	if secs >= 3600 && math.Mod(secs, float64(3600)) == 0 {
		return "bong"
	}
	if secs >= float64(60) && math.Mod(secs, float64(60)) == 0 {
		return tock
	}
	return tick
}

func Update(fileName string) bool {
	line, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
		return false
	}
	str := string(line)
	input := strings.Split(str, ",")
	tick = input[0]
	tock = input[1]
	bong = input[2]
	return true
}
