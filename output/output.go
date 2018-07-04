package output

import (
	"fmt"
	"regexp"
	"strings"
	"unsafe"

	"github.com/reiver/go-porterstemmer"
)

func bstring(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Color inspect a party and coloring Abnormal log.
func colorAb(l string) {
	r0 := regexp.MustCompile("critical")
	r1 := regexp.MustCompile("error")
	r2 := regexp.MustCompile("warn")

	for _, w := range strings.Split(l, " ") {
		stem := porterstemmer.StemString(w)
		if r0.MatchString(stem) {
			fmt.Printf("\x1b[31m%s\x1b[0m\n", l)
			return
		}
		if r1.MatchString(stem) {
			fmt.Printf("\x1b[31m%s\x1b[0m\n", l)
			return
		}
		if r2.MatchString(stem) {
			fmt.Printf("\x1b[33m%s\x1b[0m\n", l)
			return
		}
	}
	fmt.Println(l)
}

//LoopLines loop lines to output.
func LoopLines(lines [][]byte, color bool) {
	if color {
		for _, l := range lines {
			colorAb(bstring(l))
		}
		return
	}
	for _, l := range lines {
		fmt.Println(bstring(l))
	}
}
